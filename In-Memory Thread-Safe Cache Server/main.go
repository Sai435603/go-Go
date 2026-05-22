package main

import (
	"bufio"
	"container/list"
	"fmt"
	"hash/fnv"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	numShards   = 32
	shardCap    = 1000 // max items per shard (32,000 total)
	aofFileName = "appendonly.aof"
)

type item struct {
	key        string
	value      string
	expiration int64
}

type shard struct {
	mu       sync.RWMutex
	capacity int
	items    map[string]*list.Element
	lruList  *list.List
}

type Cache struct {
	shards  []*shard
	aofChan chan string
}

func NewCache() *Cache {
	c := &Cache{
		shards:  make([]*shard, numShards),
		aofChan: make(chan string, 10000), // buffer to avoid blocking workers
	}

	for i := 0; i < numShards; i++ {
		c.shards[i] = &shard{
			capacity: shardCap,
			items:    make(map[string]*list.Element),
			lruList:  list.New(),
		}
	}

	c.loadAOF()
	go c.aofWriter()

	return c
}

// hash determines which shard a key belongs to
func (c *Cache) getShard(key string) *shard {
	h := fnv.New32a()
	h.Write([]byte(key))
	return c.shards[h.Sum32()%numShards]
}

func (c *Cache) Set(key, value string, ttlSeconds int) {
	var exp int64
	if ttlSeconds > 0 {
		exp = time.Now().Add(time.Duration(ttlSeconds) * time.Second).UnixNano()
	}

	s := c.getShard(key)
	s.mu.Lock()

	if elem, ok := s.items[key]; ok {
		s.lruList.MoveToFront(elem)
		elem.Value.(*item).value = value
		elem.Value.(*item).expiration = exp
	} else {
		if s.lruList.Len() >= s.capacity {
			c.evictOldest(s)
		}
		newElem := s.lruList.PushFront(&item{key, value, exp})
		s.items[key] = newElem
	}
	s.mu.Unlock()

	// Async persistence - Format: SET key expiration value
	c.aofChan <- fmt.Sprintf("SET %s %d %s\n", key, exp, value)
}

func (c *Cache) Get(key string) (string, bool) {
	s := c.getShard(key)
	s.mu.Lock()
	defer s.mu.Unlock()

	if elem, ok := s.items[key]; ok {
		it := elem.Value.(*item)
		if it.expiration > 0 && time.Now().UnixNano() > it.expiration {
			c.removeElement(s, elem)
			return "", false
		}
		s.lruList.MoveToFront(elem)
		return it.value, true
	}

	return "", false
}

func (c *Cache) Delete(key string) {
	s := c.getShard(key)
	s.mu.Lock()
	if elem, ok := s.items[key]; ok {
		c.removeElement(s, elem)
	}
	s.mu.Unlock()

	c.aofChan <- fmt.Sprintf("DEL %s\n", key)
}

// Helpers for shard management (must be called with lock held)
func (c *Cache) evictOldest(s *shard) {
	elem := s.lruList.Back()
	if elem != nil {
		c.removeElement(s, elem)
	}
}

func (c *Cache) removeElement(s *shard, elem *list.Element) {
	s.lruList.Remove(elem)
	it := elem.Value.(*item)
	delete(s.items, it.key)
}

//Persistence(AOF)

func (c *Cache) aofWriter() {
	f, err := os.OpenFile(aofFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("failed to open AOF: %v", err)
		return
	}
	defer f.Close()

	for cmd := range c.aofChan {
		f.WriteString(cmd)
	}
}

func (c *Cache) loadAOF() {
	f, err := os.Open(aofFileName)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("failed to load AOF: %v", err)
		}
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) == 0 {
			continue
		}

		// Loading back the spaces correctly
		if parts[0] == "SET" && len(parts) >= 4 {
			key := parts[1]
			exp, _ := strconv.ParseInt(parts[2], 10, 64)
			value := strings.Join(parts[3:], " ") // Rejoin the spaces!

			s := c.getShard(key)
			s.mu.Lock()
			if elem, ok := s.items[key]; ok {
				elem.Value.(*item).value = value
				elem.Value.(*item).expiration = exp
				s.lruList.MoveToFront(elem)
			} else {
				if s.lruList.Len() >= s.capacity {
					c.evictOldest(s)
				}
				s.items[key] = s.lruList.PushFront(&item{key, value, exp})
			}
			s.mu.Unlock()
		} else if parts[0] == "DEL" && len(parts) >= 2 {
			s := c.getShard(parts[1])
			s.mu.Lock()
			if elem, ok := s.items[parts[1]]; ok {
				c.removeElement(s, elem)
			}
			s.mu.Unlock()
		}
	}
}

// TCP Server

func handleConnection(conn net.Conn, cache *Cache) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		parts := strings.Split(text, " ")
		cmd := strings.ToUpper(parts[0])

		switch cmd {
		case "SET":
			if len(parts) < 3 {
				conn.Write([]byte("ERR missing args\n"))
				continue
			}
			key := parts[1]
			value := strings.Join(parts[2:], " ") // Join everything after the key
			cache.Set(key, value, 0)
			conn.Write([]byte("OK\n"))

		case "SETEX": // Use this for TTLs SETEX <key> <time in sec> <value>
			if len(parts) < 4 {
				conn.Write([]byte("ERR missing args (SETEX key ttl value)\n"))
				continue
			}
			key := parts[1]
			ttl, _ := strconv.Atoi(parts[2])
			value := strings.Join(parts[3:], " ") // Join everything after the ttl
			cache.Set(key, value, ttl)
			conn.Write([]byte("OK\n"))

		case "GET":
			if len(parts) < 2 {
				conn.Write([]byte("ERR missing key\n"))
				continue
			}
			if val, ok := cache.Get(parts[1]); ok {
				conn.Write([]byte(fmt.Sprintf("%s\n", val)))
			} else {
				conn.Write([]byte("(nil)\n"))
			}

		case "DEL":
			if len(parts) < 2 {
				continue
			}
			cache.Delete(parts[1])
			conn.Write([]byte("OK\n"))

		case "QUIT":
			return
		}
	}
}

func main() {
	cache := NewCache()
	port := ":8000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("server start failed: %v", err)
	}
	defer listener.Close()

	log.Printf("cache server listening on %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn, cache)
	}
}

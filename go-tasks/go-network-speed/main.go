package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// get network bytes (rx, tx)
func getNetBytes() (uint64, uint64) {
	file, _ := os.Open("/proc/net/dev")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rx, tx uint64

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip headers
		if !strings.Contains(line, ":") {
			continue
		}

		parts := strings.Split(line, ":")
		iface := strings.TrimSpace(parts[0])

		// skip loopback
		if iface == "lo" {
			continue
		}

		data := strings.Fields(parts[1])

		rxBytes, _ := strconv.ParseUint(data[0], 10, 64)
		txBytes, _ := strconv.ParseUint(data[8], 10, 64)

		rx += rxBytes
		tx += txBytes
	}

	return rx, tx
}

func main() {
	prevRx, prevTx := getNetBytes()

	for {
		time.Sleep(1 * time.Second)

		currRx, currTx := getNetBytes()

		down := currRx - prevRx
		up := currTx - prevTx

		prevRx = currRx
		prevTx = currTx

		// clear screen
		fmt.Print("\033[H\033[2J")

		fmt.Printf("⬇ Download: %.2f KB/s\n", float64(down)/1024)
		fmt.Printf("⬆ Upload  : %.2f KB/s\n", float64(up)/1024)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

// Course method ,middlewares goes into sep files
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Crud Api")

	router := mux.NewRouter()

	courses = append(courses, Course{CourseId: "1", CourseName: "Sai raam", CoursePrice: 69, Author: &Author{Fullname: "sai", Website: "sairaam.in"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "BGMI", CoursePrice: 699, Author: &Author{Fullname: "UNQ Gamer", Website: "unqsairaam.in"}})

	//routing
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listening to the port 8000
	log.Fatal(http.ListenAndServe(":8000", router))

}

//controllers - sep flies

//serve home route

// [writer, reader]
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> welcome sai raam sai raam <h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println("params of the getOneCourse Endpoint ", params)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode(fmt.Sprintf("No Course Found With the Id %s", params["id"]))
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course...")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send the data")
		return
	}

	// handling case if a body is {} [empty JSON]

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	// checking course is already is there in the courses
	for _, crs := range courses {
		if crs.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Sorry this course is already exist :(")
			return
		}
	}

	//gen unq id , and conv to str
	// then append it to courses slice

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update on one course called...")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove, add with my id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			// removed
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course id not found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting one course called...")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}

	json.NewEncoder(w).Encode("Course id not found")
}

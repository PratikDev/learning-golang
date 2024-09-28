package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Struct for course - import
type Course struct {
	CourseId string  `json:"id"`
	Title    string  `json:"title"`
	Price    int     `json:"price"`
	Discount int     `json:"discount"`
	Author   *Author `json:"author"`
}

// Course auther - import
type Author struct {
	Fullname string   `json:"fullname"`
	Website  string   `json:"website"`
	Socials  []string `json:"socials"`
}

func main() {
	const PORT = 3000

	// seeding
	courses = append(courses, Course{CourseId: "1", Title: "Reactjs", Price: 100, Discount: 20, Author: &Author{Fullname: "Pratik Dev", Website: "https://pratik.dev", Socials: []string{"https://github.com/pratikdev"}}}, Course{CourseId: "2", Title: "NextJS", Price: 200, Discount: 25, Author: &Author{Fullname: "Pratik Dev", Website: "https://pratik.dev", Socials: []string{"https://facebook.com/pratikdev"}}})

	// routing
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getCourse).Methods("GET")
	router.HandleFunc("/course", createCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	portString := ":" + strconv.Itoa(PORT)
	log.Fatal(http.ListenAndServe(portString, router))
}

// fake DB
var courses []Course

/* helper func */
func isValidDiscount(discount int) bool {
	return discount >= 0 && discount <= 100
}

func (c *Course) isValid() bool {
	return (c.Title != "" && !(c.Price < 0) && isValidDiscount(c.Discount) && c.Author != nil)
}

/* controllers - import */

// home handler route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to home</h1>"))
}

// get all courses list
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// get a course by id
func getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	courseId := params["id"]

	for _, course := range courses {
		if course.CourseId == courseId {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "No courses found with the given id"})
}

// add a course
func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// if body is nil
	if r.Body == nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Body is not allowed to be null"})
		return
	}

	var course Course

	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Error parsing the request body"})
		return
	}

	// if course if invalid
	if !course.isValid() {
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request"})
		return
	}

	// if course title is duplicate
	for _, _course := range courses {
		if _course.Title == course.Title {
			json.NewEncoder(w).Encode(map[string]string{"message": "Duplicate course title is not allowed"})
			return
		}
	}

	// generating random course id
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// adding new course to the slice
	courses = append(courses, course)

	json.NewEncoder(w).Encode(map[string]string{"message": "Course created successfully"})
}

// update a course
func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// read the params
	params := mux.Vars(r)
	courseId := params["id"]

	// loop thru courses and search for `courseId`
	for index, course := range courses {
		if course.CourseId == courseId {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course

			err := json.NewDecoder(r.Body).Decode(&course)
			if err != nil {
				json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request"})
				return
			}

			course.CourseId = courseId
			courses = append(courses, course)
			json.NewEncoder(w).Encode(map[string]string{"message": "Course updated succesfully"})
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "No course found with the given id"})
}

// delete a course
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	courseId := params["id"]

	for index, course := range courses {
		if course.CourseId == courseId {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Course deleted succesfully"})
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Course not found with the given id"})
}

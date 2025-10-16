package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Model for courses - file

type Course struct {
	CourseId    string  `json:"id"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Auther      *Auther `json:"author"`
}

type Auther struct {
	Name    string `json:"fullname"`
	Website string `json:"website"`
}

// fake db
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()

	//seeding data
	courses = append(courses, Course{CourseId: "2", CourseName: "React.js", CoursePrice: 599, Auther: &Auther{Name: "Jai Srivastava", Website: "http://byteflow.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Next.js", CoursePrice: 299, Auther: &Auther{Name: "Jai Srivastava", Website: "http://byteflow.com"}})

	//routes
	r.HandleFunc("/", servehome).Methods("GET")
	r.HandleFunc("/getallcourse", getAllCourses).Methods("GET")
	r.HandleFunc("/getcoursebyid/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/addcourse", addCourse).Methods("POST")
	r.HandleFunc("/updatecourse/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/deletecourse/{id}", deleteCourse).Methods("DELETE")
	r.HandleFunc("/deleteall", deleteAllCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":3000", r))
}

//controllers

// serve home route
func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to API Lecture."))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Course By Id.")
	w.Header().Set("Content-Type", "application/json")

	//grab id form body of the api call
	params := mux.Vars(r)

	//action plan to get the course and return it
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given Id.")
}

func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Course.")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data.")
	}

	//{} data case
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data found.")
		return
	}

	for _, v := range courses {
		if course.CourseName == v.CourseName {
			json.NewEncoder(w).Encode("Course Already Present.")
			return
		}
	}

	//generate id and convert them to string and append them into the slice
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Course.")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data.")
	}

	params := mux.Vars(r)

	for i, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			var coursePtr Course
			_ = json.NewDecoder(r.Body).Decode(&coursePtr)
			coursePtr.CourseId = params["id"]
			courses = append(courses, coursePtr)
			json.NewEncoder(w).Encode(coursePtr)
			return
		}
	}
	json.NewEncoder(w).Encode("No such course found.")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Course.")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data.")
	}

	params := mux.Vars(r)

	for i, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			json.NewEncoder(w).Encode("Course Deleted.")
			return
		}
	}
	json.NewEncoder(w).Encode("No such course found.")
}

func deleteAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete All Course.")
	courses = append(courses[:0])

	json.NewEncoder(w).Encode("All Data Deleted.")

}

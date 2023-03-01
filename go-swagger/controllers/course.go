package controllers

import (
	"encoding/json"
	"fmt"
	"go-swagger/models"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// fake DB
var courses []models.Course

func SeedData() {
	//seeding
	courses = append(courses, models.Course{CourseId: "1", CourseName: ".Net", CoursePrice: 5000,
		Author: &models.Author{FullName: "Aarti Chhasiya", Website: "net.in"}})
	courses = append(courses, models.Course{CourseId: "2", CourseName: ".Net Core", CoursePrice: 5000,
		Author: &models.Author{FullName: "Aarti Parmar", Website: "netcore.in"}})
	courses = append(courses, models.Course{CourseId: "3", CourseName: "Java", CoursePrice: 5000,
		Author: &models.Author{FullName: "Jon Sopheg", Website: "java.in"}})
	courses = append(courses, models.Course{CourseId: "4", CourseName: "React JS", CoursePrice: 5000,
		Author: &models.Author{FullName: "Ven solh", Website: "reactjs.in"}})
	courses = append(courses, models.Course{CourseId: "5", CourseName: "MERN JS", CoursePrice: 5000,
		Author: &models.Author{FullName: "Pooja Mehra", Website: "mernjs.in"}})
	courses = append(courses, models.Course{CourseId: "6", CourseName: "Python", CoursePrice: 5000,
		Author: &models.Author{FullName: "Balaguru swami", Website: "python.in"}})
	courses = append(courses, models.Course{CourseId: "7", CourseName: "GoLang", CoursePrice: 5000,
		Author: &models.Author{FullName: "Swen john", Website: "golang.in"}})
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json") //set header
	json.NewEncoder(w).Encode(courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json") //set header

	//grab id from request
	params := mux.Vars(r)

	//loop through couses, find matching is and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found for the given id")
	return
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json") //set header

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about data will be like - {}
	var course models.Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//generate unique id, convert it to string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one courses")
	w.Header().Set("Content-Type", "application/json") //set header

	//grab id from request
	params := mux.Vars(r)

	//as it's a no db call and it's a in-memory operation, loop through slices-find item by id got in request-remove that item-add another item with the same id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var NewCourse Course
			_ = json.NewDecoder(r.Body).Decode(&NewCourse)
			NewCourse.CourseId = params["id"]
			courses = append(courses, NewCourse)
			json.NewEncoder(w).Encode(NewCourse)
			return
		}
	}
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one courses")
	w.Header().Set("Content-Type", "application/json") //set header

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("The record is deleted")
			break
		}
	}
}

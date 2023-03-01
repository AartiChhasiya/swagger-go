package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/AartiChhasiya/swagger-go/go-openapi/controllers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to API Section")

	controllers.SeedData()

	r := mux.NewRouter()

	//routing
	r.HandleFunc("/api", controllers.ServeHome).Methods("GET")
	r.HandleFunc("/api/courses", controllers.GetAllCourses).Methods("GET")
	r.HandleFunc("/api/course/{id}", controllers.GetOneCourse).Methods("GET")
	r.HandleFunc("/api/course", controllers.CreateOneCourse).Methods("POST")
	r.HandleFunc("/api/course/{id}", controllers.UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/api/course/{id}", controllers.DeleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":7000", r))

}

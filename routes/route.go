package routes

import (
	"newgo/test/service/handler"

	"github.com/gorilla/mux"
)

// all routes
func Addroutes(r *mux.Router) {
	routesCollege()
	routesCourse()
	routesReadingMaterial()
	routesTask()
	routesUserProfile()
}

// reading material curd operation routes
func routesReadingMaterial() {
	r.HandleFunc("/addreadingmaterial", handler.AddReadingMaterialHandler).Methods("POST")
	r.HandleFunc("/getreadingmaterial", handler.GetReadingMaterialHandler).Methods("GET")
	r.HandleFunc("/deletereadingmaterial/{id}", handler.DeleteReadingMaterialHandler).Methods("DELETE")
	r.HandleFunc("/updatereadingmaterial/{id}", handler.UpdateReadingMaterialHandler).Methods("PUT")

}

// colleges curd operation routes
func routesCollege() {
	r.HandleFunc("/addcollege", handler.AddCollegeHandler).Methods("POST")
	r.HandleFunc("/getcollege", handler.GetCollegeHandler).Methods("GET")
	r.HandleFunc("/deletecollege/{id}", handler.DeletecollegeHandler).Methods("DELETE")
	r.HandleFunc("/updatecollege/{id}", handler.UpdatecollegeHandler).Methods("PUT")

}

// course curd operation routes
func routesCourse() {
	r.HandleFunc("/addcourse", handler.AddCourseHandler).Methods("POST")
	r.HandleFunc("/getcourse", handler.GetCourseHandler).Methods("GET")
	r.HandleFunc("/deletecourse/{id}", handler.DeleteCourseHandler).Methods("DELETE")
	r.HandleFunc("/updatecourse/{id}", handler.UpdateCourseHandler).Methods("PUT")

}

// task curd operation routes
func routesTask() {
	r.HandleFunc("/addtask", handler.AddTaskHandler).Methods("POST")
	r.HandleFunc("/gettask", handler.GetTaskHandler).Methods("GET")
	r.HandleFunc("/deletetask/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/updatetask/{id}", handler.UpdateTaskHandler).Methods("PUT")

}

// userprofile curd operation routes
func routesUserProfile() {
	r.HandleFunc("/addprofile", handler.AddUserProfileHandler).Methods("POST")
	r.HandleFunc("/updateprofile/{roll_no}", handler.UpdateUserProfileHandler).Methods("PUT")
	r.HandleFunc("/deleteprofile/{roll_no}", handler.DeleteUserProfileHandler).Methods("DELETE")
	r.HandleFunc("/getprofile", handler.GetUserProfileHandler).Methods("GET")

}

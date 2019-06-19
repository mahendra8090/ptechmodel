package handler

import (
	"encoding/json"
	"net/http"
	"newgo/test/model"
)

/*
handler for curd opertion of user profile
*/

// userprofile handler
func AddUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var u model.User_profile
	_ = json.NewDecoder(r.Body).Decode(&u)
	json.NewEncoder(w).Encode(u)

	model.Adduserprofile(&u)
}

// update user profile handler
func UpdateUserProfileHandler(w http.ResponseWriter, r *http.Request) {

}

// delete user profile handler
func DeleteUserProfileHandler(w http.ResponseWriter, r *http.Request) {

}

// get user profile handler
func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {

}

// course update handler
func UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {

}

// course delete handler
func DeleteCourseHandler(w http.ResponseWriter, r *http.Request) {

}

// course getting handler
func GetCourseHandler(w http.ResponseWriter, r *http.Request) {

}

// add course handler
func AddCourseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var u model.Course
	_ = json.NewDecoder(r.Body).Decode(&u)
	json.NewEncoder(w).Encode(u)

	model.AddCourse(&u)
}

/* curd operation for college data
 */
// add college data handler
func AddCollegeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var u model.College
	_ = json.NewDecoder(r.Body).Decode(&u)
	json.NewEncoder(w).Encode(u)

	model.AddCollege(&u)
}

// update college data handler
func UpdateCollegeHandler(w http.ResponseWriter, r *http.Request) {

}

// delete college data handler
func DeleteCollegeHandler(w http.ResponseWriter, r *http.Request) {

}

// get college handler
func GetCollegeHandler(w http.ResponseWriter, r *http.Request) {

}

/*
 handlers for curd operation of reading material
*/
// add reading material handler
func AddReadingMaterialHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var u model.ReadingMaterial
	_ = json.NewDecoder(r.Body).Decode(&u)
	json.NewEncoder(w).Encode(u)

	model.AddReadingMaterial(&u)
}

//update reading material handler
func UpdateReadingMaterialHandler(w http.ResponseWriter, r *http.Request) {

}

// delete reading material handler
func DeleteReadingMaterialHandler(w http.ResponseWriter, r *http.Request) {

}

// get reading material handler
func GetReadingMaterialHandler(w http.ResponseWriter, r *http.Request) {

}

/*
handler for curd operation of task
*/

//  add task handler
func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var u model.Task
	_ = json.NewDecoder(r.Body).Decode(&u)
	json.NewEncoder(w).Encode(u)

	model.AddTask(&u)
}

// update task handler
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
}

// delete task handler
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

}

// get task handler
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {

}

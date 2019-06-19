package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Users struct {
	ID              string `json:"id"`
	User_Profile_Id string `json:"user_profileid"`
	Email_Id        string `json:"emailid"`
	Password        string `json:"password"`
	Auth_Id         string `json:"authid"`
	Created_at      string `json:"careatedat"`
	Update_at       string `json:"updateat"`
	Active_status   string `json:"activatestatus"`
	Points          string `json:"points"`
}
type User_profile struct {
	Roll_no      string     `form:"roll_no" query:"roll_no" json:"roll_no,omitempty",not null`
	FirstName    string     `form:"first_name" query:"first_name" json:"first_name,omitempty"`
	LastName     string     `form:"last_name" query:"last_name" json:"last_name,omitempty"`
	Dob          string     `form:"dob" query:"dob" json:"dob,omitempty",not null`
	College_id   string     `form:"college_id" query:"college_id" json:"college_id,omitempty",not null`
	Dept         string     `form:"dept" query:"dept" json:"dept,omitempty",not null`
	Phone_no     string     `form:"phone_no" query:"phone_no" json:"phone_no,omitempty",not null`
	Alt_Phone_no string     `form:"alt_phone_no" query:"alt_phone_no" json:"alt_phone_no,omitempty"`
	Batch_id     string     `form:"batch_id" query:"batch_id" json:"batch_id,omitempty"`
	State        string     `form:"state" query:"state" json:"state,omitempty"`
	City         string     `form:"city" query:"city" json:"city,omitempty"`
	Srn_usn      string     `form:"srn_usn" query:"srn_usn" json:"srn_usn"`
	Semester     string     `form:"semester" query:"semester" json:"semester"`
	CreatedAt    time.Time  `form:"date_created" query:"date_created" json:"date_created"`
	UpdatedAt    time.Time  `form:"date_updated" query:"date_updated" json:"date_updated"`
	DeletedAt    *time.Time `sql:"index"`
}

type Course struct {
	gorm.Model
	CourseID          string `form:"courseid" query:"courseid" json:"courseid,omitempty"`
	CourseName        string `form:"coursename" query:"coursename" json:"coursename,omitempty"`
	CourseDescription string `form:"coursedescription" query:"coursedescription" json:"coursedescription,omitempty"`
	CourseSylabbus    string `form:"coursesylabbus" query:"coursesylabbus" json:"coursesylabbus,omitempty"`
}
type College struct {
	gorm.Model
	Name       string `form:"name" query:"name" json:"name"`
	City       string `form:"city" query:"city" json:"city"`
	State      string `form:"state" query:"state" json:"state"`
	University string `form:"university" query:"university" json:"university"`
	CollegeID  string `form:"collegeid" query:"collegeid" json:"collegeid"`
}
type ReadingMaterial struct {
	gorm.Model
	MaterialID string `form:"materialid" query:"materialid" json:"materialid"`
	Path       string `form:"path" query:"path" json:"path"`
	UpdatePath string `form:"updatepath" query:"updatepath" json:"updatepath"`
}
type Task struct {
	gorm.Model
	CourseID          string `form:"courseid" query:"courseid" json:"courseid"`
	TaskName          string `form:"taskname" query:"taskname" json:"taskname"`
	ReadingMaterialID string `form:"readingmaterialid" query:"updareadingmaterialidtepath" json:"readingmaterialid"`
	VideoLink         string `form:"videolink" query:"videolink" json:"videolink"`
	Marks             string `form:"marks" query:"marks" json:"marks"`
	Description       string `form:"description" query:"description" json:"description"`
	QuestionIDs       string `form:"questionids" query:"questionids" json:"questionids"`
	McqID             string `form:"mcqid" query:"mcqid" json:"mcqid"`
}

func Adduser(u *Users) (*Users, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return u, err

}
func AddCourse(c *Course) (*Course, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err

}
func DeleteCourse(id int) error {
	var tempcollege College
	tempcollege.CollegeID = ""
	if db := db.Find(&tempcollege); db.Error != nil {
		return db.Error
	}
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Delete(&tempcollege).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Commit().Error
	return err

}
func UpdateCourse(course *Course) (*Course, error) {
	var tempCourse Course
	tempCourse.ID = course.ID
	if db := db.Find(&tempCourse); db.Error != nil {
		return nil, db.Error
	}
	if course.CourseName != "" {
		tempCourse.CourseName = course.CourseName
	}
	if course.CourseDescription != "" {
		tempCourse.CourseDescription = course.CourseDescription
	}
	if course.CourseSylabbus != "" {
		tempCourse.CourseSylabbus = course.CourseSylabbus
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Save(&tempCourse).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return &tempCourse, err

}
func GetCourse(id string) (*Course, error) {
	var course Course
	course.CourseID = id
	if db := db.Find(&course); db.Error != nil {
		return nil, db.Error
	}
	return &course, nil

}

func AddCollege(c *College) (*College, error) {

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err
}
func DeleteCollege(id int) error {
	var tempcollege College
	tempcollege.CollegeID = ""
	if db := db.Find(&tempcollege); db.Error != nil {
		return db.Error
	}
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Delete(&tempcollege).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Commit().Error
	return err
}
func UpdateCollege(college *College) (*College, error) {
	var tempCollege College
	tempCollege.ID = college.ID
	if db := db.Find(&tempCollege); db.Error != nil {
		return nil, db.Error
	}
	if college.Name != "" {
		tempCollege.Name = college.Name
	}
	if college.City != "" {
		tempCollege.City = college.City
	}
	if college.State != "" {
		tempCollege.State = college.State
	}
	if college.University != "" {
		tempCollege.University = college.University
	}
	if college.CollegeID != "" {
		tempCollege.CollegeID = college.CollegeID
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Save(&tempCollege).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return &tempCollege, err

}
func GetCollege(id string) (*College, error) {
	var college College
	college.CollegeID = id
	if db := db.Find(&college); db.Error != nil {
		return nil, db.Error
	}
	return &college, nil
}

// curd operation for reading material
func AddReadingMaterial(c *ReadingMaterial) (*ReadingMaterial, error) {

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err

}
func DeleteReadingMaterial(id int) error {
	var tempmaterial ReadingMaterial
	tempmaterial.MaterialID = ""
	if db := db.Find(&tempmaterial); db.Error != nil {
		return db.Error
	}
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Delete(&tempmaterial).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Commit().Error
	return err
}
func UpdateReadingMaterial(material *ReadingMaterial) (*ReadingMaterial, error) {
	var tempmaterial ReadingMaterial
	tempmaterial.MaterialID = material.MaterialID
	if db := db.Find(&tempmaterial); db.Error != nil {
		return nil, db.Error
	}

	if material.Path != "" {
		tempmaterial.Path = material.Path
	}
	if material.UpdatePath != "" {
		tempmaterial.UpdatePath = material.UpdatePath
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Save(&tempmaterial).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return &tempmaterial, err
}
func GetReadingMaterial(id string) (*ReadingMaterial, error) {
	var material ReadingMaterial
	material.MaterialID = id
	if db := db.Find(&material); db.Error != nil {
		return nil, db.Error
	}
	return &material, nil
}

// curd operation for Task

func AddTask(c *Task) (*Task, error) {

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err

}
func DeleteTask(id int) error {
	var temptask Task
	temptask.TaskName = ""
	if db := db.Find(&temptask); db.Error != nil {
		return db.Error
	}
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Delete(&temptask).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Commit().Error
	return err
}
func UpdateTask(task *Task) (*Task, error) {
	var temptask Task
	temptask.TaskName = task.TaskName
	if db := db.Find(&temptask); db.Error != nil {
		return nil, db.Error
	}

	if task.ReadingMaterialID != "" {
		temptask.ReadingMaterialID = task.ReadingMaterialID
	}
	if task.VideoLink != "" {
		temptask.VideoLink = task.VideoLink
	}
	if task.Description != "" {
		temptask.Description = task.Description
	}
	if task.QuestionIDs != "" {
		temptask.QuestionIDs = task.QuestionIDs
	}
	if task.McqID != "" {
		temptask.McqID = task.McqID
	}
	if task.Marks != "" {
		temptask.Marks = task.Marks
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Save(&temptask).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return &temptask, err
}
func GetTask(id string) (*Task, error) {
	var task Task
	task.TaskName = id
	if db := db.Find(&task); db.Error != nil {
		return nil, db.Error
	}
	return &task, nil
}
func Init() {
	fmt.Print("in init")
	var err error
	db, err = gorm.Open("mysql", "root:securepassword@tcp(localhost:3306)/ptech?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Print(err)
	}
	if err == nil {
		fmt.Print("no error")
	}
	db.AutoMigrate(&Users{})

	db.AutoMigrate(&User_profile{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&College{})
	db.AutoMigrate(&ReadingMaterial{})
	db.AutoMigrate(&Task{})

}
func Adduserprofile(c *User_profile) (*User_profile, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err
}
func DeleteuserProfile(id int) error {
	var tempprofile User_profile
	tempprofile.Roll_no = ""
	if db := db.Find(&tempprofile); db.Error != nil {
		return db.Error
	}
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Delete(&tempprofile).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Commit().Error
	return err

}
func UpdateuserProfile(userprofile *User_profile) (*User_profile, error) {

	var tempprofile User_profile
	tempprofile.Roll_no = userprofile.Roll_no
	if db := db.Find(&tempprofile); db.Error != nil {
		return nil, db.Error
	}

	if userprofile.FirstName != "" {
		tempprofile.FirstName = userprofile.FirstName
	}
	if userprofile.LastName != "" {
		tempprofile.LastName = userprofile.LastName
	}
	if userprofile.Dob != "" {
		tempprofile.Dob = userprofile.Dob
	}
	if userprofile.College_id != "" {
		tempprofile.College_id = userprofile.College_id
	}
	if userprofile.Dept != "" {
		tempprofile.Dept = userprofile.Dept
	}
	if userprofile.Phone_no != "" {
		tempprofile.Phone_no = userprofile.Phone_no
	}
	if userprofile.Alt_Phone_no != "" {
		tempprofile.Alt_Phone_no = userprofile.Alt_Phone_no
	}
	if userprofile.Batch_id != "" {
		tempprofile.Batch_id = userprofile.Batch_id
	}
	if userprofile.State != "" {
		tempprofile.State = userprofile.State
	}
	if userprofile.Semester != "" {
		tempprofile.Semester = userprofile.Semester
	}
	if userprofile.Srn_usn != "" {
		tempprofile.Srn_usn = userprofile.Srn_usn
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Save(&tempprofile).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return &tempprofile, err
}
func GetuserProfile(id string) (*User_profile, error) {
	var profile User_profile
	profile.Roll_no = id
	if db := db.Find(&profile); db.Error != nil {
		return nil, db.Error
	}
	return &profile, nil

}

// gygfhgjh
type Account struct {
	gorm.Model
	AccountID string `gorm:"not null;unique;index"`
	Password  uint32
}

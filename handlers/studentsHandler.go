package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/iostrovok/go-convert"

	elst "students_rest_api/elastic/student"
	"students_rest_api/models/response"
	model "students_rest_api/models/student"
)

func ListStudentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := model.SliceToJson(elst.New().ListStudents())
	if err != nil {
		_, _ = w.Write(response.New(1, err.Error(), nil).ToString())
		return
	}
	_, _ = w.Write(response.New(0, "", jsonData).ToString())
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = r.ParseForm()
	id := strings.Split(r.URL.String(), "/")[2] // 0 -> "", 1 -> "students", 2 -> {id}
	if err := elst.New().Id(convert.Int32(id)).DeleteStudent(); err != nil {
		_, _ = w.Write(response.New(1, err.Error(), nil).ToString())
		return
	}
	_, _ = w.Write(response.New(0, "", "true").ToString())
}

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		logrus.Errorf("ParseForm() err: %v", err)
		_, _ = w.Write(response.New(1, err.Error(), nil).ToString())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var student *model.Student
	if err := decoder.Decode(&student); err != nil {
		logrus.Errorf("UpdateStudentHandler() Decode() err: %v", err)
		_, _ = w.Write(response.New(1, err.Error(), nil).ToString())
		return
	}
	if err := elst.New().Body(student).UpdateStudent(); err != nil {
		logrus.Errorf("UpdateStudentHandler() UpdateStudent() err: %v", err)
		_, _ = w.Write(response.New(1, err.Error(), nil).ToString())
		return
	}
	_, _ = w.Write(response.New(0, "", nil).ToString())
}

func CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		logrus.Errorf("ParseForm() err: %v", err)
		_, _ = w.Write(response.New(1, err.Error(), nil).ToString())
		return
	}
	decoder := json.NewDecoder(r.Body)
	var student *model.Student
	if err := decoder.Decode(&student); err != nil {
		logrus.Errorf("CreateStudentHandler() Decode() err: %v", err)
		_, _ = w.Write(response.New(1, err.Error(), nil).ToString())
		return
	}
	if err := elst.New().Body(student).CreateStudent(); err != nil {
		logrus.Errorf("CreateStudentHandler() CreateStudent() err: %v", err)
		_, _ = w.Write(response.New(1, err.Error(), nil).ToString())
		return
	}
	_, _ = w.Write(response.New(0, "", "true").ToString())
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = r.ParseForm()
	id := strings.Split(r.URL.String(), "/")[2] // 0 -> "", 1 -> "students", 2 -> {id}
	if re := regexp.MustCompile("[0-9]+"); re.MatchString(id) {
		jsonData, err := elst.New().Id(convert.Int32(id)).GetStudent().ToJson()
		if err != nil {
			_, _ = w.Write(response.New(1, err.Error(), nil).ToString())
			return
		}
		_, _ = w.Write(response.New(0, "", jsonData).ToString())
		return
	}
	_, _ = w.Write(response.New(1, "Bad id of student, please, try another", nil).ToString())
}

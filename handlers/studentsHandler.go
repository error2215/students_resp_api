package handlers

import (
	"net/http"
	"strings"

	"github.com/iostrovok/go-convert"

	elst "students_rest_api/elastic/student"
	"students_rest_api/models/response"
	model "students_rest_api/models/student"
)

func ListStudentsHandler(w http.ResponseWriter, r *http.Request) {
	json, err := model.SliceToJson(elst.New().ListStudents())
	if err != nil {
		_, _ = w.Write(response.New(1, err.Error(), nil).ToString())
		return
	}
	_, _ = w.Write(response.New(0, "", json).ToString())
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	id := strings.Split(r.URL.String(), "/")[2] // 0 -> "", 1 -> "students", 2 -> {id}
	err := elst.New().Id(convert.Int32(id)).DeleteStudent()
	if err != nil {
		_, _ = w.Write(response.New(1, err.Error(), nil).ToString())
		return
	}
	_, _ = w.Write(response.New(0, "", nil).ToString())
}

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
}

func CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
}

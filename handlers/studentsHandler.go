package handlers

import (
	"net/http"

	elst "students_rest_api/elastic/student"
	"students_rest_api/models/response"
	model "students_rest_api/models/student"
)

func ListStudentsHandler(w http.ResponseWriter, r *http.Request) {
	json, err := model.SliceToJson(elst.New().ListStudents())
	if err != nil {
		_, _ = w.Write(response.New(500, err.Error(), nil).ToString())
	}
	_, _ = w.Write(response.New(0, "", json).ToString())
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
}

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
}

func CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
}

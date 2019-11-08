package main

import (
	"net/http"
	"students_rest_api/handlers"

	"students_rest_api/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/students", func(r chi.Router) {
		r.Get("/", handlers.ListStudentsHandler) // get list of all students

		r.Delete("{/id}", handlers.DeleteStudentHandler) // remove student
		r.Put("/{id}", handlers.CreateStudentHandler)    // add new student
		r.Get("/{id}", handlers.GetStudentHandler)       // get student by id
		r.Post("/{id}", handlers.UpdateStudentHandler)   // update student's data
	})

	r.Route("/lessons", func(r chi.Router) {
		r.Get("/", handlers.ListLessonsHandler) // get list of all lessons

		r.Delete("{/id}", handlers.DeleteLessonHandler) // remove lesson
		r.Put("/{id}", handlers.CreateLessonHandler)    // add new student
		r.Get("/{id}", handlers.GetLessonHandler)       // get lesson by id
		r.Post("/{id}", handlers.UpdateLessonHandler)   // update lessons's data
	})

	log.Info("Application started on port: " + config.GlobalConfig.AppPort)
	err := http.ListenAndServe(":"+config.GlobalConfig.AppPort, r)
	if err != nil {
		log.Info(err)
	}
}

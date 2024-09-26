package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRouters() http.Handler {

	r := chi.NewRouter()

	handler := &TaskHandler{}

	r.Post("/submit-task", handler.SubmitTask)
	r.Get("/task-update", handler.GetTaskUpdate)

	return r

}

package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRouters() http.Handler {

	r := chi.NewRouter()

	r.Post("/submit-task", SubmitTask)
	r.Get("/task-update", GetTaskUpdate)

	return r

}

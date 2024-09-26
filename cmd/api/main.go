package main

import (
	"distributed-task-queue/internal/api"
	"net/http"
)

func main() {

	router := api.RegisterRouters()

	http.ListenAndServe(":8080", router)
}

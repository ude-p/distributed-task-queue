package main

import (
	"distributed-task-queue/internal/api"
	"distributed-task-queue/pkg/configs"
	"distributed-task-queue/pkg/utils"
	"fmt"
	"net/http"
)

func main() {
	cfg := configs.Env.Api
	router := api.RegisterRouters()
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	fmt.Printf("Listening on Address: %s", addr)

	err := http.ListenAndServe(addr, router)
	utils.LogFatalError(err, "Error occured")

}

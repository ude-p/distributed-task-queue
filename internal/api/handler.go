package api

import (
	"context"
	"distributed-task-queue/pkg/configs"
	"distributed-task-queue/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type taskData struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Priority string `json:"priority"`
}

func SubmitTask(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		utils.PrintError(nil, "Content-Type must be application/json")
		return
	}
	var task taskData

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if utils.PrintError(err, "Error occured") {
		return
	}

	task.Id = uuid.NewString()

	rdb, err := configs.InitRedis()
	utils.LogFatalError(err, "Error occured")

	taskJSON, err := json.Marshal(task)
	if utils.PrintError(err, "Error occured") {
		return
	}

	err = rdb.Set(context.Background(), task.Id, taskJSON, 0).Err()
	if utils.PrintError(err, "Error occured") {
		return
	}

}

func GetTaskUpdate(w http.ResponseWriter, r *http.Request) {

}

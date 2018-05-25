package controllers

import (
	"encoding/json"

	"github.com/bzppx/bzppx-codepub/app/models"
)

type TaskLogController struct {
	BaseController
}

// 获取 task_logs
func (this *TaskLogController) GetTaskLogs() {

	taskLogIdStr := this.GetString("taskLog_ids", "")
	if taskLogIdStr == "" {
		this.jsonSuccess("", nil, "")
	}

	var taskLogIds []string
	json.Unmarshal([]byte(taskLogIdStr), &taskLogIds)

	taskLogs, err := models.TaskLogModel.GetTaskLogsByTaskLogIds(taskLogIds)
	if err != nil {
		this.jsonError(err.Error())
	}

	this.jsonSuccess("ok", taskLogs)
}

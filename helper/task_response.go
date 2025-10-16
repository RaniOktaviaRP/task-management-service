package helper

import (
	"task-management/model/domain"
	"task-management/model/web"
)

func ToTaskResponse(task domain.Task) web.TaskResponse {
	return web.TaskResponse{
		Id:             task.Id,
		ProjectId:      task.ProjectId,
		Title:          task.Title,
		Status:         task.Status,
		Priority:       task.Priority,
		Effort:         task.Effort,
		DifficultyLevel: task.DifficultyLevel,
		Deliverable:    task.Deliverable,
		Bottleneck:     task.Bottleneck,
		ContinueTomorrow: task.ContinueTomorrow,
		Progress:       task.Progress,
		CreatedAt:      task.CreatedAt,
		UpdatedAt:      task.UpdatedAt,
	}
}

func ToTaskResponses(tasks []domain.Task) []web.TaskResponse {
	var taskResponses []web.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, ToTaskResponse(task))
	}
	return taskResponses
}
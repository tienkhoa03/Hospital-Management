package utils

import (
	"BE_Hospital_Management/internal/domain/dto"
	"BE_Hospital_Management/internal/domain/entity"
)

func MapToTaskResponse(task *entity.Task, staffUID, managerUID int64) *dto.TaskInfoResponse {
	return &dto.TaskInfoResponse{
		TaskID:      task.Id,
		StaffUID:    staffUID,
		AssignerUID: managerUID,
		Title:       task.Title,
		Description: task.Description,
		BeginTime:   task.BeginTime,
		FinishTime:  task.FinishTime,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

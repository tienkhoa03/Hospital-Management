package task

import (
	"BE_Hospital_Management/internal/domain/entity"
	"time"

	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_task_repository.go

type TaskRepository interface {
	GetDB() *gorm.DB
	CreateTask(tx *gorm.DB, task *entity.Task) (*entity.Task, error)
	GetAllTask() ([]*entity.Task, error)
	GetTaskById(taskId int64) (*entity.Task, error)
	GetTaskByStaffId(staffId int64) (*entity.Task, error)
	GetTasksFromIds(taskIds []int64) ([]*entity.Task, error)
	UpdateTask(tx *gorm.DB, task *entity.Task) (*entity.Task, error)
	ExistsOverlapTaskOfStaff(staffId int64, beginTime, endTime time.Time) (bool, error)
}

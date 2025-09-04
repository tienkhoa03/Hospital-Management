package task

import (
	"BE_Hospital_Management/internal/domain/entity"
	"BE_Hospital_Management/internal/domain/filter"
	"time"

	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_task_repository.go

type TaskRepository interface {
	GetDB() *gorm.DB
	CreateTask(tx *gorm.DB, task *entity.Task) (*entity.Task, error)
	GetAllTask() ([]*entity.Task, error)
	GetTaskById(taskId int64) (*entity.Task, error)
	GetTasksByStaffId(staffId int64) ([]*entity.Task, error)
	GetTasksByManagerId(managerId int64) ([]*entity.Task, error)
	GetTasksByManagerIdAndStaffId(managerId, staffId int64) ([]*entity.Task, error)
	GetTasksFromIds(taskIds []int64) ([]*entity.Task, error)
	UpdateTask(tx *gorm.DB, task *entity.Task) (*entity.Task, error)
	ExistsOverlapTaskOfStaff(staffId int64, beginTime, endTime time.Time) (bool, error)
	DeleteTaskById(tx *gorm.DB, taskId int64) error
	GetTasksByStaffIdWithFilter(staffId int64, taskFilter *filter.TaskFilter) ([]*entity.Task, error)
	GetTasksByManagerIdWithFilter(managerId int64, taskFilter *filter.TaskFilter) ([]*entity.Task, error)
	GetTasksByManagerIdAndStaffIdWithFilter(managerId, staffId int64, taskFilter *filter.TaskFilter) ([]*entity.Task, error)
}

package task

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/internal/domain/entity"
	"BE_Hospital_Management/internal/domain/filter"
	"time"

	"gorm.io/gorm"
)

type PostgreSQLTaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &PostgreSQLTaskRepository{db: db}
}

func (r *PostgreSQLTaskRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLTaskRepository) GetAllTask() ([]*entity.Task, error) {
	var tasks = []*entity.Task{}
	result := r.db.Model(&entity.Task{}).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *PostgreSQLTaskRepository) GetTaskById(taskId int64) (*entity.Task, error) {
	var task = entity.Task{}
	result := r.db.Model(&entity.Task{}).Where("id = ?", taskId).First(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (r *PostgreSQLTaskRepository) GetTasksByStaffId(staffId int64) ([]*entity.Task, error) {
	var tasks []*entity.Task
	result := r.db.Model(&entity.Task{}).Where("staff_id = ?", staffId).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *PostgreSQLTaskRepository) GetTasksByManagerId(managerId int64) ([]*entity.Task, error) {
	var tasks []*entity.Task
	result := r.db.Model(&entity.Task{}).Where("assigner_id = ?", managerId).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *PostgreSQLTaskRepository) GetTasksByManagerIdAndStaffId(managerId, staffId int64) ([]*entity.Task, error) {
	var tasks []*entity.Task
	result := r.db.Model(&entity.Task{}).Where("assigner_id = ? AND staff_id = ?", managerId, staffId).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *PostgreSQLTaskRepository) GetTasksFromIds(taskIds []int64) ([]*entity.Task, error) {
	var tasks []*entity.Task
	result := r.db.Model(&entity.Task{}).Where("id IN ?", taskIds).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *PostgreSQLTaskRepository) CreateTask(tx *gorm.DB, task *entity.Task) (*entity.Task, error) {
	result := tx.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, result.Error
}

func (r *PostgreSQLTaskRepository) UpdateTask(tx *gorm.DB, task *entity.Task) (*entity.Task, error) {
	result := tx.Model(&entity.Task{}).Where("id = ?", task.Id).Updates(task)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedTask = entity.Task{}
	result = tx.First(&updatedTask, task.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedTask, nil
}

func (r *PostgreSQLTaskRepository) ExistsOverlapTaskOfStaff(staffId int64, beginTime, endTime time.Time) (bool, error) {
	var task entity.Task
	err := r.db.Where("staff_id = ?", staffId).Where("finish_time > ? AND begin_time < ? AND status = ?", beginTime, endTime, constant.AppointmentStatusScheduled).Limit(1).Take(&task).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *PostgreSQLTaskRepository) DeleteTaskById(tx *gorm.DB, taskId int64) error {
	result := tx.Model(&entity.Task{}).Where("id = ?", taskId).Update("status", constant.TaskStatusCanceled)
	return result.Error
}

func (r *PostgreSQLTaskRepository) GetTasksByStaffIdWithFilter(staffId int64, taskFilter *filter.TaskFilter) ([]*entity.Task, error) {
	var tasks []*entity.Task
	db := r.db.Model(&entity.Task{}).Where("staff_id = ?", staffId)
	db = taskFilter.ApplyFilter(db)
	result := db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *PostgreSQLTaskRepository) GetTasksByManagerIdWithFilter(managerId int64, taskFilter *filter.TaskFilter) ([]*entity.Task, error) {
	var tasks []*entity.Task
	db := r.db.Model(&entity.Task{}).Where("assigner_id = ?", managerId)
	db = taskFilter.ApplyFilter(db)
	result := db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *PostgreSQLTaskRepository) GetTasksByManagerIdAndStaffIdWithFilter(managerId, staffId int64, taskFilter *filter.TaskFilter) ([]*entity.Task, error) {
	var tasks []*entity.Task
	db := r.db.Model(&entity.Task{}).Where("assigner_id = ? AND staff_id = ?", managerId, staffId)
	db = taskFilter.ApplyFilter(db)
	result := db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

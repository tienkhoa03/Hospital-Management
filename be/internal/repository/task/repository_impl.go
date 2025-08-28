package task

import (
	"BE_Hospital_Management/internal/domain/entity"
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

func (r *PostgreSQLTaskRepository) GetTaskByStaffId(staffId int64) (*entity.Task, error) {
	var task = entity.Task{}
	result := r.db.Model(&entity.Task{}).Where("staff_id = ?", staffId).First(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
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
	err := r.db.Where("staff_id = ?", staffId).Where("end_time > ? AND begin_time < ?", beginTime, endTime).Limit(1).Take(&task).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

package repository

import (
	"fmt"
	"go-rest-api/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	// 指定ユーザーのタスク一覧を取得する
	GetAllTasks(tasks *[]models.Task, userId uint) error

	// 指定したタスクIDのタスク情報を取得する
	GetTaskById(task *models.Task, userId uint, taskId uint) error

	// タスクの新規作成
	CreateTask(task *models.Task) error

	// タスクの更新
	UpdateTask(task *models.Task, userId uint, taskId uint) error

	// タスクの削除
	DeleteTask(userId uint, taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAllTasks(tasks *[]models.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTaskById(task *models.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error; err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) CreateTask(task *models.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) UpdateTask(task *models.Task, userId uint, taskId uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Update("title", task.Title)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}

func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&models.Task{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}

package usecase

import (
	"go-rest-api/models"
	"go-rest-api/repository"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]models.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (models.TaskResponse, error)
	CreateTask(task models.Task) (models.TaskResponse, error)
	UpdateTask(task models.Task, userId uint, taskId uint) (models.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr}
}

func (tu *taskUsecase) GetAllTasks(userId uint) ([]models.TaskResponse, error) {
	tasks := []models.Task{}

	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}

	resTasks := []models.TaskResponse{}
	for _, v := range tasks {
		t := models.TaskResponse{
			Id:        v.Id,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}

	return resTasks, nil
}

func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (models.TaskResponse, error) {
	task := models.Task{}

	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return models.TaskResponse{}, err
	}

	resTask := models.TaskResponse{
		Id:        task.Id,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (tu *taskUsecase) CreateTask(task models.Task) (models.TaskResponse, error) {
	if err := tu.tr.CreateTask(&task); err != nil {
		return models.TaskResponse{}, err
	}

	resTask := models.TaskResponse{
		Id:        task.Id,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (tu *taskUsecase) UpdateTask(task models.Task, userId uint, taskId uint) (models.TaskResponse, error) {
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return models.TaskResponse{}, err
	}

	resTask := models.TaskResponse{
		Id:        task.Id,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}

	return nil
}

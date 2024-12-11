package repository

import "github.com/clrajapaksha/to-do-list-app/entities"

type TaskRepository interface {
	Save(task *entities.Task) (*entities.Task, error)
	FindAll() ([]entities.Task, error)
	FindByID(id string) (*entities.Task, error)
	Delete(task *entities.Task) error
}

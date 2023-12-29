package repository

import (
	"github.com/keima483/calender-app/models"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Task   string
	Date   string
	UserId int
}

func TaskFromModel(model models.TaskModel) Task {
	return Task{
		Task: model.Task,
		Date: model.Date,
	}
}

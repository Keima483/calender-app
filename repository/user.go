package repository

import (
	"github.com/keima483/calender-app/models"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string  `json:"-"`
	Tasks    []*Task `gorm:"foreignKey:UserId"`
}

func (u *User) AddTasks(tasks []*Task)  {
	newTasks := tasks
	if u.Tasks != nil {
		newTasks = append(u.Tasks, newTasks...)
	}
	u.Tasks = newTasks
}

func (u *User) AddTask(task *Task) {
	newTask := []*Task {task}
	if u.Tasks != nil {
		newTask = append(newTask, u.Tasks...)
	}
	u.Tasks = newTask
}

func UserFromModel(model models.UserModel) User {
	return User{
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
	}
}

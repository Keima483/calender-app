package services

import (
	"errors"

	"github.com/keima483/calender-app/initializers"
	"github.com/keima483/calender-app/models"
	"github.com/keima483/calender-app/repository"
)

func GetHolidays() []repository.Holiday {
	var holidays []repository.Holiday
	initializers.DB.Find(&holidays)
	return holidays
}

func getTasks(userId int) []repository.Task {
	var tasks []repository.Task
	initializers.DB.Where("user_id = ?", userId).Find(&tasks)
	return tasks
}

func GetTasks(userId int) ([]repository.Task, error) {
	var user repository.User
	initializers.DB.First(&user, userId)
	if user.Email == "" {
		return []repository.Task{}, errors.New("no user with such id")
	}
	return getTasks(userId), nil
} 

func AddHolidays(holidays []*models.HolidayModel)  []*repository.Holiday {
	var hol []*repository.Holiday
	for i := 0; i < len(holidays); i++ {
		hol = append(hol, repository.HolidayFromModel(*holidays[i]))
	}
	initializers.DB.Create(&hol)
	return hol
}

func AddTask(task models.TaskModel, userId int) ([]repository.Task, error) {
	user := new(repository.User)
	initializers.DB.First(&user, userId)
	if user.Email == "" {
		return []repository.Task{}, errors.New("no User with this id")
	}
	tsk := repository.TaskFromModel(task)
	user.AddTask(&tsk)
	initializers.DB.Save(&user)
	return getTasks(userId), nil
}

func AddTasks(tasks []models.TaskModel, userId int) ([]repository.Task, error) {
	user := new(repository.User)
	initializers.DB.First(&user, userId)
	if user.Email == "" {
		return []repository.Task{}, errors.New("no User with this id")
	}
	var tsks []*repository.Task
	for i := 0; i < len(tasks); i++ {
		task := repository.TaskFromModel(tasks[i])
		tsks = append(tsks, &task)
	}
	user.AddTasks(tsks)
	initializers.DB.Save(&user)
	return getTasks(userId), nil
}

func GetDayDetail(date string, userId int) ([]repository.Task, []repository.Holiday, error) {
	var user repository.User
	initializers.DB.First(&user, userId)
	if user.Email == "" {
		return []repository.Task{}, []repository.Holiday{} , errors.New("no user with such id")
	}
	var tasks []repository.Task
	initializers.DB.Where("user_id = ? and date = ?", userId, date).Find(&tasks)
	var holidays []repository.Holiday
	initializers.DB.Where("date = ?", date).Find(&holidays)
	return tasks, holidays, nil
}

func DeleteTask(taskId int) error {
	var task repository.Task
	initializers.DB.Find(&task, taskId)
	if task.Task == "" {
		return errors.New("no task with this id")
	}
	initializers.DB.Delete(&task)
	return nil
}
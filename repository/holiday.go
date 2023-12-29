package repository

import (
	"github.com/keima483/calender-app/models"
	"gorm.io/gorm"
)

type Holiday struct {
	gorm.Model
	HolidayName string
	Date        string
}

func HolidayFromModel(model models.HolidayModel) *Holiday {
	return &Holiday{
		HolidayName: model.HolidayName,
		Date:        model.Date,
	}
}

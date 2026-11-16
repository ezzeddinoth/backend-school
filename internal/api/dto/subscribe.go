package dto

import "school/internal/database/model"

type Subscribe struct {
	ID        *int   `json:"id" binding:"required"`
	StudentID *int   `json:"studentID" binding:"required"`
	MenuID    *int   `json:"menuID" binding:"required"`
	Day       string `json:"day" binding:"required"`
}

type PostSubscribe struct {
	StudentID *int   `json:"studentID" binding:"required"`
	MenuID    *int   `json:"menuID" binding:"required"`
	Day       string `json:"day" binding:"required"`
}

func SubscribeModelGetMapper(m model.Subscribe) Subscribe {
	return Subscribe{
		ID:        m.ID,
		StudentID: m.StudentID,
		MenuID:    m.MenuID,
		Day:       m.Day,
	}
}

func SubscribeModelModelMapper(d PostSubscribe) model.Subscribe {
	return model.Subscribe{
		StudentID: d.StudentID,
		MenuID:    d.MenuID,
		Day:       d.Day,
	}
}

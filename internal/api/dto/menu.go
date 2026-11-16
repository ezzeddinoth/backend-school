package dto

import "school/internal/database/model"

type Menu struct {
	ID   *int   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type PostMenu struct {
	Name string `json:"name" binding:"required"`
}

func MenuModelGetMapper(model model.Menu) Menu {
	return Menu{
		ID:   model.ID,
		Name: model.Name,
	}
}

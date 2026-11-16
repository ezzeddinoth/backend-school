package repository

import (
	"school/internal/database/model"

	"gorm.io/gorm"
)

type Repository struct {
	Student  BaseRepository[model.Student]
	Menu     BaseRepository[model.Menu]
	Subscribe BaseRepository[model.Subscribe]
}

func SetupRepositories(db *gorm.DB) *Repository {
	return &Repository{
		Student:  newBaseRepository[model.Student](db),
		Menu:     newBaseRepository[model.Menu](db),
		Subscribe: newBaseRepository[model.Subscribe](db),
	}
}

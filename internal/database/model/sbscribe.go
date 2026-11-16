package model

import "os"

type Subscribe struct {
	ID        *int   `gorm:"column:id;primaryKey;autoIncrement"`
	StudentID *int   `gorm:"column:student_id;not null;uniqueIndex:idx_subscribe_unique"`
	MenuID    *int   `gorm:"column:menu_id;not null;uniqueIndex:idx_subscribe_unique"`
	Day       string `gorm:"column:day;type:varchar(50);not null;uniqueIndex:idx_subscribe_unique"`

	Student Student `gorm:"foreignKey:StudentID;references:ID"`
	Menu    Menu    `gorm:"foreignKey:MenuID;references:ID"`
}

func (Subscribe) TableName() string {
	if os.Getenv("GO_TEST_MODE") == "true" {
		return "subscribe"
	}
	return "subscribe"
}

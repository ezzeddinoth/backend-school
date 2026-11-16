package model

import "os"

type Student struct {
    ID   *int  `gorm:"primaryKey;autoIncrement"`
    Name string `gorm:"type:varchar(255);not null"`

    Subscriptions []Subscribe `gorm:"foreignKey:StudentID"`
}


func (Student) TableName() string {
	if os.Getenv("GO_TEST_MODE") == "true" {
		return "student"
	}
	return "student"
}

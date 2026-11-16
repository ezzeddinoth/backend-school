package model

import "os"


type Menu struct {
    ID   *int  `gorm:"primaryKey;autoIncrement"`
    Name string `gorm:"type:varchar(255);not null"`

    Subscriptions []Subscribe `gorm:"foreignKey:MenuID"`
}

func (Menu) TableName() string {
	if os.Getenv("GO_TEST_MODE") == "true" {
		return "menu"
	}
	return "menu"
}

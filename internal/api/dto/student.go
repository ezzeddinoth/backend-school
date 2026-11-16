package dto

import "school/internal/database/model"

type Student struct {
	ID   *int   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`

	Subscriptions []StudentSubscrip `json:"subscription" `
}
type StudentSubscrip struct {
	Day  string `json:"day" binding:"required"`
	Menu string `json:"menu" binding:"required"`
}

type PostStudent struct {
	Name string `json:"name" binding:"required"`
}

func StudentModelGetMapper(model model.Student) Student {
	subscrips := make([]StudentSubscrip, 0, len(model.Subscriptions))
	for _, s := range model.Subscriptions {
		subscrips = append(subscrips, StudentSubscrip{
			Day:  s.Day,
			Menu: s.Menu.Name,
		})
	}
	return Student{
		ID:            model.ID,
		Name:          model.Name,
		Subscriptions: subscrips,
	}
}

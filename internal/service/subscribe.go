package service

import (
	"school/internal/api/dto"
	"school/internal/database"
	"school/internal/database/model"
	"school/internal/database/repository"
)

type SubscribeService interface {
	BulkCreate(subscribes []dto.PostSubscribe) ([]dto.Subscribe, error)
}

type subscribeService struct {
	subscribeRepository repository.BaseRepository[model.Subscribe]
}

func NewSubscribeService() SubscribeService {
	return &subscribeService{
		subscribeRepository: database.Repository.Subscribe,
	}
}

func (s *subscribeService) BulkCreate(dtos []dto.PostSubscribe) ([]dto.Subscribe, error) {
	subscribes := make([]model.Subscribe, 0, len(dtos))
	for _, d := range dtos {
		subscribes = append(subscribes, dto.SubscribeModelModelMapper(d))
	}

	err := s.subscribeRepository.BulkCreate(subscribes)
	if err != nil {
		return nil, err
	}

	response := make([]dto.Subscribe, 0, len(subscribes))
	for _, s := range subscribes {
		response = append(response, dto.SubscribeModelGetMapper(s))
	}

	return response, nil
}

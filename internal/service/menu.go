package service

import (
	"school/internal/api/dto"
	"school/internal/database"
	"school/internal/database/model"
	"school/internal/database/repository"
)

type MenuService interface {
	GetAll() ([]dto.Menu, error)
	Create(menu dto.PostMenu) (dto.Menu, error)
	Update(menu dto.Menu) (dto.Menu, error)
	Delete(id int) error
}

type menuService struct {
	menuRepository repository.BaseRepository[model.Menu]
}

func NewMenuService() MenuService {
	return &menuService{
		menuRepository: database.Repository.Menu,
	}
}

func (s *menuService) GetAll() ([]dto.Menu, error) {
	menus, err := s.menuRepository.GetAll()
	if err != nil {
		return nil, err
	}

	response := []dto.Menu{}
	for _, s := range menus {
		response = append(response, dto.MenuModelGetMapper(s))
	}

	return response, nil
}

func (s *menuService) Create(menu dto.PostMenu) (dto.Menu, error) {
	modelMenu := model.Menu{
		Name: menu.Name,
	}
	err := s.menuRepository.Create(&modelMenu)
	if err != nil {
		return dto.Menu{}, err
	}

	return dto.MenuModelGetMapper(modelMenu), nil
}

func (s *menuService) Update(menu dto.Menu) (dto.Menu, error) {
	_, err := s.menuRepository.FindMustExistId(*menu.ID)
	if err != nil {
		return dto.Menu{}, err
	}

	modelMenu := model.Menu{
		ID:   menu.ID,
		Name: menu.Name,
	}

	err = s.menuRepository.Update(&modelMenu)
	if err != nil {
		return dto.Menu{}, err
	}

	return dto.MenuModelGetMapper(modelMenu), nil
}

func (s *menuService) Delete(id int) error {
	_, err := s.menuRepository.FindMustExistId(id)
	if err != nil {
		return err
	}

	return s.menuRepository.Delete(id)
}

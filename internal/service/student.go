package service

import (
	"cmp"
	"school/internal/api/dto"
	"school/internal/database"
	"school/internal/database/model"
	"school/internal/database/repository"
	"slices"
)

type StudentService interface {
	GetAll() ([]dto.Student, error)
	Create(student dto.PostStudent) (dto.Student, error)
	Update(student dto.Student) (dto.Student, error)
	Delete(id int) error
}

type studentService struct {
	studentRepository repository.BaseRepository[model.Student]
}

func NewStudentService() StudentService {
	return &studentService{
		studentRepository: database.Repository.Student,
	}
}

func (s *studentService) GetAll() ([]dto.Student, error) {
	students, err := s.studentRepository.GetAll("Subscriptions.Menu")
	if err != nil {
		return nil, err
	}

	response := []dto.Student{}
	for _, s := range students {
		response = append(response, dto.StudentModelGetMapper(s))
	}

	slices.SortFunc(response, func(a, b dto.Student) int {
		return cmp.Compare(*a.ID, *b.ID)
	})
	return response, nil
}

func (s *studentService) Create(student dto.PostStudent) (dto.Student, error) {
	modelStudent := model.Student{
		Name: student.Name,
	}
	err := s.studentRepository.Create(&modelStudent)
	if err != nil {
		return dto.Student{}, err
	}

	return dto.StudentModelGetMapper(modelStudent), nil
}

func (s *studentService) Update(student dto.Student) (dto.Student, error) {
	_, err := s.studentRepository.FindMustExistId(*student.ID)
	if err != nil {
		return dto.Student{}, err
	}

	modelStudent := model.Student{
		ID:   student.ID,
		Name: student.Name,
	}

	err = s.studentRepository.Update(&modelStudent)
	if err != nil {
		return dto.Student{}, err
	}

	return dto.StudentModelGetMapper(modelStudent), nil
}

func (s *studentService) Delete(id int) error {
	_, err := s.studentRepository.FindMustExistId(id)
	if err != nil {
		return err
	}

	return s.studentRepository.Delete(id)
}

package handler

import (
	"school/internal/api/dto"
	"school/internal/service"

	"github.com/stretchr/testify/mock"
)

type MockStudentService struct {
	service.StudentService
	mock.Mock
}

func (m *MockStudentService) GetAll() ([]dto.Student, error) {
	args := m.Called()
	return args.Get(0).([]dto.Student), args.Error(1)
}

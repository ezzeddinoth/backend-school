package service

import (
	"school/internal/database/model"
	"school/internal/database/repository"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStudentRepository struct {
	repository.BaseRepository[model.Student]
	mock.Mock
}

func (m *MockStudentRepository) GetAll(sts ...string) ([]model.Student, error) {
	args := m.Called()
	return args.Get(0).([]model.Student), args.Error(1)
}

func TestStudentService_GetAll(t *testing.T) {
	repo := new(MockStudentRepository)
	svc := &studentService{
		studentRepository: repo,
	}
	repo.On("GetAll").Return([]model.Student{}, nil)

	result, err := svc.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

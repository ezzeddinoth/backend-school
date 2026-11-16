package repository

import (
	"school/internal/exception"
	"school/internal/utils"

	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	Create(entity *T) error
	BulkCreate(entities []T) error
	GetAll(preloads ...string) ([]T, error)
	GetByID(id int, preloads ...string) (*T, error)
	GetByIDs(ids []int) ([]T, error)
	GetAllExceptIDs(ids []int) ([]T, error)
	Update(entity *T) error
	BulkUpdate(entities []T) error
	Delete(id int) error
	BulkDelete(ids []int) error
	Count() (int64, error)
	FindMustExistIds(ids []int) ([]T, error)
	getDB() *gorm.DB
	FindMustExistId(id int) (*T, error)
}

// BaseRepository provides common CRUD operations
type baseRepository[T any] struct {
	db *gorm.DB
}

// NewBaseRepository initializes the base repository
func newBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db: db}
}

// Create creates a new record in the database
func (r *baseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

// BulkCreate creates a new records in the database
func (r *baseRepository[T]) BulkCreate(entity []T) error {
	return r.db.Create(entity).Error
}

// GetAll retrieves all records with optional preloading
func (r *baseRepository[T]) GetAll(preloads ...string) ([]T, error) {
	var entities []T
	query := r.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	err := query.Find(&entities).Error
	return entities, err
}

// GetByID retrieves a single record by its ID with optional preloading
func (r *baseRepository[T]) GetByID(id int, preloads ...string) (*T, error) {
	var entity T
	query := r.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	err := query.First(&entity, id).Error
	return &entity, err
}

func (r *baseRepository[T]) GetByIDs(ids []int) ([]T, error) {
	var entity []T

	err := r.db.Find(&entity, ids).Error
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *baseRepository[T]) GetAllExceptIDs(ids []int) ([]T, error) {
	var entity []T

	err := r.db.
		Where("id NOT IN ?", ids).
		Find(&entity).Error

	if err != nil {
		return nil, err
	}

	return entity, nil
}

// Update updates an existing record in the database
func (r *baseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

// BulkUpdate updates existing records in the database. If value doesn't contain a matching primary key, value is inserted
func (r *baseRepository[T]) BulkUpdate(entity []T) error {
	return r.db.Save(entity).Error
}

// Delete deletes a record by its ID
func (r *baseRepository[T]) Delete(id int) error {
	entity := new(T)
	return r.db.Delete(entity, id).Error
}

// BulkDelete deletes records by its IDs
func (r *baseRepository[T]) BulkDelete(ids []int) error {
	entity := new(T)
	return r.db.Delete(entity, ids).Error
}

// Count returns the total count of records in the database
func (r *baseRepository[T]) Count() (int64, error) {
	var count int64
	entity := new(T)
	err := r.db.Model(entity).Count(&count).Error
	return count, err
}

func (r *baseRepository[T]) FindMustExistIds(ids []int) ([]T, error) {
	var entity []T

	if len(ids) < 1 {
		return entity, nil
	}
	s := utils.UniqueIntSlice(ids)

	err := r.db.Find(&entity, s).Error
	if err != nil {
		return nil, err
	}

	if len(s) > len(entity) {
		return nil, &exception.NotFound{Message: "one or more ids not found"}
	}

	return entity, nil
}

func (r *baseRepository[T]) FindMustExistId(id int) (*T, error) {
	var entity []T

	err := r.db.Find(&entity, id).Error
	if err != nil {
		return nil, err
	}

	if len(entity) < 1 {
		return nil, &exception.NotFound{Message: "id not found"}
	}

	return &entity[0], nil
}

func (r *baseRepository[T]) getDB() *gorm.DB {
	return r.db
}

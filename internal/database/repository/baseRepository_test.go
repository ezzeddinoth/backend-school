package repository

import (
	"os"
	"testing"

	"github.com/go-playground/assert/v2"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestEntity represents an entity used in tests
type TestEntity struct {
	ID   int    `gorm:"primaryKey;autoIncrement;column:id"`
	Name string `gorm:"column:name"`
}

func (TestEntity) GetColumns() []string {
	return []string{"id", "name"}
}

// SetupTestRepository sets up an in-memory SQLite database and returns a new BaseRepository instance
func SetupTestRepository[T any](t *testing.T) BaseRepository[T] {
	os.Setenv("GO_TEST_MODE", "true")
	// Set up in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate schema (create tables)
	err = db.AutoMigrate(&TestEntity{})
	if err != nil {
		t.Fatalf("Failed to migrate schema: %v", err)
	}

	// Initialize and return the repository with the database connection
	return newBaseRepository[T](db)
}

func TestBaseRepository_Create(t *testing.T) {
	repo := SetupTestRepository[TestEntity](t)

	// Create a new entity
	entity := &TestEntity{Name: "Test Name"}
	err := repo.Create(entity)

	// Assertions
	assert.Equal(t, err, nil)
	assert.NotEqual(t, entity.ID, 0) // Ensure that ID is populated after creation
}

func TestBaseRepository_GetAll(t *testing.T) {
	// Use SetupTestRepository to avoid redundancy
	repo := SetupTestRepository[TestEntity](t)

	// Add a couple of records
	repo.Create(&TestEntity{Name: "Test Name 1"})
	repo.Create(&TestEntity{Name: "Test Name 2"})

	// Retrieve all records
	entities, err := repo.GetAll()

	// Assertions
	assert.Equal(t, err, nil)
	assert.Equal(t, len(entities), 2)
}

func TestBaseRepository_GetByID(t *testing.T) {
	// Use SetupTestRepository to avoid redundancy
	repo := SetupTestRepository[TestEntity](t)

	// Create a record
	entity := &TestEntity{Name: "Test Name"}
	repo.Create(entity)

	// Retrieve the record by ID
	result, err := repo.GetByID(entity.ID)

	// Assertions
	assert.Equal(t, err, nil)
	assert.Equal(t, result.ID, entity.ID)
	assert.Equal(t, result.Name, entity.Name)
}

func TestBaseRepository_Update(t *testing.T) {
	// Use SetupTestRepository to avoid redundancy
	repo := SetupTestRepository[TestEntity](t)

	// Create a record
	entity := &TestEntity{Name: "Old Name"}
	repo.Create(entity)

	// Update the record
	entity.Name = "Updated Name"
	err := repo.Update(entity)

	// Assertions
	assert.Equal(t, err, nil)

	// Retrieve the updated record
	updatedEntity, _ := repo.GetByID(entity.ID)
	assert.Equal(t, updatedEntity.Name, "Updated Name")
}

func TestBaseRepository_Delete(t *testing.T) {
	// Use SetupTestRepository to avoid redundancy
	repo := SetupTestRepository[TestEntity](t)

	// Create a record
	entity := &TestEntity{Name: "Test Name"}
	repo.Create(entity)

	// Delete the record
	err := repo.Delete(entity.ID)

	// Assertions
	assert.Equal(t, err, nil)

	// Verify deletion
	_, err = repo.GetByID(entity.ID)
	assert.NotEqual(t, err, nil) // Error should occur because the entity is deleted
}

package test

import (
	"context"
	"database/sql"
	"golang-restful-api/internal/config"
	"golang-restful-api/internal/domain"
	"golang-restful-api/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	return config.NewDB()
}

func TestCreateCategory(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	categoryRepository := repository.NewCategoryRepository(db)

	ctx := context.Background()
	category := domain.Category{
		Name: "Test Category",
	}

	result := categoryRepository.Create(ctx, category)

	assert.NotEqual(t, 0, result.Id)
	assert.Equal(t, "Test Category", result.Name)

	// Cleanup
	categoryRepository.Delete(ctx, result.Id)
}

func TestFindByIdCategory(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	categoryRepository := repository.NewCategoryRepository(db)

	ctx := context.Background()
	category := domain.Category{
		Name: "Test Find Category",
	}

	created := categoryRepository.Create(ctx, category)
	result, err := categoryRepository.FindById(ctx, created.Id)

	assert.Nil(t, err)
	assert.Equal(t, created.Id, result.Id)
	assert.Equal(t, "Test Find Category", result.Name)

	// Cleanup
	categoryRepository.Delete(ctx, created.Id)
}

func TestFindAllCategory(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	categoryRepository := repository.NewCategoryRepository(db)

	ctx := context.Background()
	results := categoryRepository.FindAll(ctx)

	assert.NotNil(t, results)
}

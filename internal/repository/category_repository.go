package repository

import (
    "context"
    "database/sql"
    "errors"
    "golang-restful-api/internal/domain"
)

type CategoryRepositoryImpl struct {
    DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) domain.CategoryRepository {
    return &CategoryRepositoryImpl{DB: db}
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context, category domain.Category) domain.Category {
    query := "INSERT INTO categories (name) VALUES (?)"
    result, err := repository.DB.ExecContext(ctx, query, category.Name)
    if err != nil {
        panic(err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        panic(err)
    }

    category.Id = id
    return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, category domain.Category) domain.Category {
    query := "UPDATE categories SET name = ? WHERE id = ?"
    _, err := repository.DB.ExecContext(ctx, query, category.Name, category.Id)
    if err != nil {
        panic(err)
    }

    return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, categoryId int64) {
    query := "DELETE FROM categories WHERE id = ?"
    _, err := repository.DB.ExecContext(ctx, query, categoryId)
    if err != nil {
        panic(err)
    }
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, categoryId int64) (domain.Category, error) {
    query := "SELECT id, name, created_at, updated_at FROM categories WHERE id = ?"
    row := repository.DB.QueryRowContext(ctx, query, categoryId)

    var category domain.Category
    err := row.Scan(&category.Id, &category.Name, &category.CreateAt, &category.UpdateAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return category, errors.New("category not found")
        }
        panic(err)
    }

    return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context) []domain.Category {
    query := "SELECT id, name, created_at, updated_at FROM categories"
    rows, err := repository.DB.QueryContext(ctx, query)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    var categories []domain.Category
    for rows.Next() {
        var category domain.Category
        err := rows.Scan(&category.Id, &category.Name, &category.CreateAt, &category.UpdateAt)
        if err != nil {
            panic(err)
        }
        categories = append(categories, category)
    }

    return categories
}
package domain

import (
	"time"
	"context"
)

type Category struct {
	Id  	int64    `json:"id"`
	Name 	string   `json:"name"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
}

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type CategoryUpdateRequest struct {
	Id   int64 `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type CategoryResponse struct {
	Id  	int64    `json:"id"`
	Name    string   `json:"name"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type CategoryRepository interface {
	Create(ctx context.Context, category Category) Category
	Update(ctx context.Context, category Category) Category
	Delete(ctx context.Context, categoryId int64)
	FindById(ctx context.Context, categoryId int64) (Category, error)
	FindAll(ctx context.Context) []Category
}

type CategoryService interface {
	Create(ctx context.Context, request CategoryCreateRequest) CategoryResponse
	Update(ctx context.Context, request CategoryUpdateRequest) CategoryResponse
	Delete(ctx context.Context, categoryId int64)
	FindById(ctx context.Context, categoryId int64) CategoryResponse
	FindAll(ctx context.Context) []CategoryResponse
}
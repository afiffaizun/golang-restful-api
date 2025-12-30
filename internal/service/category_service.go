package service

import (
	"context"
	"golang-restful-api/internal/domain"
	"golang-restful-api/pkg/helper"
)

type CategoryServiceImpl struct {
	CategoryRepository domain.CategoryRepository
}

func NewCategoryService(categoryRepository domain.CategoryRepository) domain.CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request domain.CategoryCreateRequest) domain.CategoryResponse {

	err := helper.ValidateStruct(request)
	if len(err) > 0 {
		panic(err)
	}

	//Create Category
	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Create(ctx, category)

	return toCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request domain.CategoryUpdateRequest) domain.CategoryResponse {

	//Validate request
	err := helper.ValidateStruct(request)
	if len(err) > 0 {
		panic(err)
	}

	//Check Category exist
	category, err2 := service.CategoryRepository.FindById(ctx, request.Id)
	if err2 != nil {
		panic(err2)
	}

	//Update Category
	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, category)

	return toCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int64) {

	//Check Category exist
	_, err := service.CategoryRepository.FindById(ctx, categoryId)
	if err != nil {
		panic(err)
	}

	//Delete Category
	service.CategoryRepository.Delete(ctx, categoryId)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int64) domain.CategoryResponse {

	category, err := service.CategoryRepository.FindById(ctx, categoryId)
	if err != nil {
		panic(err)
	}

	return toCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []domain.CategoryResponse {
	categories := service.CategoryRepository.FindAll(ctx)

	var categoryResponses []domain.CategoryResponse

	for _, category := range categories {
		categoryResponses = append(categoryResponses, toCategoryResponse(category))
	}

	return categoryResponses
}

func toCategoryResponse(category domain.Category) domain.CategoryResponse {
	return domain.CategoryResponse{
		Id:       category.Id,
		Name:     category.Name,
		CreateAt: category.CreateAt,
		UpdateAt: category.UpdateAt,
	}
}

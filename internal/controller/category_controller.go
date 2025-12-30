package controller

import (
	"encoding/json"
	"golang-restful-api/internal/domain"
	"golang-restful-api/pkg/helper"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CategoryController struct {
	CategoryService domain.CategoryService
}

func NewCategoryController(categoryService domain.CategoryService) *CategoryController {
	return &CategoryController{
		CategoryService: categoryService,
	}
}

func (controller *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	var request domain.CategoryCreateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.Create(r.Context(), request)

	webResponse := helper.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   categoryResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(webResponse)
}

func (controller *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	var request domain.CategoryUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.PanicIfError(err)

	categoryId := chi.URLParam(r, "categoryId")
	id, err := strconv.ParseInt(categoryId, 10, 64)
	helper.PanicIfError(err)

	request.Id = id

	categoryResponse := controller.CategoryService.Update(r.Context(), request)
	webResponse := helper.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(webResponse)
}

func (controller *CategoryController) FindById(w http.ResponseWriter, r *http.Request) {
	categoryId := chi.URLParam(r, "categoryId")
	id, err := strconv.ParseInt(categoryId, 10, 64)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(r.Context(), id)
	webResponse := helper.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(webResponse)
}

func (controller *CategoryController) FindAll(w http.ResponseWriter, r *http.Request) {
	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := helper.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponses,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(webResponse)
}

func (controller *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	categoryId := chi.URLParam(r, "categoryId")
	id, err := strconv.ParseInt(categoryId, 10, 64)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), id)
	webResponse := helper.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(webResponse)
}

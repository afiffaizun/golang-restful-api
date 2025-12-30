package router

import (
    "golang-restful-api/internal/controller"
    "golang-restful-api/internal/middleware"
    "golang-restful-api/pkg/helper"
    "net/http"

    "github.com/go-chi/chi/v5"
)

func NewRouter(categoryController *controller.CategoryController) http.Handler {
    r := chi.NewRouter()

    // Middleware untuk panic recovery
    r.Use(func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            defer func() {
                if err := recover(); err != nil {
                    helper.ErrorHandler(w, r, err)
                }
            }()
            next.ServeHTTP(w, r)
        })
    })

    // Public routes
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to Golang RESTful API"))
    })

    // Protected routes
    r.Route("/api/categories", func(r chi.Router) {
        r.Use(middleware.AuthMiddleware)
        r.Get("/", categoryController.FindAll)
        r.Get("/{categoryId}", categoryController.FindById)
        r.Post("/", categoryController.Create)
        r.Put("/{categoryId}", categoryController.Update)
        r.Delete("/{categoryId}", categoryController.Delete)
    })

    return r
}
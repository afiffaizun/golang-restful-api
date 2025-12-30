package main

import (
    "fmt"
    "golang-restful-api/internal/config"
    "golang-restful-api/internal/controller"
    "golang-restful-api/internal/repository"
    "golang-restful-api/internal/service"
    "golang-restful-api/pkg/router"
    "net/http"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Warning: .env file not found")
    } else {
        fmt.Println("Environment variables loaded successfully")
        // Debug: print DB config (jangan print password di production!)
        fmt.Printf("DB_HOST: %s, DB_PORT: %s, DB_NAME: %s\n", 
            os.Getenv("DB_HOST"), 
            os.Getenv("DB_PORT"), 
            os.Getenv("DB_NAME"))
    }

    // Database connection
    db := config.NewDB()
    defer db.Close()

    // Initialize layers
    categoryRepository := repository.NewCategoryRepository(db)
    categoryService := service.NewCategoryService(categoryRepository)
    categoryController := controller.NewCategoryController(categoryService)

    // Setup router
    r := router.NewRouter(categoryController)

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Printf("Server running on port %s\n", port)
    err = http.ListenAndServe(":"+port, r)
    if err != nil {
        panic(err)
    }
}
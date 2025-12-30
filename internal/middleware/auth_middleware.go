package middleware

import (
    "encoding/json"
    "golang-restful-api/pkg/helper"
    "net/http"
    "os"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        apiKey := r.Header.Get("X-API-Key")
        expectedKey := os.Getenv("API_KEY")

        if apiKey != expectedKey {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusUnauthorized)

            webResponse := helper.WebResponse{
                Code:   http.StatusUnauthorized,
                Status: "Unauthorized",
                Errors: "Invalid or missing API key",
            }

            json.NewEncoder(w).Encode(webResponse)
            return
        }

        next.ServeHTTP(w, r)
    })
}
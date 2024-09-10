
package middleware

import (
    "context"
    "net/http"
    "strings"
    "your_project_path/models"

    "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key") // Replace with your secret key

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
            return
        }

        tokenStr := parts[1]
        claims := &Claims{}

        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Attach user ID to context
        ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

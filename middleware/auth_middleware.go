package middleware

import (
    "net/http"
    "strings"
    "github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }
        
        tokenString := strings.Split(authHeader, " ")[1]
        token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("your_secret_key"), nil
        })
        
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            r.Header.Set("email", claims["email"].(string))
            next.ServeHTTP(w, r)
        } else {
            http.Error(w, "Forbidden", http.StatusForbidden)
        }
    })
}

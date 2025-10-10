package middleware

import (
	"net/http"
	"strings"
	"task-management/helper"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

var jwtSecret = []byte("rahasia")

// Middleware untuk http.Handler
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helper.WriteUnauthorized(w, "Missing Authorization header")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			helper.WriteUnauthorized(w, "Invalid Authorization format")
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			helper.WriteUnauthorized(w, "Invalid or expired token")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Middleware untuk httprouter.Handle
func JWTAuthHttprouter(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helper.WriteUnauthorized(w, "Missing Authorization header")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			helper.WriteUnauthorized(w, "Invalid Authorization format")
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			helper.WriteUnauthorized(w, "Invalid or expired token")
			return
		}

		h(w, r, ps)
	}
}

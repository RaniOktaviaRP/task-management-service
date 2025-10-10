package controller

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"task-management/helper"
	"task-management/model/web"
)

var jwtSecret = []byte("rahasia")

// LoginHandler untuk autentikasi user dengan email & password lalu menghasilkan JWT token
// @Summary Login user
// @Tags auth
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 200 {object} web.WebResponse
// @Failure 401 {object} web.WebResponse
// @Router /login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Dummy check (sementara, sebelum DB)
	if email != "admin@mail.com" || password != "123" {
		helper.WriteUnauthorized(w, "Email/password salah")
		return
	}

	// Buat JWT (tanpa role)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // expired 1 jam
	})

	tokenString, err := token.SignedString(jwtSecret)
	helper.PanicIfError(err)

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data: map[string]string{
			"token": tokenString,
		},
	}

	helper.WriteToResponseBody(w, response)
}

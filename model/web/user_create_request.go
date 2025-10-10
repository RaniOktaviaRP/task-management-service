package web

type UserRegisterRequest struct {
    FullName string `json:"full_name" validate:"required"` // tambahkan ini
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
    Role     string `json:"role" validate:"required,oneof=SE SCE"`
}

type UserLoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

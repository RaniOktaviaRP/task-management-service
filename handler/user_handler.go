package handler

import (
    "encoding/json"
    "net/http"
    "task-management/service"
)

type UserHandler struct {
    UserService *service.UserServiceImpl
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    json.NewDecoder(r.Body).Decode(&req)

    access, refresh, err := h.UserService.Login(r.Context(), req.Email, req.Password)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte(err.Error()))
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "access_token":  access,
        "refresh_token": refresh,
    })
}

func (h *UserHandler) Refresh(w http.ResponseWriter, r *http.Request) {
    var req struct {
        RefreshToken string `json:"refresh_token"`
    }
    json.NewDecoder(r.Body).Decode(&req)

    access, refresh, err := h.UserService.Refresh(r.Context(), req.RefreshToken)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte(err.Error()))
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "access_token":  access,
        "refresh_token": refresh,
    })
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
    var req struct {
        RefreshToken string `json:"refresh_token"`
    }
    json.NewDecoder(r.Body).Decode(&req)

    _ = h.UserService.Logout(r.Context(), req.RefreshToken)
    w.WriteHeader(http.StatusOK)
}

package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Refresh(w http.ResponseWriter, r *http.Request, ps httprouter.Params) // new
	Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

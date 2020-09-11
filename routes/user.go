package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"people/interfaces"
	"people/models"
	"people/utils/errors"
	"people/utils/validators"
)

type UserHandler struct {
	UserService interfaces.UserService
}

func NewUserHandler(r *mux.Router, us interfaces.UserService ){
	handler := &UserHandler{
		UserService: us,
	}
	r.HandleFunc("/users", handler.List).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handler.Get).Methods("GET")
	r.HandleFunc("/users", handler.Add).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", handler.Update).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", handler.Remove).Methods("DELETE")
}

func (u *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := r.Context()
	user, err := u.UserService.Get(ctx,id)
	if err == errors.NotFound {
		errors.ResponseErrorMessage(w, http.StatusNotFound, errors.NotFound)
		return
	}

	if err != nil {
		errors.ResponseErrorMessage(w, http.StatusInternalServerError, errors.ErrUnhandled)
		return
	}

	errors.ResponseObject(w, http.StatusOK, user);
}

func (u *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := u.UserService.List(ctx)
	if err != nil {
		errors.ResponseErrorMessage(w, http.StatusInternalServerError, errors.ErrUnhandled)
		return
	}

	errors.ResponseObject(w, http.StatusOK, users)
}

func (u *UserHandler) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if(user == models.User{} || !validators.IsEmailValid(user.Email) || user.Name == "" || user.Email == "" || user.Password == ""){
		errors.ResponseErrorMessage(w, http.StatusBadRequest, errors.InvalidParamInput)
		return
	}

	err := u.UserService.Add(ctx, &user)
	if err != nil {
		errors.ResponseErrorMessage(w, http.StatusInternalServerError, errors.ErrUnhandled)
		return
	}

	errors.ResponseObject(w, http.StatusCreated, user)
}

func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := u.UserService.Update(ctx, &user)

	if err == errors.NotFound {
		errors.ResponseErrorMessage(w, http.StatusNotFound, errors.NotFound)
		return
	}

	if err != nil {
		errors.ResponseErrorMessage(w, http.StatusBadRequest, errors.InvalidParamInput)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := r.Context()
	err := u.UserService.Remove(ctx, id)

	if err == errors.NotFound {
		errors.ResponseErrorMessage(w, http.StatusNotFound, errors.NotFound)
		return
	}

	if err != nil {
		errors.ResponseErrorMessage(w, http.StatusInternalServerError, errors.ErrUnhandled)
		return
	}

	w.WriteHeader(http.StatusOK)
}
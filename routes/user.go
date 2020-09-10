package router

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"people/interfaces"
	"people/models"
	"people/utils/errors"
	"strconv"
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
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ctx := r.Context()
	user, err := u.UserService.Get(ctx,id)
	if err == errors.NotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	buffer, err := json.Marshal(user)
	if err!=nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	_, err = w.Write(buffer)
}

func (u *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := u.UserService.List(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	buffer, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	_, err = w.Write(buffer)

}

func (u *UserHandler) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := u.UserService.Add(ctx, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	buffer, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type","application/json")
	_, err = w.Write(buffer)
}

func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := u.UserService.Update(ctx, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ctx := r.Context()
	_, err = u.UserService.Get(ctx, id)
	if err == errors.NotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = u.UserService.Remove(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
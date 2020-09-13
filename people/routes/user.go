package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"people/interfaces"
	"people/models"
	"people/utils"
	"people/utils/validators"
)

type UserHandler struct {
	UserService interfaces.UserService
}

func NewUserHandler(r *mux.Router, us interfaces.UserService ){
	handler := &UserHandler{
		UserService: us,
	}
	r.HandleFunc("/users", handler.List).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/users/{id:[0-9]+}", handler.Get).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/users", handler.Add).Methods(http.MethodPost)
	r.HandleFunc("/users", handler.Update).Methods(http.MethodPut)
	r.HandleFunc("/users/{id:[0-9]+}", handler.Remove).Methods(http.MethodDelete)

}

func (u *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := r.Context()
	user, err := u.UserService.Get(ctx,id)
	if err == utils.NotFound {
		log.Println(err.Error())
		utils.ResponseErrorMessage(w, http.StatusNotFound, utils.NotFound)
		return
	}

	if err != nil {
		log.Println(err.Error())
		utils.ResponseErrorMessage(w, http.StatusInternalServerError, utils.ErrUnhandled)
		return
	}

	utils.ResponseObject(w, http.StatusOK, user);
}

func (u *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := u.UserService.List(ctx)
	if err != nil {
		log.Println(err.Error())
		utils.ResponseErrorMessage(w, http.StatusInternalServerError, utils.ErrUnhandled)
		return
	}

	utils.ResponseObject(w, http.StatusOK, users)
}

func (u *UserHandler) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if(user == models.User{} || !validators.IsEmailValid(user.Email) || user.Name == "" || user.Email == "" || user.Password == ""){
		utils.ResponseErrorMessage(w, http.StatusBadRequest, utils.InvalidParamInput)
		return
	}

	err := u.UserService.Add(ctx, &user)
	if err != nil {
		log.Println(err.Error())
		utils.ResponseErrorMessage(w, http.StatusInternalServerError, utils.ErrUnhandled)
		return
	}

	utils.ResponseObject(w, http.StatusCreated, user)
}

func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := u.UserService.Update(ctx, &user)

	if err == utils.NotFound {
		utils.ResponseErrorMessage(w, http.StatusNotFound, utils.NotFound)
		return
	}

	if err != nil {
		utils.ResponseErrorMessage(w, http.StatusBadRequest, utils.InvalidParamInput)
		return
	}

	utils.ResponseObject(w, http.StatusOK, nil)
}

func (u *UserHandler) Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Println(id)
	ctx := r.Context()
	err := u.UserService.Remove(ctx, id)

	if err == utils.NotFound {
		utils.ResponseErrorMessage(w, http.StatusNotFound, utils.NotFound)
		return
	}

	if err != nil {
		utils.ResponseErrorMessage(w, http.StatusInternalServerError, utils.ErrUnhandled)
		return
	}

	utils.ResponseObject(w, http.StatusOK, nil)
}
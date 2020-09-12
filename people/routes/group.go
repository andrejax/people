package router

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"people/interfaces"
	"people/models"
	"people/utils/errors"
)

type GroupHandler struct {
	GroupService interfaces.GroupService
}

func NewGroupHandler(r *mux.Router, gr interfaces.GroupService ){
	handler := &GroupHandler{
		GroupService: gr,
	}
	r.HandleFunc("/groups", handler.List).Methods("GET")
	r.HandleFunc("/groups/{id:[0-9]+}", handler.Get).Methods("GET")
	r.HandleFunc("/groups", handler.Add).Methods("POST")
	r.HandleFunc("/groups/{id:[0-9]+}", handler.Update).Methods("PUT")
	r.HandleFunc("/groups/{id:[0-9]+}", handler.Remove).Methods("DELETE")
}

func (g *GroupHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := r.Context()
	group, err := g.GroupService.Get(ctx,id)
	if err == errors.NotFound {
		errors.ResponseErrorMessage(w, http.StatusNotFound, errors.NotFound)
		return
	}

	if err != nil {
		errors.ResponseErrorMessage(w, http.StatusInternalServerError, errors.ErrUnhandled)
		return
	}

	errors.ResponseObject(w, http.StatusOK, group);
}

func (g *GroupHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	groups, err := g.GroupService.List(ctx)
	if err != nil {
		log.Println(err.Error())
		errors.ResponseErrorMessage(w, http.StatusInternalServerError, errors.ErrUnhandled)
		return
	}

	errors.ResponseObject(w, http.StatusOK, groups)
}

func (g *GroupHandler) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var group models.Group
	json.NewDecoder(r.Body).Decode(&group)

	if(cmp.Equal(group,models.Group{}) || group.Name == "" ){
		errors.ResponseErrorMessage(w, http.StatusBadRequest, errors.InvalidParamInput)
		return
	}

	err := g.GroupService.Add(ctx, &group)
	if err != nil {
		errors.ResponseErrorMessage(w, http.StatusInternalServerError, errors.ErrUnhandled)
		return
	}

	errors.ResponseObject(w, http.StatusCreated, group)
}

func (g *GroupHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var group models.Group
	json.NewDecoder(r.Body).Decode(&group)

	err := g.GroupService.Update(ctx, &group)

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

func (g *GroupHandler) Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := r.Context()
	err := g.GroupService.Remove(ctx, id)

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
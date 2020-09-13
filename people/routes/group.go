package router

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/mux"
	"net/http"
	"people/interfaces"
	"people/models"
	"people/utils"
)

type GroupHandler struct {
	GroupService interfaces.GroupService
}

func NewGroupHandler(r *mux.Router, gr interfaces.GroupService ){
	handler := &GroupHandler{
		GroupService: gr,
	}
	r.HandleFunc("/groups", handler.List).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/groups/{id:[0-9]+}", handler.Get).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/groups", handler.Add).Methods(http.MethodPost)
	r.HandleFunc("/groups", handler.Update).Methods(http.MethodPut)
	r.HandleFunc("/groups/{id:[0-9]+}", handler.Remove).Methods(http.MethodDelete)
}

func (g *GroupHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := r.Context()
	group, err := g.GroupService.Get(ctx,id)
	if err == utils.NotFound {
		utils.ResponseErrorMessage(w, http.StatusNotFound, utils.NotFound, err)
		return
	}

	if err != nil {
		utils.ResponseErrorMessage(w, http.StatusInternalServerError, utils.ErrUnhandled, err)
		return
	}

	utils.ResponseObject(w, http.StatusOK, group);
}

func (g *GroupHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	groups, err := g.GroupService.List(ctx)
	if err != nil {
		utils.ResponseErrorMessage(w, http.StatusInternalServerError, utils.ErrUnhandled, err)
		return
	}

	utils.ResponseObject(w, http.StatusOK, groups)
}

func (g *GroupHandler) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var group models.Group
	json.NewDecoder(r.Body).Decode(&group)

	if(cmp.Equal(group,models.Group{}) || group.Name == "" ){
		utils.ResponseErrorMessage(w, http.StatusBadRequest, utils.InvalidParamInput, utils.InvalidParamInput)
		return
	}

	err := g.GroupService.Add(ctx, &group)
	if err != nil {
		utils.ResponseErrorMessage(w, http.StatusInternalServerError, utils.ErrUnhandled, err)
		return
	}

	utils.ResponseObject(w, http.StatusCreated, group)
}

func (g *GroupHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var group models.Group
	json.NewDecoder(r.Body).Decode(&group)

	err := g.GroupService.Update(ctx, &group)

	if err == utils.NotFound {
		utils.ResponseErrorMessage(w, http.StatusNotFound, utils.NotFound, err)
		return
	}

	if err != nil {
		utils.ResponseErrorMessage(w, http.StatusBadRequest, utils.InvalidParamInput, err)
		return
	}

	utils.ResponseObject(w, http.StatusOK, group)
}

func (g *GroupHandler) Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := r.Context()
	err := g.GroupService.Remove(ctx, id)

	if err == utils.NotFound {
		utils.ResponseErrorMessage(w, http.StatusNotFound, utils.NotFound, err)
		return
	}

	if err != nil {
		utils.ResponseErrorMessage(w, http.StatusInternalServerError, utils.ErrUnhandled,err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"people/models"
	"people/tests"
	"strings"
	"testing"
)

func TestAddUser_Created(t *testing.T) {
	mockUser := models.User{
		Password: "pass",
		Email:    "email@google.com",
		Name:     "name",
		Id:       "0",
	}
	mockUserService := new(tests.UserService)
	handler := UserHandler {
		UserService: mockUserService,
	}

	j, err := json.Marshal(mockUser)
	assert.NoError(t, err)

	mockUserService.On("Add", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)

	req, err := http.NewRequest(echo.POST, "/users", bytes.NewReader(j))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	handler.Add(w,req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockUserService.AssertExpectations(t)
}

func TestAddUser_ServerError(t *testing.T) {
	mockUser := models.User{
		Password: "pass",
		Email:    "email@google.com",
		Name:     "name",
		Id:       "0",
	}
	mockUserService := new(tests.UserService)
	mockUserService.On("Add", mock.Anything, mock.AnythingOfType("*models.User")).Return(fmt.Errorf("Something went wrong."))

	j, err := json.Marshal(mockUser)
	assert.NoError(t, err)

	req, err := http.NewRequest(echo.POST, "/users", bytes.NewReader(j))
	assert.NoError(t, err)


   handler := UserHandler {
   		UserService: mockUserService,
   	}

	w := httptest.NewRecorder()
	handler.Add(w,req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockUserService.AssertExpectations(t)
}


func TestAddUser_MissingParameters_BadRequest(t *testing.T) {
	mockUser := struct {
		Test string
	}{
		"test@google.com",
	}
	mockUserService := new(tests.UserService)

	j, err := json.Marshal(mockUser)
	assert.NoError(t, err)

	req, err := http.NewRequest(echo.POST, "/users", bytes.NewReader(j))
	assert.NoError(t, err)

	handler := UserHandler {
		UserService: mockUserService,
	}

	w := httptest.NewRecorder()
	handler.Add(w,req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestRemoveUser_Success(t *testing.T){
	id := "1"
	mockUserService := new(tests.UserService)
	handler := UserHandler {
		UserService: mockUserService,
	}

	mockUserService.On("Remove", mock.Anything, id).Return(nil)

	req, err := http.NewRequest(echo.DELETE, "/users/"+id, strings.NewReader(""))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	var vars = make(map[string]string)
	vars["id"] = id
	req = mux.SetURLVars(req, vars)
	handler.Remove(w,req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockUserService.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T){
	mockUser := models.User{
		Password: "pass",
		Email:    "email@google.com",
		Name:     "name",
		Id:       "0",
	}
	mockUserService := new(tests.UserService)
	handler := UserHandler {
		UserService: mockUserService,
	}

	j, err := json.Marshal(mockUser)
	assert.NoError(t, err)

	mockUserService.On("Update", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)

	req, err := http.NewRequest(echo.POST, "/users", bytes.NewReader(j))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	var vars = make(map[string]string)
	vars["id"] = mockUser.Id
	req = mux.SetURLVars(req, vars)
	handler.Update(w,req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockUserService.AssertExpectations(t)
}

func TestListUsers_Succees(t *testing.T){
	mockUser:= models.User {
		Name:     "test",
		Password: "pstest",
		Email:    "emtest",
		Id:       "1",
	}
	mockUsers := make([]models.User,0)
	mockUsers = append(mockUsers, mockUser)

	mockUserService := new(tests.UserService)

	mockUserService.On("List", mock.Anything).Return(mockUsers, nil)

	req, err := http.NewRequest(echo.GET, "/users", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := UserHandler {
		UserService: mockUserService,
	}

	handler.List(rec,req)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserService.AssertExpectations(t)
}

func TestGetUser_Succees(t *testing.T){
	mockUser:= models.User {
		Name:     "test",
		Password: "pstest",
		Email:    "emtest",
		Id:       "1",
	}
	mockUserService := new(tests.UserService)

	mockUserService.On("Get", mock.Anything, mockUser.Id).Return(mockUser, nil)

	req, err := http.NewRequest(echo.GET, "/users/"+mockUser.Id, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	handler := UserHandler {
		UserService: mockUserService,
	}

	var vars = make(map[string]string)
	vars["id"] = mockUser.Id
	req = mux.SetURLVars(req, vars)
	handler.Get(rec,req)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserService.AssertExpectations(t)
}
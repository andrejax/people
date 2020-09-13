package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	NotFound = errors.New("Not found")
	InvalidParamInput = errors.New("Invalid parameters")
	ErrConflict = errors.New("Already exist")
	ErrUnhandled = errors.New("Something went wrong")
	)

type ErrorMessage struct {
	Message string 	`json:message`
}

func ResponseErrorMessage(wr http.ResponseWriter, status int, er error) {
	eo := &ErrorMessage{
		Message: er.Error(),
	}
	ResponseObject(wr, status, eo)
}

func ResponseObject(wr http.ResponseWriter, status int, o interface{}) {
	wr.Header().Set("Content-Type","application/json")
	wr.WriteHeader(status)
	buffer, _ := json.Marshal(o)
	wr.Write(buffer)
}
package docs

import (
	"people/models"
	"people/utils"
)

// swagger:operation POST /users Users addUser
// ---
// summary: Creates a new user.
// parameters:
// - name: user
//   in: body
//   description: user model
//   required: true
//   schema:
//      "$ref": "#/definitions/User"
// responses:
//   200:
//      description: user Response
//      "$ref": "#/responses/userResponse"
//   400:
//      description: error Response
//      "$ref": "#/responses/errorResponse"

// Error response
//swagger:response errorResponse
type errorResponseWrapper struct {
	// in:body
	Body utils.ErrorMessage
}

// User response
// swagger:response userResponse
type userResponseWrapper struct {
	// in:body
	Body models.User
}

// Users response
// swagger:response usersResponse
type usersResponseWrapper struct {
	// in:body
	Body []models.User
}

// swagger:parameters userParam
type userParamWrapper struct {
	// in:body
	Body models.User
}


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
//   201:
//      "$ref": "#/responses/userResponse"
//   400:
//      "$ref": "#/responses/errorResponse"


// swagger:operation PUT /users Users updateUser
// ---
// summary: Updates an existing user.
// parameters:
// - name: user
//   in: body
//   description: user model
//   required: true
//   schema:
//      "$ref": "#/definitions/User"
// responses:
//   200:
//      "$ref": "#/responses/userResponse"
//   400:
//      "$ref": "#/responses/errorResponse"
//   404:
//      "$ref": "#/responses/errorResponse"

// swagger:operation DELETE /users/{id} Users deleteUser
// ---
// summary: Deletes an existing user.
// parameters:
// - name: id
//   in: path
//   description: id of user
//   type: string
//   required: true
// responses:
//   200:
//      description: user deleted
//   404:
//     "$ref": "#/responses/errorResponse"


// swagger:operation Get /users/{id} Users getUser
// ---
// summary: Gets an existing user.
// parameters:
// - name: id
//   in: path
//   description: id of user
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/userResponse"
//   "404":
//     "$ref": "#/responses/errorResponse"

// swagger:route Get /users Users listUsers
//---
//// summary: Gets a list of users.
// responses:
//   200: usersResponse
//   400: errorResponse


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


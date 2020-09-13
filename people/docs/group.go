package docs

import (
	"people/models"
)

// swagger:operation POST /groups Groups addGroup
// ---
// summary: Creates a new group.
// parameters:
// - name: group
//   in: body
//   description: group model
//   required: true
//   schema:
//      "$ref": "#/definitions/Group"
// responses:
//   201:
//      "$ref": "#/responses/groupResponse"
//   400:
//      "$ref": "#/responses/errorResponse"


// swagger:operation PUT /groups Groups updateGroup
// ---
// summary: Updates an existing group.
// parameters:
// - name: group
//   in: body
//   description: group model
//   required: true
//   schema:
//      "$ref": "#/definitions/Group"
// responses:
//   200:
//      "$ref": "#/responses/groupResponse"
//   400:
//      "$ref": "#/responses/errorResponse"
//   404:
//      "$ref": "#/responses/errorResponse"

// swagger:operation DELETE /groups/{id} Groups deleteGroup
// ---
// summary: Deletes an existing group.
// parameters:
// - name: id
//   in: path
//   description: id of group
//   type: string
//   required: true
// responses:
//   "200":
//      description: group deleted
//   "404":
//     "$ref": "#/responses/errorResponse"


// swagger:operation Get /groups/{id} Groups getGroup
// ---
// summary: Gets an existing group.
// parameters:
// - name: id
//   in: path
//   description: id of group
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/groupResponse"
//   "404":
//     "$ref": "#/responses/errorResponse"

// swagger:route Get /groups Groups listGroups
//---
//// summary: Gets a list of groups.
// responses:
//   200: groupsResponse
//   400: errorResponse

// Group response
// swagger:response groupResponse
type groupResponseWrapper struct {
	// in:body
	Body models.Group
}

// Group response
// swagger:response groupsResponse
type groupssResponseWrapper struct {
	// in:body
	Body []models.Group
}

// Groups response
// swagger:parameters groupParam
type groupParamWrapper struct {
	// in:body
	Body models.Group
}
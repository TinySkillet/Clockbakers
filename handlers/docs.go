// Package classification of Clockmakers API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
//	- application/json
//
// swagger:meta

package handlers

//
// NOTE: Types defined here are purely for documentation purpose
// these types are not used by any of the handlers

// GenericError is a generic error messages returned by the server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Valdiation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// collection of the errors
	// in: body
}

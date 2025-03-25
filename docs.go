// Package classification Clockbakers API
//
// # Documentation for Clockbakers API
//
// The Clockbakers API provides a backend service for managing bakery orders, reviews, and user authentication for the Clockbakers bakery.
//
// # Authentication
// This API uses JWT-based authentication for secure access.
// Users must include a valid JWT token in the Authorization header.
//
// # Contact Information
//
//	Contact:
//	  <mirageaditya@gmail.com>
//
//	License:
//	  MIT
//
// swagger:meta
package main

// Swagger spec configuration
//
// Schemes: http
// BasePath: /v1/
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// SecurityDefinitions:
//  Bearer:
//    type: apiKey
//    name: Authorization
//    in: header

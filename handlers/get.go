package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TinySkillet/ClockBakers/internal/database"
	m "github.com/TinySkillet/ClockBakers/models"
	"github.com/google/uuid"
)

// swagger:route GET /healthz system healthz
// Check if the API server is running
// responses:
//   200: emptyResponse

// HandleHealthz returns a 200 OK response if the server is running
func (a *APIServer) HandleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(struct{}{})
}

// swagger:route GET /error system error
// Test error handling
// responses:
//   400: emptyResponse

// HandleError returns a 400 Bad Request response for testing error handling
func (a *APIServer) HandleError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(struct{}{})
}

// swagger:route GET /users/{id} users getUserById
// Get a user by their ID
// responses:
//   302: userResponse
//   400: errorResponse
//   404: errorResponse

// swagger:parameters getUserById
type getUserByIdParams struct {
	// User ID
	// in: query
	// required: true
	ID string `json:"id"`
}

// HandleGetUserById retrieves a user by their UUID
func (a *APIServer) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	queries := a.getQueries()
	query := r.URL.Query()
	idString := query.Get("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		m.RespondWithError(w, "Invalid id!", http.StatusBadRequest)
		return
	}
	dbUser, err := queries.GetUserByID(r.Context(), id)
	if err != nil {
		m.RespondWithError(w, "User not found!", http.StatusNotFound)
		return
	}
	user := m.DBUserToUser(dbUser)
	m.RespondWithJSON(w, user, http.StatusFound)
}

// swagger:route GET /users users getUsers
// Get users with optional filters
// responses:
//   200: usersResponse
//   404: errorResponse

// swagger:parameters getUsers
type getUsersParams struct {
	// First name filter
	// in: query
	FirstName string `json:"first-name"`

	// Last name filter
	// in: query
	LastName string `json:"last-name"`

	// Phone number filter
	// in: query
	PhoneNo string `json:"phone-no"`

	// Email filter
	// in: query
	Email string `json:"email"`
}

// swagger:response usersResponse
type usersResponseWrapper struct {
	// List of users
	// in: body
	Body []m.User
}

// HandleGetUsers retrieves users with optional filters
func (a *APIServer) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	queries := a.getQueries()
	query := r.URL.Query()
	first_name := query.Get("first-name")
	last_name := query.Get("last-name")
	phone_no := query.Get("phone-no")
	email := query.Get("email")
	dbUsers, err := queries.GetUsers(r.Context(), database.GetUsersParams{
		Column1: first_name,
		Column2: last_name,
		Column3: phone_no,
		Column4: email,
	})
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusNotFound)
		return
	}
	users := m.DBUsersToUsers(dbUsers)
	m.RespondWithJSON(w, users, http.StatusOK)
}

// swagger:route GET /products products getProducts
// Get products with optional filters
// responses:
//   200: productsResponse
//   400: errorResponse

// swagger:parameters getProducts
type getProductsParams struct {
	// Product name filter
	// in: query
	Name string `json:"name"`

	// Category filter
	// in: query
	Category string `json:"category"`

	// Minimum price filter
	// in: query
	MinPrice float64 `json:"min-price"`

	// Maximum price filter
	// in: query
	MaxPrice float64 `json:"max-price"`
}

// swagger:response productsResponse
type productsResponseWrapper struct {
	// List of products
	// in: body
	Body []m.Product
}

// HandleGetProducts retrieves products with optional filters
func (a *APIServer) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	// parse query parameters
	query := r.URL.Query()
	name := query.Get("name")
	category := query.Get("category")
	minPriceStr := query.Get("min-price")
	maxPriceStr := query.Get("max-price")
	var minPrice, maxPrice *float64
	if minPriceStr != "" {
		val, err := strconv.ParseFloat(minPriceStr, 64)
		if err != nil {
			m.RespondWithError(w, "Invalid min_price value in query", http.StatusBadRequest)
			return
		}
		minPrice = &val
	}
	if maxPriceStr != "" {
		val, err := strconv.ParseFloat(maxPriceStr, 64)
		if err != nil {
			m.RespondWithError(w, "Invalid max_price value in query", http.StatusBadRequest)
			return
		}
		maxPrice = &val
	}
	queries := a.getQueries()
	products, err := queries.GetProducts(r.Context(), database.GetProductsParams{
		Column1: name,
		Column2: minPrice,
		Column3: maxPrice,
		Column4: category,
	})
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}
	prds := m.DBProductsToProducts(products)
	m.RespondWithJSON(w, prds, http.StatusOK)
}

// swagger:route GET /categories categories getCategories
// Get all product categories
// responses:
//   200: categoriesResponse
//   500: errorResponse

// swagger:response categoriesResponse
type categoriesResponseWrapper struct {
	// List of categories
	// in: body
	Body []string
}

// HandleGetCategories retrieves all available product categories
func (a *APIServer) HandleGetCategories(w http.ResponseWriter, r *http.Request) {
	queries := a.getQueries()
	dbCats, err := queries.GetCategories(r.Context())
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.RespondWithJSON(w, dbCats, http.StatusOK)
}

func (a *APIServer) HandleGetOrderById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	orderID := query.Get("id")
	id, err := uuid.Parse(orderID)
	if err != nil {
		m.RespondWithError(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	dbOrder, err := queries.GetOrder(r.Context(), id)
	if err != nil {
		m.RespondWithError(w, "Order not found", http.StatusNotFound)
		return
	}

	order := m.DBOrderToOrder(dbOrder)
	m.RespondWithJSON(w, order, http.StatusOK)
}

func (a *APIServer) HandleListOrders(w http.ResponseWriter, r *http.Request) {
	userIDParam := r.URL.Query().Get("uid")
	status := r.URL.Query().Get("status")

	var userID *uuid.UUID
	if userIDParam != "" {
		parsedID, err := uuid.Parse(userIDParam)
		if err != nil {
			m.RespondWithError(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		userID = &parsedID
	}

	queries := a.getQueries()
	dbOrders, err := queries.ListOrders(r.Context(), database.ListOrdersParams{
		Column1: userID,
		Column2: status,
	})
	if err != nil {
		m.RespondWithError(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	orders := m.DBOrdersToOrders(dbOrders)
	m.RespondWithJSON(w, orders, http.StatusOK)
}

func (a *APIServer) HandleGetPopularItems(w http.ResponseWriter, r *http.Request) {
	query := a.getQueries()
	dbItems, err := query.GetPopularItems(r.Context())
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	items := m.DBPopularProductsToPopularProducts(dbItems)
	m.RespondWithJSON(w, items, http.StatusOK)
}

func (a *APIServer) HandleGetReviews(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	productID := query.Get("pid")
	id := query.Get("id")
	userID := query.Get("uid")

	var reviewParams m.GetReviewsParams

	if productID != "" {
		parsedProductID, err := uuid.Parse(productID)
		if err != nil {
			m.RespondWithError(w, "Invalid product ID", http.StatusBadRequest)
			return
		}
		reviewParams.ProductID = uuid.NullUUID{UUID: parsedProductID, Valid: true}
	}

	if id != "" {
		parsedID, err := uuid.Parse(id)
		if err != nil {
			m.RespondWithError(w, "Invalid review ID", http.StatusBadRequest)
			return
		}
		reviewParams.ID = uuid.NullUUID{UUID: parsedID, Valid: true}
	}

	if userID != "" {
		parsedUserID, err := uuid.Parse(userID)
		if err != nil {
			m.RespondWithError(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		reviewParams.UserId = uuid.NullUUID{UUID: parsedUserID, Valid: true}
	}

	queries := a.getQueries()
	reviews, err := queries.GetReviews(r.Context(), database.GetReviewsParams{
		Column1: reviewParams.ID.UUID,
		Column2: reviewParams.ProductID.UUID,
		Column3: reviewParams.UserId.UUID,
	})
	if err != nil {
		m.RespondWithError(w, "Failed to fetch reviews: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var responseReviews []m.Review
	for _, dbReview := range reviews {
		responseReviews = append(responseReviews, m.DBReviewToReview(dbReview))
	}
	m.RespondWithJSON(w, responseReviews, http.StatusOK)
}

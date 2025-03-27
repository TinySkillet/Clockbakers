package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/TinySkillet/ClockBakers/internal/database"
	m "github.com/TinySkillet/ClockBakers/models"
	"github.com/google/uuid"
)

// swagger:route PUT /v1/user users updateUser
// Update an existing user's information
// responses:
//   200: userResponse
//   400: errorResponse

// swagger:parameters updateUser
type updateUserParams struct {
	// The ID of the user to update
	// in: query
	// required: true
	UserID string `json:"uid"`
	// User information to update
	// in: body
	// required: true
	Body m.User
}

// HandleUpdateUser updates a user's information
func (a *APIServer) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	idStr := query.Get("uid")
	if idStr == "" {
		m.RespondWithError(w, "Invalid query parameters: uid!", http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(idStr)
	if err != nil {
		m.RespondWithError(w, "Parameter uid should be a uuid!", http.StatusBadRequest)
		return
	}

	params := m.User{}
	err = m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	phoneNo := sql.NullString{String: params.PhoneNo, Valid: params.PhoneNo != ""}
	email := sql.NullString{String: params.Email, Valid: params.Email != ""}

	queries := a.getQueries()
	dbUpdatedUser, err := queries.UpdateUser(r.Context(), database.UpdateUserParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		PhoneNo:   phoneNo,
		Address:   email,
		ID:        uid,
	})

	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser := m.DBUserToUser(dbUpdatedUser)
	m.RespondWithJSON(w, updatedUser, http.StatusOK)
}

// swagger:route PUT /v1/category category updateCategory
// Update an existing product category
// responses:
//   200: categoryResponse
//   400: errorResponse

// swagger:parameters updateCategory
type updateCategoryParams struct {
	// The current name of the category
	// in: query
	// required: true
	OldCategory string `json:"old-cat"`
	// The new name for the category
	// in: query
	// required: true
	NewCategory string `json:"new-cat"`
}

// HandleUpdateCategory updates a product category name
func (a *APIServer) HandleUpdateCategory(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	oldCatName := query.Get("old-cat")
	if oldCatName == "" {
		m.RespondWithError(w, "Invalid query parameter: old-cat!", http.StatusBadRequest)
		return
	}

	newCatName := query.Get("new-cat")
	if newCatName == "" {
		m.RespondWithError(w, "Invalid query parameter: new-cat!", http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	updatedCat, err := queries.UpdateCategory(r.Context(), database.UpdateCategoryParams{
		Name:   newCatName,
		Name_2: oldCatName,
	})
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}
	m.RespondWithJSON(w, updatedCat, http.StatusOK)
}

// swagger:route PUT /v1/product product updateProduct
// Update an existing product
// responses:
//   200: productResponse
//   400: errorResponse

// swagger:parameters updateProduct
type updateProductParams struct {
	// Product information to update
	// in: body
	// required: true
	Body m.Product
}

// HandleUpdateProduct updates a product's information
func (a *APIServer) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	var params m.Product
	if err := m.FromJSON(r, &params); err != nil {
		m.RespondWithError(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	updateParams := database.UpdateProductParams{
		UpdatedAt: time.Now().UTC(),
	}

	if params.SKU != "" {
		updateParams.Sku = sql.NullString{String: params.SKU, Valid: true}
	}
	if params.Name != "" {
		updateParams.Name = sql.NullString{String: params.Name, Valid: true}
	}
	if params.Description != "" {
		updateParams.Description = sql.NullString{String: params.Description, Valid: true}
	}
	if params.Price != 0 {
		updateParams.Price = sql.NullFloat64{Float64: float64(params.Price), Valid: true}
	}
	if params.StockQty != 0 {
		updateParams.StockQty = sql.NullInt32{Int32: int32(params.StockQty), Valid: true}
	}
	if params.CategoryName != "" {
		updateParams.Category = sql.NullString{String: params.CategoryName, Valid: true}
	}

	queries := a.getQueries()
	dbUpdatedProduct, err := queries.UpdateProduct(r.Context(), updateParams)
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedProduct := m.DBProductToProduct(dbUpdatedProduct)
	m.RespondWithJSON(w, updatedProduct, http.StatusOK)
}

// swagger:route PUT /v1/cart cart reduceItemQtyFromCart
// Reduce the quantity of an item in the user's shopping cart
// responses:
//   200: emptyResponse
//   400: errorResponse

// swagger:parameters reduceItemQtyFromCart
type reduceItemQtyFromCartParams struct {
	// The quantity to reduce
	// in: query
	// required: true
	Quantity string `json:"qty"`
	// The product ID of the item
	// in: query
	// required: true
	ProductID string `json:"pid"`
	// The user ID associated with the cart
	// in: query
	// required: true
	UserID string `json:"uid"`
}

// HandleReduceItemQtyFromCart reduces the quantity of an item in the cart
func (a *APIServer) HandleReduceItemQtyFromCart(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	quantity := query.Get("qty")
	product_id := query.Get("pid")
	user_id := query.Get("uid")

	qty, err := strconv.Atoi(quantity)
	if err != nil {
		m.RespondWithError(w, "Invalid query parameter: qty!", http.StatusBadRequest)
		return
	}

	pid, err := uuid.Parse(product_id)
	if err != nil {
		m.RespondWithError(w, "Invalid query parameter: pid!", http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(user_id)
	if err != nil {
		m.RespondWithError(w, "Invalid query parameter: cart_id!", http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	err = queries.ReduceQuantityFromCart(r.Context(), database.ReduceQuantityFromCartParams{
		Quantity:  int32(qty),
		ProductID: pid,
		UserID:    uid,
	})

	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}
	m.RespondWithJSON(w, struct{}{}, http.StatusOK)
}

// swagger:route PUT /v1/order order updateOrderStatus
// Update the status of an existing order
// responses:
//   200: orderResponse
//   400: errorResponse
//   500: errorResponse

// swagger:parameters updateOrderStatus
type updateOrderStatusParams struct {
	// The ID of the order to update
	// in: query
	// required: true
	ID string `json:"id"`
	// The updated order status
	// in: body
	// required: true
	Body m.Order
}

// HandleUpdateOrderStatus updates the status of an order
func (a *APIServer) HandleUpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	orderID := query.Get("id")
	id, err := uuid.Parse(orderID)
	if err != nil {
		m.RespondWithError(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	params := m.Order{}
	err = m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!", http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	dbOrder, err := queries.UpdateOrderStatus(r.Context(), database.UpdateOrderStatusParams{
		ID:     id,
		Status: database.OrderStatus(params.Status),
	})
	if err != nil {
		m.RespondWithError(w, "Failed to update order status", http.StatusInternalServerError)
		return
	}

	order := m.DBOrderToOrder(dbOrder)
	m.RespondWithJSON(w, order, http.StatusOK)
}

// swagger:route PUT /v1/review review updateReview
// Update an existing product review
// responses:
//   200: reviewResponse
//   400: errorResponse
//   500: errorResponse

// swagger:parameters updateReview
type updateReviewParams struct {
	// The Rating of the Review
	// in: body
	// name: rating
	// description: Rating of the review to update
	// required: true
	// type: int
	// format: int
	Rating int32 `json:"rating"`

	// The Review Comment
	// in: body
	// name: comment
	// description: Comment of the review to update
	// required: true
	// type: string
	// format: uuid
	Comment string `json:"comment"`
}

// HandleUpdateReview updates a product review
func (a *APIServer) HandleUpdateReview(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	review_id := query.Get("review_id")
	user_id := query.Get("user_id")

	id, err := uuid.Parse(review_id)
	if err != nil {
		m.RespondWithError(w, "Invalid review id!", http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(user_id)
	if err != nil {
		m.RespondWithError(w, "Invalid user id!", http.StatusBadRequest)
		return
	}

	params := m.Review{}
	err = m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!"+err.Error(), http.StatusBadRequest)
		return
	}

	err = params.Validate()
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	dbReview, err := queries.UpdateReview(r.Context(), database.UpdateReviewParams{
		ID:        id,
		UserID:    uid,
		Rating:    params.Rating,
		Comment:   params.Comment,
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		m.RespondWithError(w, "Failed to update review!", http.StatusInternalServerError)
		return
	}

	review := m.DBReviewToReview(dbReview)
	m.RespondWithJSON(w, review, http.StatusOK)
}

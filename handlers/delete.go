package handlers

import (
	"net/http"

	"github.com/TinySkillet/ClockBakers/internal/database"
	m "github.com/TinySkillet/ClockBakers/models"
	"github.com/google/uuid"
)

// swagger:route DELETE /categories categories deleteCategory
// Delete an existing product category
// responses:
//   200: emptyResponse
//   400: errorResponse

// swagger:parameters deleteCategory
type deleteCategoryParams struct {
	// The name of the category to delete
	// in: query
	// required: true
	Category string `json:"category"`
}

// HandleDeleteCategory deletes a product category by name
func (a *APIServer) HandleDeleteCategory(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	cat := query.Get("category")
	if cat == "" {
		m.RespondWithError(w, "Invalid query parameter: category!", http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	err := queries.DeleteCategory(r.Context(), cat)
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}
	m.RespondWithJSON(w, struct{}{}, http.StatusOK)
}

// swagger:route DELETE /products products deleteProduct
// Delete an existing product
// responses:
//   200: emptyResponse
//   400: errorResponse

// swagger:parameters deleteProduct
type deleteProductParams struct {
	// The SKU of the product to delete
	// in: query
	// required: true
	SKU string `json:"sku"`
}

// HandleDeleteProduct deletes a product by SKU
func (a *APIServer) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	sku := query.Get("sku")
	if sku == "" {
		m.RespondWithError(w, "Invalid query parameter: sku!", http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	err := queries.DeleteProduct(r.Context(), sku)
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}
	m.RespondWithJSON(w, struct{}{}, http.StatusOK)
}

// swagger:route DELETE /cart/items cart deleteItemFromCart
// Remove an item from the user's shopping cart
// responses:
//   200: emptyResponse
//   400: errorResponse

// swagger:parameters deleteItemFromCart
type deleteItemFromCartParams struct {
	// The product ID of the item to remove
	// in: query
	// required: true
	ProductID string `json:"pid"`
	// The user ID associated with the cart
	// in: query
	// required: true
	UserID string `json:"uid"`
}

// HandleDeleteItemFromCart removes a product from the user's shopping cart
func (a *APIServer) HandleDeleteItemFromCart(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	product_id := query.Get("pid")
	user_id := query.Get("uid")

	pid, err := uuid.Parse(product_id)
	if err != nil {
		m.RespondWithError(w, "Invalid query parameter: pid!", http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(user_id)
	if err != nil {
		m.RespondWithError(w, "Invalid query parameter: uid!", http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	err = queries.RemoveFromCart(r.Context(), database.RemoveFromCartParams{
		ProductID: pid,
		UserID:    uid,
	})
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}
	m.RespondWithJSON(w, struct{}{}, http.StatusOK)
}

// swagger:route DELETE /orders orders deleteOrder
// Delete an existing order
// responses:
//   200: emptyResponse
//   400: errorResponse
//   500: errorResponse

// swagger:parameters deleteOrder
type deleteOrderParams struct {
	// The ID of the order to delete
	// in: query
	// required: true
	ID string `json:"id"`
}

// HandleDeleteOrder deletes an order by ID
func (a *APIServer) HandleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	orderID := query.Get("id")
	id, err := uuid.Parse(orderID)
	if err != nil {
		m.RespondWithError(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	err = queries.DeleteOrder(r.Context(), id)
	if err != nil {
		m.RespondWithError(w, "Failed to delete order!", http.StatusInternalServerError)
		return
	}

	m.RespondWithJSON(w, struct{}{}, http.StatusOK)
}

// swagger:route DELETE /reviews reviews deleteReview
// Delete an existing product review
// responses:
//   200: emptyResponse
//   400: errorResponse
//   500: errorResponse

// swagger:parameters deleteReview
type deleteReviewParams struct {
	// The ID of the review to delete
	// in: query
	// required: true
	ID string `json:"id"`
}

// HandleDeleteReview deletes a product review by ID
func (a *APIServer) HandleDeleteReview(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	review_id := query.Get("id")
	id, err := uuid.Parse(review_id)
	if err != nil {
		m.RespondWithError(w, "Invalid review id!", http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	err = queries.DeleteReview(r.Context(), id)
	if err != nil {
		m.RespondWithError(w, "Failed to delete review!", http.StatusInternalServerError)
		return
	}

	m.RespondWithJSON(w, struct{}{}, http.StatusOK)
}

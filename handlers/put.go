package handlers

import (
	"net/http"

	"github.com/TinySkillet/ClockBakers/internal/database"
	m "github.com/TinySkillet/ClockBakers/models"
	"github.com/google/uuid"
)

func (a *APIServer) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	idStr := query.Get("id")
	if idStr == "" {
		m.RespondWithError(w, "Invalid query parameters: id!", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		m.RespondWithError(w, "Parameter id should be a uuid!", http.StatusBadRequest)
		return
	}

	params := m.User{}
	m.FromJSON(r, params)

	queries := a.getQueries()
	dbUpdatedUser, err := queries.UpdateUser(r.Context(), database.UpdateUserParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		PhoneNo:   params.PhoneNo,
		Address:   params.Address,
		ID:        id,
	})

	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser := m.DBUserToUser(dbUpdatedUser)
	m.RespondWithJSON(w, updatedUser, http.StatusOK)
}

func (a *APIServer) HandleUpdateCategory(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	oldCatName := query.Get("category")
	if oldCatName == "" {
		m.RespondWithError(w, "Invalid query parameter: sku!", http.StatusBadRequest)
		return
	}

	params := m.Category{}
	m.FromJSON(r, params)

	queries := a.getQueries()
	updatedCat, err := queries.UpdateCategory(r.Context(), database.UpdateCategoryParams{
		Name: params.CategoryName,
	})
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}
	m.RespondWithJSON(w, updatedCat, http.StatusOK)
}

func (a *APIServer) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	oldCatName := query.Get("category")
	if oldCatName == "" {
		m.RespondWithError(w, "Invalid query parameter: sku!", http.StatusBadRequest)
		return
	}

	params := m.Product{}
	m.FromJSON(r, params)

	queries := a.getQueries()
	dbUpdatedProduct, err := queries.UpdateProduct(r.Context(), database.UpdateProductParams{
		Sku:         params.SKU,
		Name:        params.Name,
		Description: params.Description,
		Price:       float32(params.Price),
		StockQty:    int32(params.StockQty),
		Category:    params.CategoryName,
	})
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedProduct := m.DBProductToProduct(dbUpdatedProduct)
	m.RespondWithJSON(w, updatedProduct, http.StatusOK)
}

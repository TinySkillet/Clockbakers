package handlers

import (
	"net/http"

	m "github.com/TinySkillet/ClockBakers/models"
)

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

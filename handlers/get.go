package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TinySkillet/ClockBakers/internal/database"
	m "github.com/TinySkillet/ClockBakers/models"
	"github.com/google/uuid"
)

// healthz handler to check if the server is up
func (a *APIServer) HandleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(struct{}{})
}

func (a *APIServer) HandleError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(struct{}{})
}

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

// handler to get categories
func (a *APIServer) HandleGetCategories(w http.ResponseWriter, r *http.Request) {
	queries := a.getQueries()

	dbCats, err := queries.GetCategories(r.Context())
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m.RespondWithJSON(w, dbCats, http.StatusOK)
}

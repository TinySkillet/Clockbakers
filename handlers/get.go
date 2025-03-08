package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TinySkillet/ClockBakers/internal/database"
	m "github.com/TinySkillet/ClockBakers/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

func (a *APIServer) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	queries := a.getQueries()
	us, err := queries.GetUsers(r.Context())
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}
	users := m.DBUsersToUsers(us)
	m.RespondWithJSON(w, users, http.StatusOK)
}

func (a *APIServer) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	dbqueries := a.getQueries()

	query := r.URL.Query()
	idString := query.Get("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		m.RespondWithError(w, "Invalid id query parameter!", http.StatusBadRequest)
		return
	}

	dbUser, err := dbqueries.GetUserByID(r.Context(), id)
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusNotFound)
		return
	}
	user := m.DBUserToUser(dbUser)
	m.RespondWithJSON(w, user, http.StatusOK)
}

func (a *APIServer) HandleGetUsersByName(w http.ResponseWriter, r *http.Request) {
	queries := a.getQueries()

	vars := mux.Vars(r)
	name := vars["name"]

	dbUsers, err := queries.GetUsersByName(r.Context(), name)
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
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
	minPriceStr := query.Get("min_price")
	maxPriceStr := query.Get("max_price")

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

package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/TinySkillet/ClockBakers/internal/database"
	middleware "github.com/TinySkillet/ClockBakers/middlewares"
	m "github.com/TinySkillet/ClockBakers/models"
	s "github.com/TinySkillet/ClockBakers/storage"
	"github.com/google/uuid"
)

func (a *APIServer) getQueries() *database.Queries {
	store, ok := a.storage.(s.PostgresStore)
	if !ok {
		a.l.Fatal("PostgresStore does not implement DataStore interface!")
	}
	return database.New(store.DB)
}

// login handler
func (a *APIServer) HandleLogin(w http.ResponseWriter, r *http.Request) {
	params := m.LoginRequest{}
	err := m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!", http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	// check if user with given email exists
	dbUser, err := queries.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		m.RespondWithError(w, "Invalid Email or Password!", http.StatusUnauthorized)
		return
	}

	// hash password of the user with that email
	hashedPw, err := m.HashPassword(params.Password)
	if err != nil {
		m.RespondWithError(w, "Invalid Password!", http.StatusBadRequest)
		return
	}

	// compare the hashes
	if m.VerifyPassword(hashedPw, dbUser.Password) {
		m.RespondWithError(w, "Invalid Email or Password!", http.StatusUnauthorized)
		return
	}

	// generate jwt token
	jwtToken, err := middleware.CreateToken(
		dbUser.ID,
		dbUser.Email,
		string(dbUser.Role),
	)
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send token in authorization header
	w.Header().Add("Authorization", "Bearer "+jwtToken)
	w.WriteHeader(http.StatusAccepted)
}

// register user handler
func (a *APIServer) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	params := m.User{}
	err := m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!", http.StatusBadRequest)
		log.Print(err)
		return
	}
	queries := a.getQueries()

	// hash the password
	hashed_password, err := m.HashPassword(params.Password)
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create user in the database
	u, err := queries.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		PhoneNo:   params.PhoneNo,
		Address:   params.Address,
		Password:  hashed_password,
		Role:      database.UserTypeCustomer,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!"+err.Error(), http.StatusBadRequest)
		return
	}

	// respond with the new user's information
	user := m.DBUserToUser(u)
	m.RespondWithJSON(w, user, http.StatusCreated)
}

// create category handler
func (a *APIServer) HandleCreateCategory(w http.ResponseWriter, r *http.Request) {
	params := m.Category{}
	err := m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!", http.StatusBadRequest)
		log.Print(err)
		return
	}

	queries := a.getQueries()
	cat, err := queries.CreateCategory(r.Context(), params.CategoryName)

	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!"+err.Error(), http.StatusBadRequest)
		return
	}

	m.RespondWithJSON(w, struct {
		Category string `json:"category_name"`
	}{Category: cat}, http.StatusCreated)
}

// create product handler
func (a *APIServer) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	params := m.Product{}
	err := m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!"+err.Error(), http.StatusBadRequest)
		log.Print(err)
		return
	}

	queries := a.getQueries()
	dbProduct, err := queries.CreateProduct(r.Context(), database.CreateProductParams{
		ID:          uuid.New(),
		Sku:         params.SKU,
		Name:        params.Name,
		Description: params.Description,
		Price:       float32(params.Price),
		StockQty:    int32(params.StockQty),
		Category:    params.CategoryName,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	})
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!"+err.Error(), http.StatusBadRequest)
		return
	}

	product := m.DBProductToProduct(dbProduct)
	m.RespondWithJSON(w, product, http.StatusCreated)
}

package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/TinySkillet/ClockBakers/internal/database"
	middleware "github.com/TinySkillet/ClockBakers/middlewares"
	m "github.com/TinySkillet/ClockBakers/models"
	s "github.com/TinySkillet/ClockBakers/storage"
	"github.com/google/uuid"
)

// APIServer utility function to get database queries
func (a *APIServer) getQueries() *database.Queries {
	store, ok := a.storage.(s.PostgresStore)
	if !ok {
		a.l.Fatal("PostgresStore does not implement DataStore interface!")
	}
	return database.New(store.DB)
}

// swagger:route POST /v1/login auth login
// Login to the application
// responses:
//   202: loginResponse
//   400: errorResponse
//   401: errorResponse
//   500: errorResponse

// swagger:parameters login
type loginParams struct {
	// Login credentials
	// in: body
	// required: true
	Body m.LoginRequest
}

// swagger:response loginResponse
type loginResponseWrapper struct {
	// Authorization token in header
	// in: header
	Authorization string
}

// HandleLogin authenticates a user and returns a JWT token
func (a *APIServer) HandleLogin(w http.ResponseWriter, r *http.Request) {
	params := m.LoginRequest{}
	err := m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!", http.StatusBadRequest)
		return
	}

	err = params.Validate()
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
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

// swagger:route POST /v1/user user createUser
// Create a new user account
// responses:
//   201: userResponse
//   400: errorResponse
//   500: errorResponse

// swagger:parameters createUser
type createUserParams struct {
	// User information for registration
	// in: body
	// required: true
	Body m.User
}

// swagger:response userResponse
type userResponseWrapper struct {
	// The created user
	// in: body
	Body m.User
}

// HandleCreateUser registers a new user and creates their shopping cart
func (a *APIServer) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	params := m.User{}
	err := m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!", http.StatusBadRequest)
		log.Print(err)
		return
	}

	err = params.Validate()
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	queries := a.getQueries()

	// hash the password
	hashed_password, err := m.HashPassword(params.Password)
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	phoneNo := sql.NullString{String: params.PhoneNo, Valid: params.PhoneNo != ""}
	address := sql.NullString{String: params.Address, Valid: params.Address != ""}

	// create user in the database
	u, err := queries.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		PhoneNo:   phoneNo,
		Address:   address,
		Password:  hashed_password,
		Role:      database.UserTypeCustomer,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!"+err.Error(), http.StatusBadRequest)
		return
	}

	// create a cart for the user
	_, err = queries.CreateCart(r.Context(), database.CreateCartParams{
		ID:     uuid.New(),
		UserID: u.ID,
	})

	if err != nil {
		m.RespondWithError(w, "An unexpected error occurred while initializing d!"+err.Error(), http.StatusInternalServerError)
		return
	}

	// respond with the new user's information
	user := m.DBUserToUser(u)
	m.RespondWithJSON(w, user, http.StatusCreated)
}

// swagger:route POST /v1/category category createCategory
// Create a new product category
// responses:
//   201: categoryResponse
//   400: errorResponse

// swagger:parameters createCategory
type createCategoryParams struct {
	// Category information
	// in: body
	// required: true
	Body m.Category
}

// swagger:response categoryResponse
type categoryResponseWrapper struct {
	// The created category
	// in: body
	Body struct {
		Category string `json:"category_name"`
	}
}

// HandleCreateCategory creates a new product category
func (a *APIServer) HandleCreateCategory(w http.ResponseWriter, r *http.Request) {
	params := m.Category{}
	err := m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!", http.StatusBadRequest)
		log.Print(err)
		return
	}

	err = params.Validate()
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
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

// swagger:route POST /v1/product product createProduct
// Create a new product
// responses:
//   201: productResponse
//   400: errorResponse

// swagger:parameters createProduct
type createProductParams struct {
	// Product information
	// in: body
	// required: true
	Body m.Product
}

// swagger:response productResponse
type productResponseWrapper struct {
	// The created product
	// in: body
	Body m.Product
}

// HandleCreateProduct creates a new bakery product
func (a *APIServer) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	params := m.Product{}
	err := m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!"+err.Error(), http.StatusBadRequest)
		log.Print(err)
		return
	}

	err = params.Validate()
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
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

// swagger:route POST /v1/cart cart addItemToCart
// Add an item to the user's shopping cart
// responses:
//   200: emptyResponse
//   400: errorResponse

// swagger:parameters addItemToCart
type addItemToCartParams struct {
	// Cart item information
	// in: body
	// required: true
	Body m.CartItem
}

// swagger:response emptyResponse
type emptyResponseWrapper struct {
	// Empty response
	// in: body
	Body struct{}
}

// HandleInsertItemInCart adds a product to the user's shopping cart
func (a *APIServer) HandleInsertItemInCart(w http.ResponseWriter, r *http.Request) {
	params := m.CartItem{}
	err := m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!", http.StatusBadRequest)
		log.Print(err)
		return
	}

	err = params.Validate()
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	queries := a.getQueries()

	err = queries.AddToCart(r.Context(), database.AddToCartParams{
		ID:        params.ID,
		Quantity:  params.Quantity,
		CartID:    params.CartID,
		ProductID: params.ProductID,
	})

	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!"+err.Error(), http.StatusBadRequest)
		return
	}

	m.RespondWithJSON(w, struct{}{}, http.StatusOK)
}

// swagger:route POST /v1/order order createOrder
// Create a new order
// responses:
//   201: orderResponse
//   400: errorResponse
//   500: errorResponse

// swagger:parameters createOrder
type createOrderParamsWrapper struct {
	// in: body
	// required: true
	Body m.Order
}

// swagger:response orderResponse
type orderResponseWrapper struct {
	// in: body
	Body m.Order
}

// HandleCreateOrder creates a new customer order with order items
func (a *APIServer) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	params := m.Order{}
	err := m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!", http.StatusBadRequest)
		return
	}

	err = params.Validate()
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	dbOrder, err := queries.CreateOrder(r.Context(), database.CreateOrderParams{
		ID:           uuid.New(),
		Status:       database.OrderStatusPending,
		TotalPrice:   params.TotalPrice,
		Quantity:     int32(params.Quantity),
		Pounds:       float32(params.Pounds),
		Message:      params.Message,
		DeliveryTime: database.DeliveryTimes(params.DeliveryTime),
		DeliveryDate: params.DeliveryDate,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
		ProductID:    params.ProductID,
		UserID:       params.UserID,
	})

	if err != nil {
		m.RespondWithError(w, "Failed to create order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	order := m.DBOrderToOrder(dbOrder)
	m.RespondWithJSON(w, order, http.StatusCreated)
}

// swagger:route POST /v1/review review createReview
// Create a new product review
// responses:
//   201: reviewResponse
//   400: errorResponse

// swagger:parameters createReview
type createReviewParams struct {
	// Review information
	// in: body
	// required: true
	Body m.Review
}

// swagger:response reviewResponse
type reviewResponseWrapper struct {
	// The created review
	// in: body
	Body m.Review
}

// swagger:response errorResponse
type errorResponseWrapper struct {
	// The error message
	// in: body
	Body struct {
		Error string `json:"error"`
	}
}

// CreateReview adds a new product review from a user
func (a *APIServer) HandleCreateReview(w http.ResponseWriter, r *http.Request) {
	params := m.Review{}
	err := m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params! "+err.Error(), http.StatusBadRequest)
		return
	}

	err = params.Validate()
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	dbReview, err := queries.CreateReview(r.Context(), database.CreateReviewParams{
		ID:        uuid.New(),
		Rating:    params.Rating,
		Comment:   params.Comment,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    params.UserID,
		ProductID: params.ProductID,
	})
	if err != nil {
		m.RespondWithError(w, "Failed to insert create review: "+err.Error(), http.StatusBadRequest)
		return
	}
	review := m.DBReviewToReview(dbReview)
	m.RespondWithJSON(w, review, http.StatusCreated)
}

// swagger:route POST /v1/address deliveryAddress createDeliveryAddress
// Create a new delivery address for a user.
//
// Responses:
//
//	201: addressResponse
//	400: errorResponse
//	500: errorResponse
//
// swagger:parameters createDeliveryAddress
type createDeliveryAddressParams struct {
	// Delivery address details
	// in: body
	// required: true
	Body struct {
		// The address details
		// example: "123 Main Street, City, Country"
		Address string `json:"address"`

		// User ID (UUID) to associate with the address
		// format: uuid
		UserID uuid.UUID `json:"user_id"`
	}
}

func (a *APIServer) HandleCreateDeliveryAddress(w http.ResponseWriter, r *http.Request) {
	params := m.DeliveryAddress{}

	err := m.FromJSON(r, &params)
	if err != nil {
		m.RespondWithError(w, "Invalid JSON params! "+err.Error(), http.StatusBadRequest)
		return
	}

	err = params.Validate()
	if err != nil {
		m.RespondWithError(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	dbAddress, err := queries.CreateDeliveryAddress(r.Context(), database.CreateDeliveryAddressParams{
		ID:      uuid.New(),
		Address: params.Address,
		UserID:  params.UserID,
	})

	if err != nil {
		m.RespondWithError(w, "Failed to create address: "+err.Error(), http.StatusBadRequest)
		return
	}

	address := m.DBDeliveryAddressToAddress(dbAddress)
	m.RespondWithJSON(w, address, http.StatusCreated)
}

func (a *APIServer) HandleResetPassword(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	email := query.Get("email")
	if email == "" {
		m.RespondWithError(w, "Invalid query parameter: email", http.StatusBadRequest)
		return
	}

	queries := a.getQueries()
	users, err := queries.GetUsers(r.Context(), database.GetUsersParams{
		Column4: email,
	})
	if err != nil {
		m.RespondWithError(w, "Couldn't find account with that email!", http.StatusBadRequest)
		return
	}

	if len(users) == 0 {
		m.RespondWithError(w, "Couldn't find account with that email!", http.StatusBadRequest)
		return
	}

	// user := users[0]

	resetCode, err := m.GenerateResetCode()
	if err != nil {
		m.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	expiresAt := time.Now().Add(3 * time.Hour).UTC()
	_, err = queries.CreatePasswordResetCode(r.Context(), database.CreatePasswordResetCodeParams{
		ID:        uuid.New(),
		Email:     email,
		Code:      resetCode,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now().UTC(),
	})

	if err != nil {
		m.RespondWithError(w, "Failed to create reset code", http.StatusInternalServerError)
		return
	}

	// subject := "Your Password Reset Code"
	// body := fmt.Sprintf("Hello %s,\n\nYour password reset code is: %s\n\nThis code will expire in 3 hours.\n", user.Username, resetCode)

	// // Setup SMTP configuration (replace with your configuration)
	// from := "no-reply@yourdomain.com"
	// password := "your-smtp-password"
	// smtpHost := "smtp.yourdomain.com"
	// smtpPort := "587"

}

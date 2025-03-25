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
		UserID:    params.UserID,
		ProductID: params.ProductID,
	})

	if err != nil {
		m.RespondWithError(w, "Invalid JSON params!"+err.Error(), http.StatusBadRequest)
		return
	}

	m.RespondWithJSON(w, struct{}{}, http.StatusOK)
}

// swagger:route POST /v1/order order createOrder
// Create a new order from cart items
// responses:
//   201: orderResponse
//   400: errorResponse
//   500: errorResponse

// swagger:parameters createOrder
type createOrderParams struct {
	// Order information including items
	// in: body
	// required: true
	Body m.Order
}

// swagger:response orderResponse
type orderResponseWrapper struct {
	// The created order
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
		ID:         uuid.New(),
		Status:     database.OrderStatusPending,
		TotalPrice: params.TotalPrice,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		UserID:     params.UserID,
	})
	if err != nil {
		m.RespondWithError(w, "Failed to create order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var createdItems []m.OrderItem
	// insert order items
	for _, item := range params.Items {
		dbOrderItem, err := queries.CreateOrderItem(r.Context(), database.CreateOrderItemParams{
			ID:              uuid.New(),
			Quantity:        int32(item.Quantity),
			PriceAtPurchase: float32(item.PriceAtPurchase),
			OrderID:         dbOrder.ID,
			ProductID:       item.ProductID,
		})
		if err != nil {
			m.RespondWithError(w, "Failed to insert order item: "+err.Error(), http.StatusInternalServerError)
			return
		}
		createdItems = append(createdItems, m.DBOrderItemToOrderItem(dbOrderItem))
	}

	order := m.DBOrderToOrder(dbOrder)
	order.Items = createdItems
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

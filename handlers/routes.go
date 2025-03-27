package handlers

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	m "github.com/TinySkillet/ClockBakers/middlewares"
	s "github.com/TinySkillet/ClockBakers/storage"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr    string
	storage s.DataStore
	l       *log.Logger
}

func NewAPIServer(listenAddr string, logger *log.Logger, s s.DataStore) *APIServer {
	return &APIServer{listenAddr, s, logger}
}

func (a *APIServer) Run() {
	// base router
	router := mux.NewRouter()

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{
			"Content-Type",
			"Authorization",
			"X-Requested-With",
			"X-CSRF-Token",
		}),
		handlers.MaxAge(3600),
		handlers.AllowCredentials(),
	)

	// cors middleware
	baseHandler := corsMiddleware(router)

	// v1 router, we can improve to other versions in the future
	v1Router := router.PathPrefix("/v1").Subrouter()

	// sub router for get methods, get requests are routed to this router
	getRouter := v1Router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/healthz", a.HandleHealthz)
	getRouter.HandleFunc("/error", a.HandleError)
	getRouter.HandleFunc("/users", m.MiddlewareValidateAdmin(a.HandleGetUsers))
	getRouter.HandleFunc("/products", a.HandleGetProducts)
	getRouter.HandleFunc("/categories", a.HandleGetCategories)
	getRouter.HandleFunc("/orders", m.MiddlewareValidateUser(a.HandleListOrders))
	getRouter.HandleFunc("/products/popular", a.HandleGetPopularItems)
	getRouter.HandleFunc("/reviews", m.MiddlewareValidateUser(a.HandleGetReviews))

	// sub router for post methods, post requests are routed to this router
	postRouter := v1Router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/login", a.HandleLogin)
	postRouter.HandleFunc("/user", a.HandleCreateUser)
	postRouter.HandleFunc("/category", m.MiddlewareValidateAdmin(a.HandleCreateCategory))
	postRouter.HandleFunc("/product", m.MiddlewareValidateAdmin(a.HandleCreateProduct))
	postRouter.HandleFunc("/cart", m.MiddlewareValidateUser(a.HandleInsertItemInCart))
	postRouter.HandleFunc("/order", m.MiddlewareValidateUser(a.HandleCreateOrder))
	postRouter.HandleFunc("/review", m.MiddlewareValidateUser(a.HandleCreateReview))

	// sub router for put methods, put requests are routed to this router
	putRouter := v1Router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/category", m.MiddlewareValidateAdmin(a.HandleUpdateCategory))
	putRouter.HandleFunc("/product", m.MiddlewareValidateAdmin(a.HandleUpdateProduct))
	putRouter.HandleFunc("/user", m.MiddlewareValidateUser(a.HandleUpdateUser))
	putRouter.HandleFunc("/cart", m.MiddlewareValidateUser(a.HandleReduceItemQtyFromCart))
	putRouter.HandleFunc("/order", m.MiddlewareValidateAdmin(a.HandleUpdateOrderStatus))
	putRouter.HandleFunc("/review", m.MiddlewareValidateUser(a.HandleUpdateReview))

	// sub router for delete methods, put requests are routed to this router
	deleteRouter := v1Router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/category", m.MiddlewareValidateAdmin(a.HandleDeleteCategory))
	deleteRouter.HandleFunc("/product", m.MiddlewareValidateAdmin(a.HandleDeleteProduct))
	deleteRouter.HandleFunc("/cart", m.MiddlewareValidateUser(a.HandleDeleteItemFromCart))
	deleteRouter.HandleFunc("/order", m.MiddlewareValidateUser(a.HandleDeleteOrder))
	deleteRouter.HandleFunc("/review", m.MiddlewareValidateUser(a.HandleDeleteReview))

	// serve swagger.json
	router.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./swagger.json")
	})

	// serve swagger docs with Redoc
	opts := middleware.RedocOpts{Title: "Clockbakers API Documentation", Template: ""}
	sh := middleware.Redoc(opts, nil)

	router.Handle("/docs", sh)

	// api server
	server := http.Server{
		Addr:         ":" + a.addr,
		Handler:      baseHandler,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// start the server on separate goroutine
	go func() {
		a.l.Print("Server started on PORT ", server.Addr)
		a.l.Fatal(server.ListenAndServe())
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// channel to check for interrupts or kill
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	// block until kill or interrupt
	sig := <-sigChan
	a.l.Print("Received terminate, graceful shutdown!", sig)
	server.Shutdown(ctx)
}

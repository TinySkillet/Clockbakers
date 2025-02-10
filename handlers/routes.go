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

	// v1 router, we can improve to other versions in the future
	v1Router := router.PathPrefix("/v1").Subrouter()

	// cors middleware
	v1Router.Use(m.CorsMiddleware)

	// sub router for get methods, get requests are routed to this router
	getRouter := v1Router.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/healthz", a.HandleHealthz)
	getRouter.HandleFunc("/error", a.HandleError)
	getRouter.HandleFunc("/user", m.MiddlewareValidateUser(a.HandleGetUser))
	getRouter.HandleFunc("/users", m.MiddlewareValidateAdmin(a.HandleGetUsers))
	getRouter.HandleFunc("/products", a.HandleGetProducts)
	getRouter.HandleFunc("/categories", a.HandleGetCategories)

	// sub router for post methods, post requests are routed to this router
	postRouter := v1Router.Methods(http.MethodPost).Subrouter()

	postRouter.HandleFunc("/login", a.HandleLogin)
	postRouter.HandleFunc("/user", a.HandleCreateUser)
	postRouter.HandleFunc("/category", m.MiddlewareValidateAdmin(a.HandleCreateCategory))
	postRouter.HandleFunc("/product", m.MiddlewareValidateAdmin(a.HandleCreateProduct))

	// sub router for put methods, put requests are routed to this router
	putRouter := v1Router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/category", m.MiddlewareValidateAdmin(a.HandleUpdateCategory))
	putRouter.HandleFunc("/product", m.MiddlewareValidateAdmin(a.HandleUpdateProduct))
	putRouter.HandleFunc("/user", m.MiddlewareValidateUser(a.HandleUpdateUser))

	// sub router for delete methods, put requests are routed to this router
	deleteRouter := v1Router.Methods(http.MethodDelete).Subrouter()

	deleteRouter.HandleFunc("/category", m.MiddlewareValidateAdmin(a.HandleDeleteCategory))
	deleteRouter.HandleFunc("/product", m.MiddlewareValidateAdmin(a.HandleDeleteProduct))

	// api server
	server := http.Server{
		Addr:         ":" + a.addr,
		Handler:      router,
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

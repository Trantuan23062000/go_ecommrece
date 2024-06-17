package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Trantuan23062000/go_ecommrece/service/user"
	"github.com/gorilla/mux"
)

// APIserver struct
type APIserver struct {
	addr string
	db   *sql.DB
}

// NewAPIServer initializes a new API server
func NewAPIServer(addr string, db *sql.DB) *APIserver {
	return &APIserver{
		addr: addr,
		db:   db,
	}
}

// Run starts the API server
func (s *APIserver) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func CreateServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (server *APIServer) RunServer() error {
	router := mux.NewRouter()

	// Setting a universal path prefix for the api routes
	subrouter := router.PathPrefix("/api/v1/").Subrouter()

	log.Println("Server listening on: http://localhost", server.addr)
	return http.ListenAndServe(server.addr, router)
}

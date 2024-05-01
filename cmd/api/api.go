package api

import (
	"api-go/services/user"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(add string, db *sql.DB) *APIServer {
	return &APIServer{addr: add, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)

	userService := user.NewHandler(userStore)
	userService.RegisterRouter(subRouter)

	log.Println("Listening On", s.addr)

	return http.ListenAndServe(s.addr, router)
}

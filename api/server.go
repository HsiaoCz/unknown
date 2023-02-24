package api

import (
	"go-hello/storage"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr     string
	authMiddleware *AuthMiddleware
	store          storage.MysqlStorage
}

func NewServer(listenAddr string, store storage.MysqlStorage) *Server {
	return &Server{
		listenAddr:     listenAddr,
		authMiddleware: NewAuthMiddleware(),
		store:          store,
	}
}

func (s *Server) Start() error {
	r := mux.NewRouter()
	r.HandleFunc("/user/signup", s.handleUserRegister).Methods("POST")
	r.HandleFunc("/user/login", s.handleUserSignup).Methods("POST")
    // thie router group that must be login
	us := r.PathPrefix("/user/login").Subrouter()
	us.Use(s.authMiddleware.AuthLoginMiddle)
	us.HandleFunc("/{id}", s.handleGetUserByID).Methods("GET")
	srv := http.Server{
		Handler:      r,
		Addr:         s.listenAddr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return srv.ListenAndServe()
}

package api

import (
	"go-hello/storage"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
	store      storage.MysqlStorage
}

func NewServer(listenAddr string, store storage.MysqlStorage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", s.handleGetUserByID).Methods("GET")
	r.HandleFunc("/user/register", s.handleUserRegister).Methods("POST")
	srv := http.Server{
		Handler:      r,
		Addr:         s.listenAddr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return srv.ListenAndServe()
}

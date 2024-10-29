package api

import (
	"encoding/json"
	"module/storage"
	"module/util"
	"net/http"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserById)
	http.HandleFunc("/user/id", s.handleDeleteUserById)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)

	json.NewEncoder(w).Encode(user)
}

func (s *Server) handleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)

	_ = util.Round2Dec(10.34434)

	json.NewEncoder(w).Encode(user)
}

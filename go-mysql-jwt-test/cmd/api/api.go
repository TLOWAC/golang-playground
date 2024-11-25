package api

import (
	"database/sql"
	"log"
	"module/service/user"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPISever(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := gin.Default()
	v1Router := router.Group("/api/v1")

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(v1Router)

	log.Println("Listening on", s.addr)

	return router.Run(s.addr)
}

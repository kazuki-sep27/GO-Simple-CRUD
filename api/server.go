package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/kazuki-sep27/simple_bank_go/db/sqlc"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//add router here
	router.POST("/accounts",server.createAccount)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
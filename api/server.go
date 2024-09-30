package api

import (
	db "github.com/Davison003/drr-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server makes HTTP requests for the banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// adding routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.PATCH("/accounts/:id", server.updateAccount)

	server.router = router
	return server
}

// Start func runs the HTTP server on a given address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

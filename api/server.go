package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/solracnet/go_finance_backend/db/sqlc"
)

type Server struct {
	store  *db.SqlStore
	router *gin.Engine
}

func NewServer(store *db.SqlStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/user", server.CreateUser)
	router.GET("/users/:username", server.GetUser)
	router.GET("/user/id/:id", server.GetUserById)

	router.POST("/category", server.CreateCategory)
	router.GET("/category/id/:id", server.GetCategory)
	router.GET("/category", server.GetCategories)
	router.DELETE("/category/:id", server.DeleteCategory)
	router.PUT("/category/:id", server.UpdateCategory)

	router.POST("/account", server.CreateAccount)
	router.GET("/account/id/:id", server.GetAccount)
	router.GET("/account", server.GetAccounts)
	router.DELETE("/account/:id", server.DeleteAccount)
	router.PUT("/account/:id", server.UpdateAccount)
	router.GET("/account/graph/:user_id/:type", server.GetAccountGraph)
	router.GET("/account-report", server.GetAccountReports)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

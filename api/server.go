package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/solracnet/go_finance_backend/db/sqlc"
)

type Server struct {
	store  *db.SqlStore
	router *gin.Engine
}

func CORSConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func NewServer(store *db.SqlStore) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.Use(CORSConfig())

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

	router.POST("/login", server.Login)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

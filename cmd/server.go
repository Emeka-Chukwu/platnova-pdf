package cmd

import (
	"net/http"
	httphandlers "pdf-go/internal/accounts/http"
	"pdf-go/internal/accounts/repository"
	"pdf-go/internal/accounts/services"
	"pdf-go/pkg"
	"pdf-go/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Config util.Config
	Router *gin.Engine
}

//// server serves out http request for our backend service

func NewServer(config util.Config) (*Server, error) {
	server := &Server{Config: config}
	router := gin.Default()
	server.Router = router

	groupRouter := router.Group("/api/v1")
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "app is unning fine at" + server.Config.HTTPServerAddress})
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":       "Resource not found",
			"route":       c.Request.URL.Path,
			"status_code": http.StatusNotFound,
		})
	})

	acctRepo := repository.NewAccountRepository()
	pkgSercice := pkg.NewPdfPackage()
	acctServices := services.NewAccountServices(acctRepo, pkgSercice)
	httphandlers.NewAccountRoutes(groupRouter, acctServices)

	return server, nil

}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

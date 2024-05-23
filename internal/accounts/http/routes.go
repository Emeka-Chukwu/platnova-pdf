package httphandlers

import (
	"pdf-go/internal/accounts/services"

	"github.com/gin-gonic/gin"
)

func NewAccountRoutes(router *gin.RouterGroup, accountService services.AccountServices) {
	accountHandler := NewAccountHandler(accountService)
	route := router.Group("/accounts")
	route.GET("/statements", accountHandler.FetchAccountStatement)
	route.GET("/statements/local", accountHandler.CreatePdfAndSave)
}

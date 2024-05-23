package httphandlers

import (
	"fmt"
	"net/http"
	domain "pdf-go/domain/account"
	"pdf-go/internal/accounts/services"
	"pdf-go/util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	templatePath         = "./templates/accounts_statement.html"
	accountStatementJson = "./account_statement.json"
)

type accountHandler struct {
	accountService services.AccountServices
}

type AccountHandler interface {
	FetchAccountStatement(ctx *gin.Context)
	CreatePdfAndSave(ctx *gin.Context)
}

func NewAccountHandler(accountService services.AccountServices) AccountHandler {
	return &accountHandler{accountService: accountService}
}

func (ah accountHandler) FetchAccountStatement(ctx *gin.Context) {
	currency := "usd"
	query := util.GetUrlQueryParams[domain.QueryParams](ctx)
	if strings.ToLower(query.Currency) == "eur" {
		currency = "eur"
	}
	status, pdf, err := ah.accountService.FetchAccountStatements(accountStatementJson, templatePath, currency)
	if err != nil {
		ctx.JSON(status, gin.H{"status": status, "error": err.Error()})
	}
	fileName := "account statement " + fmt.Sprintf("%d", time.Now().UnixNano()) + ".pdf"
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	ctx.Data(http.StatusOK, "application/pdf", pdf)
}

func (ah accountHandler) CreatePdfAndSave(ctx *gin.Context) {
	currency := "usd"
	query := util.GetUrlQueryParams[domain.QueryParams](ctx)
	if strings.ToLower(query.Currency) == "eur" {
		currency = "eur"
	}
	fileName := "account statement " + fmt.Sprintf("%d", time.Now().UnixNano()) + ".pdf"
	path := "storage/" + fileName
	err := ah.accountService.CreatePdfAndSaveOnLocal(accountStatementJson, templatePath, currency, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Successfully created and saved pdf"})
}

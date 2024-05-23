package services

import (
	"fmt"
	"net/http"
	"pdf-go/internal/accounts/repository"
	"pdf-go/pkg"
)

type accountServices struct {
	repo       repository.AccountRepository
	pkgService pkg.PdfPakage
}

type AccountServices interface {
	FetchAccountStatements(path string, htmlPath string, currency string) (int, []byte, error)
}

func NewAccountServices(repo repository.AccountRepository, pkgService pkg.PdfPakage) AccountServices {
	return &accountServices{repo: repo, pkgService: pkgService}
}

func (as *accountServices) FetchAccountStatements(path string, htmlPath string, currency string) (int, []byte, error) {
	var curencyMap map[string]string = map[string]string{
		"usd": "$", "eur": "€",
	}
	key, valid := curencyMap[currency]
	data, err := as.repo.FetchAccountStatement(path)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	if len(data) == 0 {
		return http.StatusAccepted, nil, fmt.Errorf("no data for the specified range")
	}
	pdfData := data[0]
	for _, value := range data {
		if key == value.Currency1 && valid {
			pdfData = value
		}
	}
	pdf, err := as.pkgService.ProcessPdf(pdfData, htmlPath)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, pdf, nil
}

package pkg

import (
	"bytes"
	"html/template"
	domain "pdf-go/domain/account"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type PdfPakage interface {
	ProcessPdf(model domain.AccountStatement, htmlPath string) ([]byte, error)
}
type pdfPakage struct {
}

func NewPdfPackage() PdfPakage {
	return &pdfPakage{}
}

func (pdfPakage) ProcessPdf(model domain.AccountStatement, htmlPath string) ([]byte, error) {
	var templ *template.Template
	templ, err := template.ParseFiles(htmlPath)
	if err != nil {
		return nil, err
	}
	var doc bytes.Buffer
	err = templ.Execute(&doc, model)
	if err != nil {
		return nil, err
	}
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginTop.Set(0)
	pdfg.MarginBottom.Set(0)
	pdfg.MarginLeft.Set(00)
	pdfg.MarginRight.Set(0)
	pdfg.Dpi.Set(50)
	pdfg.Cover.EnableLocalFileAccess.Set(true)
	pdfg.AddPage(wkhtmltopdf.NewPageReader(&doc))
	err = pdfg.Create()
	if err != nil {
		return nil, err
	}
	return pdfg.Bytes(), nil

}

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	//Capa
	pdf := newReport()
	pdf.AddPage()
	pdf = header(pdf)
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

func newReport() *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "")
	return pdf
}

func header(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {
	pdf.SetFont("Arial", "B", 12)
	pdf.SetFillColor(240, 240, 240)
	pdf.Cell(40, 10, fmt.Sprintf("Cliente Relatório de Revisão de Custos %s", time.Now().Format("dd/mm/YYYY")))
	return pdf
}

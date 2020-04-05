package api

import (
	"bytes"
	"io/ioutil"
	"log"
	"path"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"

	"github.com/shellbear/faktur/data"
)

// generateInvoice generate a PDF from invoice data.
// It will generate a file in the "InvoiceDir" folder and names it [UUID].pdf.
func generateInvoice(invoice *data.Invoice, conf *config) error {
	pdfGen, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	var buff bytes.Buffer
	if err := conf.Template.Execute(&buff, invoice); err != nil {
		return err
	}

	page := wkhtmltopdf.NewPageReader(ioutil.NopCloser(&buff))
	page.CacheDir.Set(conf.CacheDir)

	pdfGen.AddPage(page)
	if err = pdfGen.Create(); err != nil {
		return err
	}

	invoicePath := path.Join(conf.InvoiceDir, invoice.UUID.String()+".pdf")
	if err = pdfGen.WriteFile(invoicePath); err != nil {
		return err
	}

	return nil
}

// RunWorker starts the worker and waits for invoices.
func RunWorker(c chan data.Invoice, conf *config) {
	for {
		invoice := <-c

		log.Println("Processing invoice", invoice.UUID)

		if err := generateInvoice(&invoice, conf); err != nil {
			log.Printf("Failed to generate invoice %s. Error: %v\n", invoice.UUID, err)
		}

		log.Println("Successfully processed invoice", invoice.UUID)
	}
}

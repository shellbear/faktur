package api

import (
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
)

type config struct {
	// Number of workers to launch
	Workers uint8

	// Root folder of generated invoices
	InvoiceDir string

	// The default template to use for invoice generation
	TemplatePath string

	// The directory used to cache wkhtmltopdf web assets
	CacheDir string

	// Template content, read from Template path
	Template *template.Template
}

const (
	defaultWorkers    = "4"
	defaultInvoiceDir = "./invoices"
	defaultTemplate   = "./templates/invoice.html"
	defaultCacheDir   = "/tmp/cache-wk/"
)

// getConfig parse environment variables and will set a default value for each one if none is set.
func getConfig() (*config, error) {
	workers := os.Getenv("FAKTUR_WORKERS")
	if workers == "" {
		workers = defaultWorkers
	}

	invoiceDir := os.Getenv("FAKTUR_INVOICE_DIR")
	if invoiceDir == "" {
		invoiceDir = defaultInvoiceDir
	}

	templatePath := os.Getenv("FAKTUR_TEMPLATE")
	if templatePath == "" {
		templatePath = defaultTemplate
	}

	cacheDir := os.Getenv("FAKTUR_CACHE_DIR")
	if cacheDir == "" {
		cacheDir = defaultCacheDir
	}

	workersNb, err := strconv.ParseUint(workers, 10, 8)
	if err != nil {
		return nil, err
	}

	templateContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("default").Parse(string(templateContent))
	if err != nil {
		return nil, err
	}

	return &config{
		Workers:      uint8(workersNb),
		InvoiceDir:   invoiceDir,
		TemplatePath: templatePath,
		CacheDir:     cacheDir,
		Template:     tmpl,
	}, nil
}

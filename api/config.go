package api

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/markbates/pkger"
)

type config struct {
	// Number of workers to launch
	Workers uint8

	// Root folder of generated invoices
	InvoiceDir string

	// An optional directory for additional templates
	TemplateDir string

	// The directory used to cache wkhtmltopdf web assets
	CacheDir string

	// Templates content, read from Template path
	Template *template.Template
}

const (
	defaultWorkers    = "4"
	defaultInvoiceDir = "./invoices"
	defaultCacheDir   = "/tmp/cache-wk/"
)

func getFileName(filePath string) string {
	basename := path.Base(filePath)
	return strings.TrimSuffix(basename, path.Ext(basename))
}

func parseTemplate(t *template.Template, file io.Reader, path string) (*template.Template, error) {
	var tmpl *template.Template

	templateName := getFileName(path)
	if t == nil {
		t = template.New(templateName)
	}

	if t.Name() == templateName {
		tmpl = t
	} else {
		tmpl = t.New(templateName)
	}

	templateContent, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	_, err = tmpl.Parse(string(templateContent))
	return t, err
}

func getCustomTemplates(t *template.Template, dir string) (*template.Template, error) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}

			if t, err = parseTemplate(t, file, path); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return t, nil
}

func getDefaultTemplates(t *template.Template) (*template.Template, error) {
	err := pkger.Walk("/templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := pkger.Open(path)
			if err != nil {
				return err
			}

			if t, err = parseTemplate(t, file, path); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return t, nil
}

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

	cacheDir := os.Getenv("FAKTUR_CACHE_DIR")
	if cacheDir == "" {
		cacheDir = defaultCacheDir
	}

	workersNb, err := strconv.ParseUint(workers, 10, 8)
	if err != nil {
		return nil, err
	}

	tmpl, err := getDefaultTemplates(nil)
	if err != nil {
		return nil, err
	}

	templateDir := os.Getenv("FAKTUR_TEMPLATE_DIR")
	if templateDir != "" {
		if tmpl, err = getCustomTemplates(tmpl, templateDir); err != nil {
			return nil, err
		}
	}

	if tmpl != nil {
		for i, temp := range tmpl.Templates() {
			fmt.Printf("%d: %s\n", i, temp.Name())
		}
	}

	return &config{
		Workers:     uint8(workersNb),
		InvoiceDir:  invoiceDir,
		TemplateDir: templateDir,
		CacheDir:    cacheDir,
		Template:    tmpl,
	}, nil
}

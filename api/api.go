package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/shellbear/faktur/data"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() (*echo.Echo, error) {
	conf, err := getConfig()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(conf.InvoiceDir); os.IsNotExist(err) {
		if err := os.MkdirAll(conf.InvoiceDir, os.ModePerm); err != nil {
			return nil, err
		}
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Static("/invoices", conf.InvoiceDir)

	ch := make(chan data.Invoice)
	for i := uint8(0); i < conf.Workers; i++ {
		go RunWorker(ch, conf)
	}

	log.Println("Launched", conf.Workers, "workers")

	e.GET("/", func(c echo.Context) error {
		invoice := data.Invoice{}

		if err := invoice.Sanitize(); err != nil {
			return err
		}

		ch <- invoice

		return c.JSON(http.StatusAccepted, map[string]interface{}{
			"id":  invoice.UUID,
			"uri": fmt.Sprintf("/invoices/%s.pdf", invoice.UUID),
		})
	})

	e.POST("/", func(c echo.Context) error {
		var invoice data.Invoice

		if err := c.Bind(&invoice); err != nil {
			return err
		}

		if err := invoice.Sanitize(); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, &invoice)
	})

	return e, nil
}

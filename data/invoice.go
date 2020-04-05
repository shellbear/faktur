package data

import (
	"time"

	"github.com/google/uuid"
)

// Invoice contains all data needed to generate an invoice.
type Invoice struct {
	UUID uuid.UUID `json:"-"`

	// An unique invoice identifier.
	ID *string `json:"id"`

	// An optional title
	Title *string `json:"title"`

	// The invoice date, default to now
	Date time.Time `json:"date"`

	// An optional due date
	DueDate *time.Time `json:"due_date"`

	// A list of billed items
	Items []Item `json:"items"`

	// Information about your company
	Company *Person `json:"company"`

	// Information about the customer
	Customer *Person `json:"customer"`

	// The currency symbol, default to "$"
	Currency string `json:"currency"`

	// Some optional notes which will be added to the end of the invoice
	Notes *string `json:"notes"`

	Total float32 `json:"-"`
}

func (i *Invoice) Sanitize() (err error) {
	if i.UUID, err = uuid.NewRandom(); err != nil {
		return
	}

	if i.Date.IsZero() {
		i.Date = time.Now()
	}

	if i.Currency == "" {
		i.Currency = "$"
	}

	for index, item := range i.Items {
		i.Items[index].Total = item.Price * item.Quantity
		if item.Tax != 0 {
			i.Items[index].Total *= item.Tax
		}

		i.Total += i.Items[index].Total
	}

	return
}

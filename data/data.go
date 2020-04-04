package data

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    float32 `json:"quantity"`
}

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	County  string `json:"county"`
	Country string `json:"country"`
}

type Person struct {
	Logo    *string  `json:"logo"`
	Title   *string  `json:"title"`
	Phone   *string  `json:"phone"`
	Email   *string  `json:"email"`
	Address *Address `json:"address"`
}

// Invoice contains all data needed to generate an invoice.
type Invoice struct {
	UUID uuid.UUID `json:"-"`

	// An unique invoice identifier.
	ID *string `json:"id"`

	// An optional title
	Title *string `json:"title"`

	// The invoice date, default to now
	Date *time.Time `json:"date"`

	// An optional due date
	DueDate *time.Time `json:"due_date"`

	// A list of billed items
	Items []Item `json:"items"`

	// Information about your company
	Company Person `json:"company"`

	// Information about the customer
	Customer Person `json:"customer"`

	// The currency symbol, default to "$"
	Currency *string `json:"currency"`

	// Some optional notes which will be added to the end of the invoice
	Notes *string `json:"notes"`
}

func (i *Invoice) Sanitize() (err error) {
	if i.UUID, err = uuid.NewRandom(); err != nil {
		return
	}

	if i.Date == nil {
		now := time.Now()
		i.Date = &now
	}

	if i.Currency == nil {
		currency := "$"
		i.Currency = &currency
	}

	return
}

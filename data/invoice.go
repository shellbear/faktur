package data

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Calculate performs all dollar value calculations for each invoice.
type Calculate struct {
	SubTotal   float64 `json:"-"`
	TaxAmount  float64 `json:"-"`
	Total      float64 `json:"-"`
	AmountPaid float64 `json:"-"`
	BalanceDue float64 `json:"-"`
}

// Invoice contains data fields needed to generate a PDF Invoice.
type Invoice struct {
	UUID uuid.UUID `json:"-"`

	// An invoice identifier, currently set at 1001, increments +1 on next call
	ID *InvoiceId `json:"-"`

	// An optional title
	Title *string `json:"title"`

	// The invoice date, default to now
	Date time.Time `json:"date"`

	// An optional due date
	DueDate *time.Time `json:"due_date"`

	// A list of billed items
	Items []Item `json:"items"`

	// Information about your company
	Contractor *Contractor `json:"contractor"`

	// Information about the customer (individual or business)
	Customer *Customer `json:"customer"`

	// The currency symbol, default to "$"
	Currency string `json:"currency"`

	// Some optional notes which will be added to the end of the invoice
	Notes string `json:"-"`

	CalculateInvoice *Calculate `json:"-"`

	// TaxRate is the local tax rate
	TaxRate float64 `json:"-"`
}

var generator = NewIDGenerator() // Shared generator for Invoice instances

func (i *Invoice) Sanitize() (err error) {
	// creates a random uuid and sets the filename of each pdf as the generated uuid
	if i.UUID, err = uuid.NewRandom(); err != nil {
		return
	}

	// sets the date time
	if i.Date.IsZero() {
		i.Date = time.Now()
	}

	// default to USD
	if i.Currency == "" {
		i.Currency = "$"
	}

	// sets Invoice.Notes to Customer.Notes
	if i.Customer.Notes != "" {
		i.Notes = i.Customer.Notes
	}

	// assigns an Invoice.ID to each invoice generated.
	if i.ID == nil {
		i.ID = generator
	}

	i.ID.GenerateID()

	// string convert to float64 in order to satisfy Calculate struct
	amount_paid, err := strconv.ParseFloat(i.Customer.AmountPaid, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Calculate struct instantiation
	calculate_invoice := Calculate{
		SubTotal:   0.0,
		TaxAmount:  0.0,
		Total:      0.0,
		AmountPaid: amount_paid,
	}

	// iterate through slice, perform calculations and map to *Invoice fields
	for index, item := range i.Items {
		// Rounds the f64 to the nearest 100th position.
		i.Items[index].TaxAmount = math.Round(item.Rate*item.Quantity*item.TaxRate*100) / 100
		i.Items[index].SubTotal = item.Rate * item.Quantity

		calculate_invoice.SubTotal += i.Items[index].SubTotal
		calculate_invoice.TaxAmount += i.Items[index].TaxAmount
		calculate_invoice.Total = math.Round((calculate_invoice.SubTotal+calculate_invoice.TaxAmount)*100) / 100
		calculate_invoice.BalanceDue = calculate_invoice.Total - calculate_invoice.AmountPaid
	}

	i.CalculateInvoice = &calculate_invoice // assign values into Calculate

	return
}

type InvoiceId struct {
	sync.Mutex
	FakturInvoiceId int `json:"-"`
}

func NewIDGenerator() *InvoiceId {
	return &InvoiceId{
		FakturInvoiceId: 1000, // starting invoice id is 1000
	}
}

func (a *InvoiceId) GenerateID() int {
	a.Lock()
	defer a.Unlock()

	idToReturn := a.FakturInvoiceId
	a.FakturInvoiceId++

	return idToReturn
}

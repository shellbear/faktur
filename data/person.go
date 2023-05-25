package data

// Customer represents some individual or business customer.
type Customer struct {
	// The company or person name
	FullName   string  `json:"full_name"`
	Phone      string  `json:"phone"`
	Email      string  `json:"email"`
	Address    Address `json:"address"`
	Terms      string  `json:"terms"`
	AmountPaid string  `json:"amount_paid"`

	// Customer notes go here! these notes are also mapped to Invoice notes.
	Notes string `json:"customer_notes"`
}

// Corporate represents the 'contractor', entity issuing the invoice.
type Contractor struct {
	// A logo, can be an URL or a base64 encoded image string
	Logo Image `json:"logo"`

	// The business legal name
	BusinessName       string  `json:"business_name"`
	Phone              string  `json:"phone"`
	Email              string  `json:"email"`
	RegistrationNumber string  `json:"registration_number"`
	Address            Address `json:"address"`

	// Some optional additional notes
	Notes *string `json:"notes"`
}

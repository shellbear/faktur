package data

// Person represents someone, can be your company or the customer.
type Person struct {
	// A logo, can be an URL or a base64 encoded image string
	Logo Image `json:"logo"`

	// The company or person name
	Title              string  `json:"title"`
	Phone              string  `json:"phone"`
	Email              string  `json:"email"`
	RegistrationNumber string  `json:"registration_number"`
	Address            Address `json:"address"`

	// Some optional additional notes
	Notes *string `json:"notes"`
}

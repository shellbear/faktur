package data

// Person represents someone, can be your company or the customer.
type Person struct {
	Logo    Image   `json:"logo"`
	Title   string  `json:"title"`
	Phone   string  `json:"phone"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

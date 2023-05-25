package data

// Item is a product or a service that is billed to customer.
type Item struct {
	ItemNumber  string  `json:"item_number"`
	Description string  `json:"description"`
	Rate        float64 `json:"rate"`
	Quantity    float64 `json:"quantity"`
	TaxAmount   float64 `json:"-"` // taxable amount in USD for each row item.
	TaxRate     float64 `json:"tax_rate"`
	SubTotal    float64 `json:"-"` // (Quantity * Rate) for each row item.
}

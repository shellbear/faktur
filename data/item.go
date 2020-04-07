package data

// Item is a product or a service that is billed to customer.
type Item struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    float32 `json:"quantity"`
	Tax         float32 `json:"tax"`
	Total       float32 `json:"-"`
}

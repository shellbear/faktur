package data

// Item is a billing item.
type Item struct {
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	Quantity    float32 `json:"quantity"`
	Description string  `json:"description"`
	Tax         float32 `json:"tax"`
	Total       float32 `json:"-"`
}

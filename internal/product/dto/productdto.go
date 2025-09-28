package dto

type ProductResponse struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	ImageUrl   string  `json:"image_url"`
	Price      float64 `json:"price"`
	CategoryID string  `json:"category_id"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

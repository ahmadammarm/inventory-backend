package dto

type CategoryResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ProductResponse struct {
	ID                string           `json:"id"`
	Name              string           `json:"name"`
	ImageUrl          string           `json:"image_url"`
	Price             float64          `json:"price"`
	ProductCategoryID string           `json:"product_category_id"`
	ProductCategory   CategoryResponse `json:"product_category"`
	CreatedAt         string           `json:"created_at"`
	UpdatedAt         string           `json:"updated_at"`
}

package pkg

import "errors"

type Product struct {
	Id          int    `db:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Price       int64  `json:"price"` // in cents
	CategoryId  int    `json:"category_id" db:"category_id"`
}
type UpdateProductInput struct {
	Name        *string `json:"name" binding:"required"`
	Description *string `json:"description"`
	Price       *int64  `json:"price"` // in cents
	CategoryId  *int    `json:"category_id" db:"category_id`
}

func (i *UpdateProductInput) Validate() error {
	if i.Name == nil && i.Description == nil && i.Price == nil && i.CategoryId == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

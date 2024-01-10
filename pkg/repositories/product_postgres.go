package repositories

import (
	"fmt"
	"github.com/fdrt29/product-app/pkg"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(product pkg.Product) (int, error) {
	var id int
	createProductQuery := `INSERT INTO products (name, description, price, category_id, created_at) 
VALUES ($1, $2, $3, $4, $5) RETURNING id`
	row := r.db.QueryRow(createProductQuery, product.Name, product.Description, product.Price, product.CategoryId, time.Now())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductPostgres) GetAll() ([]pkg.Product, error) {
	var products []pkg.Product
	query := "SELECT id, name, description, price, category_id FROM products"
	err := r.db.Select(&products, query)
	return products, err
}

func (r *ProductPostgres) GetById(productId int) (pkg.Product, error) {
	var product pkg.Product
	query := "SELECT id, name, description, price, category_id FROM products WHERE id = $1"
	err := r.db.Get(&product, query, productId)
	return product, err
}

func (r *ProductPostgres) Update(productId int, input pkg.UpdateProductInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}
	if input.CategoryId != nil {
		setValues = append(setValues, fmt.Sprintf("category_id=$%d", argId))
		args = append(args, *input.CategoryId)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE products SET %s WHERE id = $%d", setQuery, argId)
	args = append(args, productId)
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *ProductPostgres) Delete(productId int) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := r.db.Exec(query, productId)
	return err
}

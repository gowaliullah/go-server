package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/gowaliullah/basic-ecommerce/domain"
	"github.com/gowaliullah/basic-ecommerce/product"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db sqlx.DB
}

// constructor or constructor functions
func NewProductRepo(db sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {

	query := `
		INSERT INTO products (
			title,
			description,
			price,
			image_url
		) VALUES (
			$1,
			$2,
			$3,
			$4
		)
			RETURNING id
	`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImageUrl)

	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil

}

func (r *productRepo) Get(id int) (*domain.Product, error) {
	var prd domain.Product
	query := `
		SELECT 
			id, 
			title,
			description,
			price,
			image_url
		FROM products
		WHERE id = $1
	`
	err := r.db.Get(&prd, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &prd, nil
}

func (r *productRepo) List(page, limit int64) ([]*domain.Product, error) {

	offset := ((page - 1) * limit) + 1

	var prdlist []*domain.Product
	query := `
		SELECT 
			id, 
			title,
			description,
			price,
			image_url
		FROM products
		LIMIT $1
		OFFSET $2	
	`
	err := r.db.Select(&prdlist, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return prdlist, nil
}

func (r *productRepo) Delete(id int) error {

	query := `DELETE FROM products WHERE id =$1`
	_, err := r.db.Exec(query, id)

	if err != nil {
		return nil
	}

	return nil
}

func (r *productRepo) Update(prd domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products SET
			title=$1,
			description=$2,
			price=$3,
			image_url=$4
		WHERE id=$5
	`
	row := r.db.QueryRow(query, prd.Title, prd.Description, prd.Price, prd.ImageUrl, prd.ID)
	err := row.Err()
	if err != nil {
		return nil, err
	}

	return &prd, nil
}

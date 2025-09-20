package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgURL      string  `json:"imageUrl" db:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(prdId int) (*Product, error)
	List() ([]*Product, error)
	Delete(prdId int) error
	Update(p Product) (*Product, error)
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

func (r *productRepo) Create(p Product) (*Product, error) {

	query := `
		INSERT INTO products (
			title,
			description,
			price,
			imageUrl
		) VALUES (
			$1,
			$2,
			$3,
			$4 
		)
			RETURNING id
	`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgURL)

	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil

}

func (r *productRepo) Get(prdId int) (*Product, error) {
	var prd Product
	query := `
		SELECT 
			id, 
			title,
			description,
			price,
			imageUrl
		FROM products
		WHERE id = $1
	`
	err := r.db.Get(&prd, query, prdId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &prd, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Delete(prdId int) error {
	var tempList []*Product

	for _, product := range r.productList {
		if product.ID != prdId {
			tempList = append(tempList, product)
		}
	}
	r.productList = tempList

	return nil
}

func (r *productRepo) Update(prd Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == prd.ID {
			r.productList[idx] = &prd // productList[0]
		}
	}

	return &prd, nil
}

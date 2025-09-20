package repo

import "github.com/jmoiron/sqlx"

type Product struct {
	ID          int     `json:"_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
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
	for _, product := range r.productList {
		if product.ID == prdId {
			return product, nil
		}
	}
	return nil, nil
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

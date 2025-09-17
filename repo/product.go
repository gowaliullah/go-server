package repo

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
	List() []*Product
	Delete(prdId int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

// constructor or constructor functions
func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {

	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
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

func (r *productRepo) Update() {

}

func generateInitialProducts(r *productRepo) {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "It's very delicious and full of vitamin C.",
		Price:       108.00,
		ImgURL:      "https://example.com/images/orange.jpg",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Crisp and sweet red apples.",
		Price:       120.50,
		ImgURL:      "https://example.com/images/apple.jpg",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Rich in potassium and easy to digest.",
		Price:       60.00,
		ImgURL:      "https://example.com/images/banana.jpg",
	}

	prd4 := Product{
		ID:          4,
		Title:       "Mango",
		Description: "The king of fruits, sweet and juicy.",
		Price:       150.75,
		ImgURL:      "https://example.com/images/mango.jpg",
	}

	prd5 := Product{
		ID:          5,
		Title:       "Watermelon",
		Description: "Refreshing and hydrating summer fruit.",
		Price:       200.00,
		ImgURL:      "https://example.com/images/watermelon.jpg",
	}

	r.productList = append(r.productList, &prd1, &prd2, &prd3, &prd4, &prd5)
}

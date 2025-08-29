package database

var productList []Product

type Product struct {
	ID          int     `json:"_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(prdId int) *Product {
	for _, product := range productList {
		if product.ID == prdId {
			return &product
		}
	}
	return nil
}

func Update(prd Product) {
	for idx, p := range productList {
		if p.ID == prd.ID {
			productList[idx] = prd
		}
	}
}

func Delete(prdId int) {
	var tempList []Product = make([]Product, 0)

	for _, product := range productList {
		if product.ID != prdId {
			tempList = append(tempList, product)
		}
	}
	productList = tempList
}

func init() {
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

	productList = append(productList, prd1, prd2, prd3, prd4, prd5)
}

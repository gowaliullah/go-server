package repo

type ProductRepo interface {
	Create()
	Get()
	List()
	Delete()
	Update()
}

package product

type service struct {
	prductRepo ProductRepo
}

func NewService(prductRepo ProductRepo) Service {
	return &service{
		prductRepo: prductRepo,
	}
}

package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name string, productType string, count int, price float64) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Store(name string, productType string, count int, price float64) (Product, error) {
	product, err := s.repository.Store(s.repository.GetNewId(), name, productType, count, price)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

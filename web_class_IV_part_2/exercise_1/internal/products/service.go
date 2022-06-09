package products

type Service interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Save(name string, productType string, count int, price float64) (Product, error)
	Delete(id int) error
	Update(id int, name string, productType string, count int, price float64) (Product, error)
	PatchNamePrice(id int, name string, price float64) (Product, error)
}

type service struct {
	repository Repository
}

func NewService() Service {
	return &service{repository: NewRepository()}
}

func (s *service) Save(name string, productType string, count int, price float64) (Product, error) {
	product, err := s.repository.Save(name, productType, count, price)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Update(id int, name string, productType string, count int, price float64) (Product, error) {
	product, err := s.repository.Update(id, name, productType, count, price)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) PatchNamePrice(id int, name string, price float64) (Product, error) {
	return s.repository.PatchNamePrice(id, name, price)
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) GetById(id int) (Product, error) {
	return s.repository.GetById(id)
}

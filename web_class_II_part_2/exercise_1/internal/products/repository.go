package products

type Product struct {
	Id    int
	Name  string
	Type  string
	Count int
	Price float64
}

var products []Product

type repository struct{}

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name string, productType string, count int, price float64) (Product, error)
	GetNewId() int
}

func NewRepository() Repository {
	return &repository{}
}

func (repository) Store(id int, name string, productType string, count int, price float64) (Product, error) {
	p := Product{id, name, productType, count, price}
	products = append(products, p)
	return p, nil
}

func (repository) GetAll() ([]Product, error) {
	return products, nil
}

func (repository) GetNewId() int {
	var newID int
	if len(products) <= 0 {
		newID = 1
	} else {
		newID = products[len(products)-1].Id + 1
	}
	return newID
}

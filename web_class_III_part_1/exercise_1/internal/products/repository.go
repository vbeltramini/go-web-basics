package products

import "fmt"

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
	Save(id int, name string, productType string, count int, price float64) (Product, error)
	GetNewId() int
	Update(id int, name string, productType string, count int, price float64) (Product, error)
}

func NewRepository() Repository {
	return &repository{}
}

func (repository) Save(id int, name string, productType string, count int, price float64) (Product, error) {
	newProduct := Product{id, name, productType, count, price}
	products = append(products, newProduct)

	return newProduct, nil
}

func (repository) Update(id int, name string, productType string, count int, price float64) (Product, error) {
	newProduct := Product{id, name, productType, count, price}
	for i := range products {
		if newProduct.Id == id {
			products[i] = newProduct
			return newProduct, nil
		}
	}
	return Product{}, fmt.Errorf("products with id: %d not found", id)
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

package products

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int     `json:"count"`
	Price float64 `json:"price"`
}

var products []Product
var baseId = 7

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Save(name string, productType string, count int, price float64) (Product, error)
	Delete(id int) (bool, error)
	Update(id int, name string, productType string, count int, price float64) (Product, error)
	PatchNamePrice(id int, name string, price float64) (Product, error)
}

type repository struct{}

func NewRepository() Repository {
	products = createSimpleProductsData()
	return &repository{}
}

func (repository) Save(name string, productType string, count int, price float64) (Product, error) {
	newProduct := Product{getNewId(), name, productType, count, price}
	products = append(products, newProduct)

	return newProduct, nil
}

func (repository) Update(id int, name string, productType string, count int, price float64) (Product, error) {
	newProduct := Product{id, name, productType, count, price}
	fmt.Println(id)
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].Id)
		if products[i].Id == id {
			products[i] = newProduct
			return newProduct, nil
		}
	}

	return Product{}, fmt.Errorf("product with id: %d not found", id)
}

func (repository) PatchNamePrice(id int, name string, price float64) (Product, error) {
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].Id)
		if products[i].Id == id {
			products[i].Name = name
			products[i].Price = price
			return products[i], nil
		}
	}

	return Product{}, fmt.Errorf("product with id: %d not found", id)
}

func (repository) Delete(id int) (bool, error) {
	for i := 0; i < len(products); i++ {
		if products[i].Id == id {
			products = append(products[:i], products[i+1:]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("product with id: %d not found", id)
}

func (repository) GetAll() ([]Product, error) {
	return products, nil
}

func (repository) GetById(id int) (Product, error) {
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].Id)
		if products[i].Id == id {
			return products[i], nil
		}
	}

	return Product{}, fmt.Errorf("product with id: %d not found", id)
}

func getNewId() int {
	baseId++
	return baseId
}

func createSimpleProductsData() []Product {
	jsonData := []byte(`
    [
        {"id":1,"name": "batata doce??","type": "Batata","count": 10,"price": 1.99},
        {"id":2,"name": "teste","type": "Batata","count": 10,"price": 1.99},
        {"id":3,"name": "batata doce??","type": "Batata","count": 10,"price": 1.99},
        {"id":4,"name": "batata doce??","type": "Batata","count": 10,"price": 1.99},
        {"id":5,"name": "batata doce??","type": "Batata","count": 10,"price": 1.99},
        {"id":6,"name": "batata doce??","type": "Batata","count": 10,"price": 1.99},
        {"id":7,"name": "batata doce??","type": "Batata","count": 10,"price": 1.99}
    ]`)

	var u []Product

	if err := json.Unmarshal(jsonData, &u); err != nil {
		log.Fatal(err)
	}
	return u
}

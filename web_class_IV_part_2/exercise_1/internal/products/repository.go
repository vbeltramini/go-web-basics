package products

import (
	"fmt"

	"github.com/vbeltramini/go-web-basics/web_class_IV_part_2/exercise_1/pkg/store"
)

var products []Product
var baseId = 7

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Save(name string, productType string, count int, price float64) (Product, error)
	Delete(id int) error
	Update(id int, name string, productType string, count int, price float64) (Product, error)
	PatchNamePrice(id int, name string, price float64) (Product, error)
}

type repository struct {
	db store.Store
}

func NewRepository() Repository {
	db := store.New(store.FileType, "../../pkg/store/data.json")
	return &repository{db: db}
}

func (r *repository) Save(name string, productType string, count int, price float64) (Product, error) {
	var products []Product
	r.db.Read(&products)
	newProduct := Product{getNewId(), name, productType, count, price}

	products = append(products, newProduct)
	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}
	return newProduct, nil
}

func (r *repository) Update(id int, name string, productType string, count int, price float64) (Product, error) {

	var products []Product
	r.db.Read(&products)

	index := -1
	for i := range products {
		if products[i].Id == id {
			index = i
		}
	}

	if index != -1 {
		products[index] = Product{id, name, productType, count, price}
		r.db.Write(products)
		return products[index], nil
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

func (r *repository) Delete(id int) error {
	var products []Product
	r.db.Read(&products)

	for i := range products {
		if products[i].Id == id {
			products = append(products[:i], products[i+1:]...)
			r.db.Write(products)
			return nil
		}
	}
	return fmt.Errorf("product with id: %d not found", id)
}

func (r repository) GetAll() ([]Product, error) {
	var products []Product
	r.db.Read(&products)

	if len(products) == 0 {
		return make([]Product, 0), nil
	}

	return products, nil
}

func (r repository) GetById(id int) (Product, error) {

	var productList []Product
	r.db.Read(&productList)

	productIndex := -1

	for i := range productList {
		if productList[i].Id == id {
			productIndex = i
		}
	}

	if productIndex == -1 {
		return Product{}, fmt.Errorf("product with id %d not founded", id)
	}
	return productList[productIndex], nil
}

func getNewId() int {
	baseId++
	return baseId
}

package store

import (
	"app/internal"
	"encoding/json"
	"os"
	"time"
)

// NewStoreProductJSON creates a new JSON file store for products.
func NewStoreProductJSON(path string) (s *StoreProductJSON) {
	s = &StoreProductJSON{
		Path: path,
	}
	return
}

// StoreProductJSON is a JSON file store for products.
type StoreProductJSON struct {
	// Path is the path to the JSON file.
	Path string
}

// ProductJSON is a JSON representation of a product.
type ProductJSON struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// ReadAll reads all products from the store.
func (s *StoreProductJSON) ReadAll() (p map[int]internal.Product, err error) {
	// open file
	f, err := os.Open(s.Path)
	if err != nil {
		return
	}
	defer f.Close()

	// decode JSON
	var pr []ProductJSON
	err = json.NewDecoder(f).Decode(&pr)
	if err != nil {
		return
	}

	// serialize
	p = make(map[int]internal.Product)
	for _, v := range pr {
		var exp time.Time
		exp, err = time.Parse(time.DateOnly, v.Expiration)
		if err != nil {
			return
		}

		p[v.Id] = internal.Product{
			Id:          v.Id,
			ProductAttributes: internal.ProductAttributes{
				Name:        v.Name,
				Quantity:    v.Quantity,
				CodeValue:   v.CodeValue,
				IsPublished: v.IsPublished,
				Expiration:  exp,
				Price:       v.Price,
			},
		}
	}

	return
}

// WriteAll writes all products to the store.
func (s *StoreProductJSON) WriteAll(p map[int]internal.Product) (err error) {
	// serialize
	var pr []ProductJSON
	for _, v := range p {
		pr = append(pr, ProductJSON{
			Id:          v.Id,
			Name:        v.Name,
			Quantity:    v.Quantity,
			CodeValue:   v.CodeValue,
			IsPublished: v.IsPublished,
			Expiration:  v.Expiration.Format(time.DateOnly),
			Price:       v.Price,
		})
	}

	// open file
	// - create if not exists / write only / truncate
	f, err := os.OpenFile(s.Path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	// encode JSON
	err = json.NewEncoder(f).Encode(pr)
	if err != nil {
		return
	}

	return
}
package repository

import "app/internal"

// NewProductMap returns a new ProductMap
func NewProductMap(db map[int]internal.Product, lastId int) (pm *ProductMap) {
	// default config
	defaultDb := make(map[int]internal.Product)
	defaultLastId := 0
	if db != nil {
		defaultDb = db
	}
	if lastId != 0 {
		defaultLastId = lastId
	}

	pm = &ProductMap{
		db:     defaultDb,
		lastId: defaultLastId,
	}
	return
}

// ProductMap is a map that represents a product repository
type ProductMap struct {
	// db is a map that represents a database
	db map[int]internal.Product
	// lastId is an int that represents the last id of the database
	lastId int
}

// Save saves a product
func (pm *ProductMap) Save(p *internal.Product) (err error) {
	// check if the product has a duplicated field
	for _, pr := range pm.db {
		if pr.Name == p.Name {
			err = internal.ErrProductDuplicatedField
			return
		}
	}

	// set the id of the product
	pm.lastId++
	(*p).Id = pm.lastId

	// save the product
	pm.db[pm.lastId] = *p

	return
}
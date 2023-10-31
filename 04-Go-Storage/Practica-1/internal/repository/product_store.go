package repository

import "app/internal"

// NewRepositoryProductStore creates a new repository for products.
func NewRepositoryProductStore(st internal.StoreProduct) (r *RepositoryProductStore) {
	r = &RepositoryProductStore{
		st: st,
	}
	return
}

// RepositoryProductStore is a repository for products.
type RepositoryProductStore struct {
	// st is the underlying store.
	st internal.StoreProduct
}

// FindById finds a product by id.
func (r *RepositoryProductStore) FindById(id int) (p internal.Product, err error) {
	// read all products
	ps, err := r.st.ReadAll()
	if err != nil {
		return
	}

	// find product
	p, ok := ps[id]
	if !ok {
		err = internal.ErrRepositoryProductNotFound
		return
	}

	return
}

// Save saves a product.
func (r *RepositoryProductStore) Save(p *internal.Product) (err error) {
	// read all products
	ps, err := r.st.ReadAll()
	if err != nil {
		return
	}

	// find max id
	var maxId int
	var cc int
	for k := range ps {
		cc++
		if cc == 1 {
			maxId = k
			continue
		}
		if k > maxId {
			maxId = k
		}
	}

	// set id
	(*p).Id = maxId + 1

	// add product
	ps[p.Id] = *p

	// write all products
	err = r.st.WriteAll(ps)
	if err != nil {
		return
	}

	return
}

// UpdateOrSave updates or saves a product.
func (r *RepositoryProductStore) UpdateOrSave(p *internal.Product) (err error) {
	// read all products
	ps, err := r.st.ReadAll()
	if err != nil {
		return
	}

	// update product
	_, ok := ps[p.Id]
	switch ok {
	case true:
		ps[p.Id] = *p
	default:
		// find max id
		var maxId int
		var cc int
		for k := range ps {
			cc++
			if cc == 1 {
				maxId = k
				continue
			}
			if k > maxId {
				maxId = k
			}
		}

		// set id
		(*p).Id = maxId + 1

		// add product
		ps[p.Id] = *p
	}

	// write all products
	err = r.st.WriteAll(ps)
	if err != nil {
		return
	}

	return
}

// Update updates a product.
func (r *RepositoryProductStore) Update(p *internal.Product) (err error) {
	// read all products
	ps, err := r.st.ReadAll()
	if err != nil {
		return
	}

	// update product
	_, ok := ps[p.Id]
	if !ok {
		err = internal.ErrRepositoryProductNotFound
		return
	}

	// update product
	ps[p.Id] = *p

	// write all products
	err = r.st.WriteAll(ps)
	if err != nil {
		return
	}

	return
}

// Delete deletes a product.
func (r *RepositoryProductStore) Delete(id int) (err error) {
	// read all products
	ps, err := r.st.ReadAll()
	if err != nil {
		return
	}

	// delete product
	_, ok := ps[id]
	if !ok {
		err = internal.ErrRepositoryProductNotFound
		return
	}

	// delete product
	delete(ps, id)

	// write all products
	err = r.st.WriteAll(ps)
	if err != nil {
		return
	}

	return
}
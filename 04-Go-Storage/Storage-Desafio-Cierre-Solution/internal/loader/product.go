package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

// NewLoaderProductJSON returns a new pointer to a LoaderProductJSON struct.
func NewLoaderProductJSON(file *os.File) *LoaderProductJSON {
	return &LoaderProductJSON{file: file}
}

// LoaderProductJSON is an struct that implements the LoaderProduct interface.
type LoaderProductJSON struct {
	// file is the file to handle read and write operations.
	file *os.File
}

// ProductJSON is the struct that represents the product data in the json file.
type ProductJSON struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Load loads the product data from the json file.
func (l *LoaderProductJSON) Load() (p []internal.Product, err error) {
	// decode the json file
	var ps []ProductJSON
	err = json.NewDecoder(l.file).Decode(&ps)
	if err != nil {
		return
	}

	// serialize the product data
	for _, v := range ps {
		p = append(p, internal.Product{
			Id: v.Id,
			ProductAttributes: internal.ProductAttributes{
				Description: v.Description,
				Price:       v.Price,
			},
		})
	}

	return
}

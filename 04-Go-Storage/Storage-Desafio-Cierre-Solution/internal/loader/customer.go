package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

// NewLoaderCustomerJSON returns a new pointer to a LoaderCustomerJSON struct.
func NewLoaderCustomerJSON(file *os.File) *LoaderCustomerJSON {
	return &LoaderCustomerJSON{file: file}
}

// LoaderCustomerJSON is an struct that implements the LoaderCustomer interface.
type LoaderCustomerJSON struct {
	// file is the file to handle read and write operations.
	file *os.File
}

// CustomerJSON is the struct that represents the customer data in the json file.
type CustomerJSON struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

// Load loads the customer data from the json file.
func (l *LoaderCustomerJSON) Load() (c []internal.Customer, err error) {
	// decode the json file
	var cs []CustomerJSON
	err = json.NewDecoder(l.file).Decode(&cs)
	if err != nil {
		return
	}

	// serialize the customer data
	for _, v := range cs {
		c = append(c, internal.Customer{
			Id: v.Id,
			CustomerAttributes: internal.CustomerAttributes{
				FirstName: v.FirstName,
				LastName:  v.LastName,
				Condition: v.Condition,
			},
		})
	}

	return
}

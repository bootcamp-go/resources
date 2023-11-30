package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

// NewCustomersJSON returns a new pointer to a CustomersJSON struct.
func NewCustomersJSON(file *os.File) *CustomersJSON {
	return &CustomersJSON{file: file}
}

// CustomersJSON is an struct that implements the LoaderCustomer interface.
type CustomersJSON struct {
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
func (l *CustomersJSON) Load() (c []internal.Customer, err error) {
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

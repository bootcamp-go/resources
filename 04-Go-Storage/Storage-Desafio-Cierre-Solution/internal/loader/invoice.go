package loader

import (
	"app/internal"
	"encoding/json"
	"os"
)

// NewLoaderInvoiceJSON returns a new pointer to a LoaderInvoiceJSON struct.
func NewLoaderInvoiceJSON(file *os.File) *LoaderInvoiceJSON {
	return &LoaderInvoiceJSON{file: file}
}

// LoaderInvoiceJSON is an struct that implements the LoaderInvoice interface.
type LoaderInvoiceJSON struct {
	// file is the file to handle read and write operations.
	file *os.File
}

// InvoiceJSON is the struct that represents the invoice data in the json file.
type InvoiceJSON struct {
	Id         int    `json:"id"`
	Datetime   string `json:"datetime"`
	Total	  float64 `json:"total"`
	CustomerId int    `json:"customer_id"`
}

// Load loads the invoice data from the json file.
func (l *LoaderInvoiceJSON) Load() (i []internal.Invoice, err error) {
	// decode the json file
	var is []InvoiceJSON
	err = json.NewDecoder(l.file).Decode(&is)
	if err != nil {
		return
	}

	// serialize the invoice data
	for _, v := range is {
		i = append(i, internal.Invoice{
			Id: v.Id,
			InvoiceAttributes: internal.InvoiceAttributes{
				Datetime:   v.Datetime,
				Total:		v.Total,
				CustomerId: v.CustomerId,
			},
		})
	}

	return
}
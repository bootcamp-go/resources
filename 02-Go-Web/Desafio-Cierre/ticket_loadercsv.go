package loader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (t *LoaderTicketCSV) Load() (t map[int]internal.TicketAttributes, err error) {
	// open the file
	f, err := os.Open(t.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records
	t := make(map[int]internal.TicketAttributes)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		
			err = fmt.Errorf("error reading record: %v", err)
			return
		}

		// serialize the record
		id := record[0]
		ticket := internal.TicketAttributes{
			Name: record[1].(string),
			Email: record[2].(string),
			Country: record[3].(string),
			Hour: record[4].(string),
			Price: record[5].(int),
		}

		// add the ticket to the map
		t[id] = ticket
	}

	return
}


	
	


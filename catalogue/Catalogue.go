package catalogue

import (
	"encoding/csv"
	"io"
	"strconv"
)

var Catalogue map[string]*Item

// LoadCatalogue reads and parses the products from catalogue.txt
func Load(bufioReader io.Reader) {
	Catalogue = make(map[string]*Item)

	reader := csv.NewReader(bufioReader)
	for {
		record, err1 := reader.Read()
		if err1 == io.EOF {
			break
		}
		price, err1 := strconv.ParseFloat(record[2], 32)
		if err1 != nil {
			panic(err1)
		}
		Catalogue[record[0]] = &Item{record[0], record[1], float32(price)}
	}
}

// Item - a single product
type Item struct {
	ItemCode string
	ItemName string
	Price    float32
}

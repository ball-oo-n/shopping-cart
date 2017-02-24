package catalogue

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

var Catalogue map[string]*Item

func init() {
	Catalogue = LoadCatalogue()
}

// LoadCatalogue reads and parses the products from catalogue.txt
func LoadCatalogue() map[string]*Item {
	PriceMap := make(map[string]*Item)

	if file, err := os.Open("../main/catalogue.txt"); err == nil {
		defer file.Close()
		reader := csv.NewReader(bufio.NewReader(file))
		for {
			record, err1 := reader.Read()
			if err1 == io.EOF {
				break
			}
			price, err1 := strconv.ParseFloat(record[2], 32)
			if err1 != nil {
				panic(err)
			}
			PriceMap[record[0]] = &Item{record[0], record[1], float32(price)}
		}
	} else {
		log.Fatal(err)
	}
	return PriceMap
}

// Item - a single product
type Item struct {
	ItemCode string
	ItemName string
	Price    float32
}

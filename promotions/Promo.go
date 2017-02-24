package promotions

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

//PromoCodes key is the promo code, value is the decimal discount
var PromoCodes map[string]float32

func init() {
	PromoCodes = LoadPromoCodes()
}

//LoadPromoCodes loads the promo codes
func LoadPromoCodes() map[string]float32 {
	pmap := make(map[string]float32)
	//----
	if file, err := os.Open("../main/promo.txt"); err == nil {
		defer file.Close()
		reader := csv.NewReader(bufio.NewReader(file))
		for {
			record, err1 := reader.Read()
			if err1 == io.EOF {
				break
			}
			discount, err1 := strconv.ParseFloat(record[1], 32)
			pmap[record[0]] = float32(discount)
			if err1 != nil {
				panic(err)
			}
		}
	} else {
		log.Fatal(err)
	}
	//----
	return pmap
}

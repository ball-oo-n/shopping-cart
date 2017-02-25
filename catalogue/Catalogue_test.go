package catalogue_test

import (
	"bufio"
	"os"
	"testing"

	. "github.com/ball-oo-n/shopping-cart/catalogue"
)

func init() {
	file, _ := os.Open("../catalogue.txt")
	defer file.Close()
	Load(bufio.NewReader(file))
}

func TestLoadCatalogue(t *testing.T) {
	expMap := map[string]Item{
		"ult_small":  Item{"ult_small", "Unlimited 1GB", 24.90},
		"ult_medium": Item{"ult_medium", "Unlimited 2GB", 29.90},
		"ult_large":  Item{"ult_large", "Unlimited 5GB", 44.90},
		"1gb":        Item{"1gb", "1 GB Data-pack", 9.90},
	}
	for code := range expMap {
		if actItem := Catalogue[code]; actItem != nil {
			if actItem.ItemCode != expMap[code].ItemCode {
				t.Errorf("Test FAILED. Actual: %v Expected: %v", actItem.ItemCode, expMap[code].ItemCode)
			}
			if actItem.ItemName != expMap[code].ItemName {
				t.Errorf("Test FAILED. Actual: %v Expected: %v", actItem.ItemName, expMap[code].ItemName)
			}
			if actItem.Price != expMap[code].Price {
				t.Errorf("Test FAILED. Actual: %v Expected: %v", actItem.Price, expMap[code].Price)
			}
		} else {
			t.Error("Test FAILED. Actual: nil Expected: not nil")
		}
	}
}

func TestPrice(t *testing.T) {
	const mapLen int = 4
	expPrice := [mapLen]float32{24.90, 29.90, 44.90, 9.90}
	testItmCode := [mapLen]string{"ult_small", "ult_medium", "ult_large", "1gb"}

	for i := 0; i < mapLen; i++ {
		actPrice := Catalogue[testItmCode[i]].Price
		if actPrice != expPrice[i] {
			t.Errorf("Test FAILED. Actual: %v Expected: %v", testItmCode[i], expPrice[i])
		}
	}
}

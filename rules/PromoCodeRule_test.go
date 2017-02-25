package rules

import (
	"bufio"
	"os"
	"testing"

	"github.com/ball-oo-n/shopping-cart/catalogue"
	"github.com/ball-oo-n/shopping-cart/shop"
)

func init() {
	file, _ := os.Open("../catalogue.txt")
	defer file.Close()
	catalogue.Load(bufio.NewReader(file))
}

//Case1:
func TestPromoCodeRule1(t *testing.T) {

	//test data
	items := map[string]*shop.Item{
		"ult_medium": &shop.Item{
			Quantity:   1,
			TotalPrice: 29.90,
		},
		"ult_large": &shop.Item{
			Quantity:   4,
			TotalPrice: (44.90 * 4),
		},
		"I<3AMAYSIM": &shop.Item{
			Quantity: 1,
		},
	}
	//expected data
	expItems := map[string]shop.Item{
		"ult_medium": shop.Item{
			Quantity:   1,
			TotalPrice: 29.90 - (29.90 * 0.1),
		},
		"ult_large": shop.Item{
			Quantity:   4,
			TotalPrice: (44.90 * 4) - ((44.90 * 4) * 0.1),
		},
		"I<3AMAYSIM": shop.Item{
			Quantity: 1,
		},
	}

	var rule PromoCodeRule
	rule.Apply(&items)

	//Assertion
	if len(items) != len(expItems) {
		t.Errorf("Test FAILED. Actual: %v Expected: %v", len(items), len(expItems))
	}

	for k, v := range items {
		if v.Quantity != expItems[k].Quantity {
			t.Errorf("Test %v Quantity FAILED. Actual: %v Expected: %v", k, v.Quantity, expItems[k].Quantity)
		}
		if v.TotalPrice != expItems[k].TotalPrice {
			t.Errorf("Test %v TotalPrice FAILED. Actual: %v Expected: %v", k, v.TotalPrice, expItems[k].TotalPrice)
		}
	}
}

//Case2: Promo Code invalid
func TestPromoCodeRule2(t *testing.T) {
	//test data
	items := map[string]*shop.Item{
		"ult_medium": &shop.Item{
			Quantity:   1,
			TotalPrice: 29.90,
		},
		"ult_large": &shop.Item{
			Quantity:   4,
			TotalPrice: (44.90 * 4),
		},
		"i<3AMAYSIM": &shop.Item{
			Quantity: 1,
		},
	}
	//expected data
	expItems := map[string]shop.Item{
		"ult_medium": shop.Item{
			Quantity:   1,
			TotalPrice: 29.90,
		},
		"ult_large": shop.Item{
			Quantity:   4,
			TotalPrice: 44.90 * 4,
		},
		"i<3AMAYSIM": shop.Item{
			Quantity: 1,
		},
	}

	var rule PromoCodeRule
	rule.Apply(&items)

	//Assertion
	if len(items) != len(expItems) {
		t.Errorf("Test FAILED. Actual: %v Expected: %v", len(items), len(expItems))
	}

	for k, v := range items {
		if v.Quantity != expItems[k].Quantity {
			t.Errorf("Test %v Quantity FAILED. Actual: %v Expected: %v", k, v.Quantity, expItems[k].Quantity)
		}
		if v.TotalPrice != expItems[k].TotalPrice {
			t.Errorf("Test %v TotalPrice FAILED. Actual: %v Expected: %v", k, v.TotalPrice, expItems[k].TotalPrice)
		}
	}
}

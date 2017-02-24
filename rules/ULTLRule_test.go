package rules

import (
	"testing"

	"github.com/ball-oo-n/shopping-cart/shop"
)

//Case1: 1x ULTS, 4x ULTL
func TestULTLRule1(t *testing.T) {
	//test data
	items := map[string]*shop.Item{
		"ult_small": &shop.Item{
			Quantity: 1,
		},
		"ult_large": &shop.Item{
			Quantity: 4,
		},
	}

	//expected data
	expItems := map[string]shop.Item{
		"ult_small": shop.Item{
			Quantity: 1,
		},
		"ult_large": shop.Item{
			Quantity:   4,
			TotalPrice: 159.60,
		},
	}

	var rule ULTLRule
	rule.Apply(&items)

	//Assertion
	for k, v := range items {
		if v.Quantity != expItems[k].Quantity {
			t.Errorf("Test FAILED. Actual: %v Expected: %v", v.Quantity, expItems[k].Quantity)
		}
		if v.TotalPrice != expItems[k].TotalPrice {
			t.Errorf("Test FAILED. Actual: %v Expected: %v", v.TotalPrice, expItems[k].TotalPrice)
		}
	}
}

//Case2: 1x ULTL only
func TestULTLRule2(t *testing.T) {
	//test data
	items := map[string]*shop.Item{
		"ult_large": &shop.Item{
			Quantity: 1,
		},
	}

	//expected data
	expItems := map[string]shop.Item{
		"ult_large": shop.Item{
			Quantity:   1,
			TotalPrice: 44.90,
		},
	}

	var rule ULTLRule
	rule.Apply(&items)

	//Assertion
	for k, v := range items {
		if v.Quantity != expItems[k].Quantity {
			t.Errorf("Test FAILED. Actual: %v Expected: %v", v.Quantity, expItems[k].Quantity)
		}
		if v.TotalPrice != expItems[k].TotalPrice {
			t.Errorf("Test FAILED. Actual: %v Expected: %v", v.TotalPrice, expItems[k].TotalPrice)
		}
	}
}

package rules

import (
	"testing"

	"github.com/ball-oo-n/shopping-cart/shop"
)

//Case1: 1x ULTM, 4x ULTL
func TestULTMRule1(t *testing.T) {
	//test data
	items := map[string]*shop.Item{
		"ult_medium": &shop.Item{
			Quantity: 1,
		},
		"ult_large": &shop.Item{
			Quantity: 4,
		},
	}
	//expected data
	expItems := map[string]shop.Item{
		"ult_medium": shop.Item{
			Quantity:   1,
			TotalPrice: 29.90,
		},
		"ult_large": shop.Item{
			Quantity: 4,
		},
		"1gb": shop.Item{
			Quantity: 1,
		},
	}

	var rule ULTMRule
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

//Case2: 3x ULTM
func TestULTMRule2(t *testing.T) {
	//test data
	items := map[string]*shop.Item{
		"ult_medium": &shop.Item{
			Quantity: 3,
		},
	}
	//expected data
	expItems := map[string]shop.Item{
		"ult_medium": shop.Item{
			Quantity:   3,
			TotalPrice: float32(29.90 * 3),
		},
		"1gb": shop.Item{
			Quantity: 3,
		},
	}

	var rule ULTMRule
	rule.Apply(&items)

	//Assertion
	if len(items) != len(expItems) {
		t.Errorf("Test FAILED. Actual: %v Expected: %v", len(items), len(expItems))
	}

	for k, v := range items {
		if v.Quantity != expItems[k].Quantity {
			t.Errorf("Test FAILED. Actual: %v Expected: %v", v.Quantity, expItems[k].Quantity)
		}
		if v.TotalPrice != expItems[k].TotalPrice {
			t.Errorf("Test FAILED. Actual: %v Expected: %v", v.TotalPrice, expItems[k].TotalPrice)
		}
	}
}

//Case2: 3x ULTM
func TestULTMRule3(t *testing.T) {
	//test data
	items := map[string]*shop.Item{
		"ult_medium": &shop.Item{
			Quantity: 3,
		},
		"1gb": &shop.Item{
			Quantity: 1,
		},
	}
	//expected data
	expItems := map[string]shop.Item{
		"ult_medium": shop.Item{
			Quantity:   3,
			TotalPrice: float32(29.90 * 3),
		},
		"1gb": shop.Item{
			Quantity: 4,
		},
	}

	var rule ULTMRule
	rule.Apply(&items)

	//Assertion
	if len(items) != len(expItems) {
		t.Errorf("Test FAILED. Actual: %v Expected: %v", len(items), len(expItems))
	}

	for k, v := range items {
		if v.Quantity != expItems[k].Quantity {
			t.Errorf("Test FAILED. Actual: %v Expected: %v", v.Quantity, expItems[k].Quantity)
		}
		if v.TotalPrice != expItems[k].TotalPrice {
			t.Errorf("Test FAILED. Actual: %v Expected: %v", v.TotalPrice, expItems[k].TotalPrice)
		}
	}
}

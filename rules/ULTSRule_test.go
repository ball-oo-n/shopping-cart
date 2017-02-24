package rules

import (
	"testing"

	"github.com/ball-oo-n/shopping-cart/shop"
)

//Case1: 1x ULTS, 4x ULTL
func TestULTSRule1(t *testing.T) {
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
			Quantity:   1,
			TotalPrice: 24.90,
		},
		"ult_large": shop.Item{
			Quantity: 4,
		},
	}

	var rule ULTSRule
	rule.Apply(&items)

	//Assertion
	for k, v := range items {
		if v.Quantity != expItems[k].Quantity {
			t.Errorf("Test FAILED. Actual: %v Expected: %v", v.Quantity, expItems[k].Quantity)
		}
		if v.TotalPrice != expItems[k].TotalPrice {
			t.Errorf("Test %v FAILED. Actual: %v Expected: %v", k, v.TotalPrice, expItems[k].TotalPrice)
		}
	}
}

//Case2: 3x ULTS only
func TestULTSRule2(t *testing.T) {
	//test data
	items := map[string]*shop.Item{
		"ult_small": &shop.Item{
			Quantity: 3,
		},
	}
	//expected data
	expItems := map[string]shop.Item{
		"ult_small": shop.Item{
			Quantity:   3,
			TotalPrice: 49.80,
		},
	}

	var rule ULTSRule
	rule.Apply(&items)

	//Assertion
	for k, v := range items {
		if v.Quantity != expItems[k].Quantity {
			t.Errorf("Test FAILED. Actual: %v Expected: %v", v.Quantity, expItems[k].Quantity)
		}
		if v.TotalPrice != expItems[k].TotalPrice {
			t.Errorf("Test %v FAILED. Actual: %v Expected: %v", k, v.TotalPrice, expItems[k].TotalPrice)
		}
	}
}

//Case3: 0x ULTS, 4x ULTL
func TestULTSRule3(t *testing.T) {
	//test data
	items := map[string]*shop.Item{
		"ult_large": &shop.Item{
			Quantity: 4,
		},
	}
	//expected data
	expItems := map[string]shop.Item{
		"ult_large": shop.Item{
			Quantity: 4,
		},
	}

	var rule ULTSRule
	rule.Apply(&items)

	//Assertion
	for k, v := range items {
		if v.Quantity != expItems[k].Quantity {
			t.Errorf("Test FAILED. Actual: %v Expected: %v", v.Quantity, expItems[k].Quantity)
		}
		if v.TotalPrice != expItems[k].TotalPrice {
			t.Errorf("Test %v FAILED. Actual: %v Expected: %v", k, v.TotalPrice, expItems[k].TotalPrice)
		}
	}
}

//Case4: 4x ULTS, 8x ULTS
func TestULTSRule4(t *testing.T) {
	//test data
	items := map[string]*shop.Item{
		"ult_small": &shop.Item{
			Quantity: 12,
		},
	}
	//expected data
	expItems := map[string]shop.Item{
		"ult_small": shop.Item{
			Quantity:   12,
			TotalPrice: 8 * 24.90,
		},
	}

	var rule ULTSRule
	rule.Apply(&items)

	//Assertion
	for k, v := range items {
		if v.Quantity != expItems[k].Quantity {
			t.Errorf("Test FAILED. Actual: %v Expected: %v", v.Quantity, expItems[k].Quantity)
		}
		if v.TotalPrice != expItems[k].TotalPrice {
			t.Errorf("Test %v FAILED. Actual: %v Expected: %v", k, v.TotalPrice, expItems[k].TotalPrice)
		}
	}
}

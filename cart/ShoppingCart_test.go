package cart_test

import (
	"testing"

	. "github.com/ball-oo-n/shopping-cart/cart"
	r "github.com/ball-oo-n/shopping-cart/rules"
	"github.com/ball-oo-n/shopping-cart/shop"
)

var actItems []shop.Item
var expItemLen int
var expItems map[string]*shop.Item
var cart ShoppingCart

func initTestData() {
	actItems = []shop.Item{
		shop.Item{Quantity: 1, ItemCode: "ult_small"},
		shop.Item{Quantity: 1, ItemCode: "ult_medium"},
		shop.Item{Quantity: 1, ItemCode: "ult_large"},
		shop.Item{Quantity: 1, ItemCode: "1gb"},
	}

	expItemLen = 4
	expItems = map[string]*shop.Item{
		"ult_small":  &shop.Item{Quantity: 1, ItemCode: "ult_small"},
		"ult_medium": &shop.Item{Quantity: 1, ItemCode: "ult_medium"},
		"ult_large":  &shop.Item{Quantity: 1, ItemCode: "ult_large"},
		"1gb":        &shop.Item{Quantity: 1, ItemCode: "1gb"},
	}

	cart = ShoppingCart{Items: make(map[string]*shop.Item)}

	for _, i := range actItems {
		cart.Add(i)
	}
}

//Case1 Adding items
func TestShoppingCart1(t *testing.T) {
	initTestData()

	//Assertion
	if len(cart.Items) != expItemLen {
		t.Errorf("Test FAILED. Actual: %v Expected: %v", len(cart.Items), expItemLen)
	}

	for code := range expItems {
		if actItem := cart.Items[code]; actItem != nil {
			if actItem.ItemCode != expItems[code].ItemCode {
				t.Errorf("Test FAILED. Actual: %v Expected: %v", actItem.ItemCode, expItems[code].ItemCode)
			}
			if actItem.Quantity != expItems[code].Quantity {
				t.Errorf("Test FAILED. Actual: %v Expected: %v", actItem.Quantity, expItems[code].Quantity)
			}
			if actItem.TotalPrice != expItems[code].TotalPrice {
				t.Errorf("Test FAILED. Actual: %v Expected: %v", actItem.TotalPrice, expItems[code].TotalPrice)
			}
		} else {
			t.Error("Test FAILED. Actual: nil Expected: not nil")
		}
	}

	//add another test data
	for _, i := range actItems {
		cart.Add(i)
	}

	expItems = map[string]*shop.Item{
		"ult_small":  &shop.Item{Quantity: 2, ItemCode: "ult_small"},
		"ult_medium": &shop.Item{Quantity: 2, ItemCode: "ult_medium"},
		"ult_large":  &shop.Item{Quantity: 2, ItemCode: "ult_large"},
		"1gb":        &shop.Item{Quantity: 2, ItemCode: "1gb"},
	}

	//Assertion
	if len(cart.Items) != expItemLen {
		t.Errorf("Test FAILED. Actual: %v Expected: %v", len(cart.Items), expItemLen)
	}

	for code := range expItems {
		if actItem := cart.Items[code]; actItem != nil {
			if actItem.ItemCode != expItems[code].ItemCode {
				t.Errorf("Test FAILED. Actual: %v Expected: %v", actItem.ItemCode, expItems[code].ItemCode)
			}
			if actItem.Quantity != expItems[code].Quantity {
				t.Errorf("Test FAILED. Actual: %v Expected: %v", actItem.Quantity, expItems[code].Quantity)
			}
			if actItem.TotalPrice != expItems[code].TotalPrice {
				t.Errorf("Test FAILED. Actual: %v Expected: %v", actItem.TotalPrice, expItems[code].TotalPrice)
			}
		} else {
			t.Error("Test FAILED. Actual: nil Expected: not nil")
		}
	}
}

func TestShoppingCart2(t *testing.T) {
	expCode := "i<3amaysim"
	cart.AddPromo("i<3amaysim")
	if cart.PromoCode != expCode {
		t.Errorf("Test FAILED. Actual: %v Expected: %v", cart.PromoCode, expCode)
	}
}

func TestCalculate1(t *testing.T) {
	//test data
	p := []float32{24.90, 29.90, 44.90, 0}
	items := map[string]*shop.Item{
		"ult_small":  &shop.Item{Quantity: 1, ItemCode: "ult_small", TotalPrice: p[0]},
		"ult_medium": &shop.Item{Quantity: 1, ItemCode: "ult_medium", TotalPrice: p[1]},
		"ult_large":  &shop.Item{Quantity: 1, ItemCode: "ult_large", TotalPrice: p[2]},
		"1gb":        &shop.Item{Quantity: 1, ItemCode: "1gb", TotalPrice: p[3]},
	}

	var expTotal float32
	expTotal = 99.7
	var cart ShoppingCart
	cart.Items = items
	cart.PricingRules = r.PricingRules

	CalculateTotal(&cart)
	//Assertion
	if cart.Total != expTotal {
		t.Errorf("Test FAILED. Actual: %v Expected: %v", cart.Total, expTotal)
	}
}

//Case2: with promo code
func TestCalculate2(t *testing.T) {
	//test data
	p := []float32{24.90, 29.90, 44.90, 0}
	items := map[string]*shop.Item{
		"ult_small":  &shop.Item{Quantity: 1, ItemCode: "ult_small", TotalPrice: p[0]},
		"ult_medium": &shop.Item{Quantity: 1, ItemCode: "ult_medium", TotalPrice: p[1]},
		"ult_large":  &shop.Item{Quantity: 1, ItemCode: "ult_large", TotalPrice: p[2]},
		"1gb":        &shop.Item{Quantity: 1, ItemCode: "1gb", TotalPrice: p[3]},
	}

	var expTotal float32
	expTotal = 99.7 - (99.7 * 0.1)
	var cart ShoppingCart
	cart.Items = items
	cart.PromoCode = "I<3AMAYSIM"
	cart.PricingRules = r.PricingRules

	CalculateTotal(&cart)
	//Assertion
	if cart.Total != expTotal {
		t.Errorf("Test FAILED. Actual: %v Expected: %v", cart.Total, expTotal)
	}
}

//Case2: with invalid promo code
func TestCalculate3(t *testing.T) {
	//test data
	p := []float32{24.90, 29.90, 44.90, 0}
	items := map[string]*shop.Item{
		"ult_small":  &shop.Item{Quantity: 1, ItemCode: "ult_small", TotalPrice: p[0]},
		"ult_medium": &shop.Item{Quantity: 1, ItemCode: "ult_medium", TotalPrice: p[1]},
		"ult_large":  &shop.Item{Quantity: 1, ItemCode: "ult_large", TotalPrice: p[2]},
		"1gb":        &shop.Item{Quantity: 1, ItemCode: "1gb", TotalPrice: p[3]},
	}

	var expTotal float32
	expTotal = 99.7
	var cart ShoppingCart
	cart.Items = items
	cart.PromoCode = "iloveamaysim"

	CalculateTotal(&cart)
	//Assertion
	if cart.Total != expTotal {
		t.Errorf("Test FAILED. Actual: %v Expected: %v", cart.Total, expTotal)
	}
}

//Case2: with invalid promo code
func TestCalculate5(t *testing.T) {
	//test data
	p := []float32{24.90, 9.90}
	items := map[string]*shop.Item{
		"ult_small": &shop.Item{Quantity: 1, ItemCode: "ult_small", TotalPrice: p[0]},
		"1gb":       &shop.Item{Quantity: 1, ItemCode: "1gb", TotalPrice: p[1]},
	}

	var expTotal float32
	expTotal = 31.32
	var cart ShoppingCart
	cart.Items = items
	cart.PromoCode = "I<3AMAYSIM"

	CalculateTotal(&cart)
	//Assertion
	if cart.Total != expTotal {
		t.Errorf("Test FAILED. Actual: %v Expected: %v", cart.Total, expTotal)
	}
}

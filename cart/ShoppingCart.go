package cart

import (
	r "github.com/ball-oo-n/shopping-cart/rules"
	"github.com/ball-oo-n/shopping-cart/shop"
)

//ShoppingCart object
type ShoppingCart struct {
	Total        float32
	Items        map[string]*shop.Item
	PromoCode    string
	PricingRules []r.Rule
}

//Add function adds an item to the cart
func (s *ShoppingCart) Add(i shop.Item) {
	mapItem, exists := s.Items[i.ItemCode]
	if exists {
		mapItem.Quantity += i.Quantity
	} else {
		s.Items[i.ItemCode] = &i
	}
}

//AddPromo function adds the promo code to the shopping cart
func (s *ShoppingCart) AddPromo(promoCode string) {
	itmObj := shop.Item{
		Quantity:   1,
		TotalPrice: 0,
		ItemCode:   promoCode,
	}
	s.Add(itmObj)
}

//CalculateTotal sums up the total price of all items in the cart
func CalculateTotal(c *ShoppingCart) {
	items := c.Items

	for _, v := range c.PricingRules {
		v.Apply(&items)
	}

	var grandTotal float32
	for _, v := range items {
		grandTotal += v.TotalPrice
	}
	c.Total = grandTotal
}

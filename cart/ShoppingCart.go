package cart

import (
	p "github.com/ball-oo-n/shopping-cart/promotions"
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
	mapItem := s.Items[i.ItemCode]
	if mapItem != nil {
		mapItem.Quantity += i.Quantity
	} else {
		s.Items[i.ItemCode] = &i
	}
}

//AddPromo function adds the promo code to the shopping cart
func (s *ShoppingCart) AddPromo(promoCode string) {
	s.PromoCode = promoCode
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
	if len(c.PromoCode) > 0 {
		if d, ok := p.PromoCodes[c.PromoCode]; ok {
			d = grandTotal * (d)
			grandTotal = grandTotal - d
		}
	}
	c.Total = grandTotal
}

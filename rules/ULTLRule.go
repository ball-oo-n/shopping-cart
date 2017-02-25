package rules

import (
	c "github.com/ball-oo-n/shopping-cart/catalogue"
	"github.com/ball-oo-n/shopping-cart/shop"
)

//ULTLRule struct
type ULTLRule struct {
}

//Apply rule
func (r ULTLRule) Apply(items *map[string]*shop.Item) {
	code := "ult_large"
	item, isSatisfied := (*items)[code]

	if isSatisfied {
		if item.Quantity > 3 {
			item.TotalPrice = 39.90 * float32(item.Quantity)
		} else {
			item.TotalPrice = c.Catalogue[code].Price * float32(item.Quantity)
		}
	}
}

package rules

import (
	c "github.com/ball-oo-n/shopping-cart/catalogue"
	"github.com/ball-oo-n/shopping-cart/shop"
)

//ULTSRule struct
type ULTSRule struct {
}

//Apply rule
func (r ULTSRule) Apply(items *map[string]*shop.Item) {
	code := "ult_small"
	item, isSatisfied := (*items)[code]

	if isSatisfied {
		if item.Quantity >= 3 {
			price := float32(item.Quantity/3) * c.Catalogue[code].Price * 2
			price += float32(item.Quantity%3) * c.Catalogue[code].Price
			item.TotalPrice = price
		} else {
			price := float32(item.Quantity%3) * c.Catalogue[code].Price
			item.TotalPrice = price
		}
	}
}

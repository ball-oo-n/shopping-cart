package rules

import (
	c "github.com/ball-oo-n/shopping-cart/catalogue"
	"github.com/ball-oo-n/shopping-cart/shop"
)

//ULTMRule struct
type ULTMRule struct {
}

//Apply struct
func (r ULTMRule) Apply(items *map[string]*shop.Item) {
	code := "ult_medium"
	item, ok := (*items)[code]

	if ok {
		item.TotalPrice = c.Catalogue[code].Price * float32(item.Quantity)

		var qty int
		if j, exists := (*items)["1gb"]; exists {
			qty = j.Quantity
		}

		if item.Quantity > 0 {
			(*items)["1gb"] = &shop.Item{
				Quantity:   item.Quantity + qty,
				ItemCode:   "1gb",
				TotalPrice: float32(0),
			}
		}
	}
}

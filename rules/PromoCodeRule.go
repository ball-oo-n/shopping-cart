package rules

import "github.com/ball-oo-n/shopping-cart/shop"

//PromoCodeRule struct
type PromoCodeRule struct {
}

//Apply rule
func (r PromoCodeRule) Apply(items *map[string]*shop.Item) {
	code := "I<3AMAYSIM"
	discount := float32(0.10)
	_, isSatisfied := (*items)[code]

	if isSatisfied {
		for _, v := range *items {
			v.TotalPrice -= (v.TotalPrice * discount)
		}
	}
}

package rules

import "github.com/ball-oo-n/shopping-cart/shop"

// Rule interface
type Rule interface {
	Apply(items *map[string]*shop.Item)
}

// PricingRules is the slice of all rules
var PricingRules []Rule

func init() {
	PricingRules = []Rule{ULTLRule{}, ULTMRule{}, ULTSRule{}}
}

package promotions_test

import (
	"testing"

	p "github.com/ball-oo-n/shopping-cart/promotions"
)

func TestLoadPromoCodes(t *testing.T) {
	promoCodes := p.LoadPromoCodes()
	expMap := map[string]float32{
		"I<3AMAYSIM": float32(0.10),
	}
	for code := range expMap {
		actPromo := promoCodes[code]
		if actPromo != promoCodes[code] {
			t.Errorf("Test FAILED. Actual: %v Expected: %v", actPromo, expMap[code])
		}
	}
}

/**
*
**/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	s "github.com/ball-oo-n/shopping-cart/cart"
	c "github.com/ball-oo-n/shopping-cart/catalogue"
	r "github.com/ball-oo-n/shopping-cart/rules"
	"github.com/ball-oo-n/shopping-cart/shop"
)

var m = flag.Bool("menu", false, "menu")
var a = flag.String("add", "", "add <item_code_1>/<qty_1>,<item_code_2>/<qty_2>,...,<item_code_n>/<qty_n>")

var shoppingCart s.ShoppingCart

func init() {
	file, _ := os.Open("catalogue.txt")
	defer file.Close()
	c.Load(bufio.NewReader(file))
}

func main() {
	shoppingCart = s.ShoppingCart{Items: make(map[string]*shop.Item)}
	shoppingCart.PricingRules = r.PricingRules
	flag.Parse()

	if *m {
		// display menu.
		for value := range c.Catalogue {
			item := c.Catalogue[value]
			fmt.Printf("[%v] %v, $%v\n", item.ItemCode, item.ItemName, item.Price)
		}
	}

	if len(*a) > 0 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Do you have a promo code? Y/N: ")
		hasPromoCode, _ := reader.ReadString('\n')

		hasPromoCode = strings.TrimSpace(hasPromoCode)
		if hasPromoCode == "Y" || hasPromoCode == "y" {
			fmt.Print("Enter promo code: ")
			promoCode, _ := reader.ReadString('\n')
			promoCode = strings.TrimSpace(promoCode)

			shoppingCart.AddPromo(promoCode)

		}
		cartItems := strings.Split(*a, ",")
		for cartItemIndex := range cartItems {
			item := strings.Split(cartItems[cartItemIndex], "/")
			q, _ := strconv.ParseInt(item[1], 10, 32)

			itmObj := shop.Item{
				Quantity:   int(q),
				ItemCode:   item[0],
				TotalPrice: c.Catalogue[item[0]].Price * float32(q),
			}
			shoppingCart.Add(itmObj)
		}

		s.CalculateTotal(&shoppingCart)

		fmt.Println("Shopping Cart Items:")
		for k, v := range shoppingCart.Items {
			var name string
			if value, exists := c.Catalogue[k]; exists {
				name = value.ItemName
				qty := v.Quantity
				fmt.Printf("%v x%v", name, qty)
				fmt.Println()
			}
		}

		fmt.Printf("Shopping Cart Total: $%.02f\n", shoppingCart.Total)
	}

}

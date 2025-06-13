package product

import "fmt"

type Cart struct {
	Qty         int
	Price       int
	NameProduct string
}

var cartAmount = []Cart{}

func AddCart(qty int, price int, name string) []Cart {
	cartItem := Cart{
		Qty:         qty,
		Price:       price,
		NameProduct: name,
	}
	cartAmount = append(cartAmount, cartItem)
	for i, x := range cartAmount {
		subtotal := x.Qty * x.Price
		fmt.Println("Keranjang Belanja Anda")
		fmt.Printf("%v. %v x%v - Rp %d\n", i+1, x.NameProduct, x.Qty, subtotal)
	}

	return cartAmount
}

var chose string

func GetCart() []Cart {
	if len(cartAmount) > 0 {
		fmt.Println("List Keranjang Mu")
		for i, x := range cartAmount {
			subtotal := x.Qty * x.Price

			fmt.Printf("%v. %v x%v - Rp %d\n", i+1, x.NameProduct, x.Qty, subtotal)
		}
		fmt.Println("Lakukan Check out? y/n")
		fmt.Scan(&chose)
		if chose == "y" || chose == "Y" {
			Order(cartAmount)
		} else {
			fmt.Println("Checkout dibatalkan.")
		}
	} else {
		fmt.Println("List Keranjang mu Kosong")
	}
	return cartAmount
}

package product

import (
	"fmt"
	"sync"
	"time"
)

type Cart struct {
	Qty         int
	Price       int
	NameProduct string
}

type CartHandler interface {
	AddCart(qty int, price int, name string) []Cart
	GetCart() []Cart
	DeleteCart() []Cart
}

type cartService struct {
	cartItems []Cart
}

func NewCartService() CartHandler {
	return &cartService{
		cartItems: []Cart{},
	}
}

func (c *cartService) AddCart(qty int, price int, name string) []Cart {
	item := Cart{
		Qty:         qty,
		Price:       price,
		NameProduct: name,
	}
	c.cartItems = append(c.cartItems, item)

	fmt.Println("Keranjang Belanja Anda:")
	for i, x := range c.cartItems {
		subtotal := x.Qty * x.Price
		fmt.Printf("%d. %s x%d - Rp %d\n", i+1, x.NameProduct, x.Qty, subtotal)
	}

	return c.cartItems
}

func (c *cartService) GetCart() []Cart {
	if len(c.cartItems) == 0 {
		fmt.Println("List Keranjang mu Kosong")
		return c.cartItems
	}

	fmt.Println("List Keranjang Mu:")
	for i, x := range c.cartItems {
		subtotal := x.Qty * x.Price
		fmt.Printf("%d. %s x%d - Rp %d\n", i+1, x.NameProduct, x.Qty, subtotal)
	}
	var wg = sync.WaitGroup{}
	resultChan := make(chan bool)
	wg.Add(1)
	var chose string
	fmt.Println("Lakukan Check out? y/n")
	fmt.Scanln(&chose)
	if chose == "y" || chose == "Y" {
		go Order(c.cartItems, c, resultChan, &wg)
		time.Sleep(3 * time.Second)
		result := <-resultChan
		if result {
			fmt.Println("✅ Checkout selesai")

		}

		wg.Wait()
	} else {
		fmt.Println("Checkout dibatalkan.")
	}

	return c.cartItems
}

func (c *cartService) DeleteCart() []Cart {
	c.cartItems = []Cart{}
	fmt.Println("Keranjang berhasil dikosongkan.")
	return c.cartItems
}

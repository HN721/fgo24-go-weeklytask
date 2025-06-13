package product

import "fmt"

func Order(cart []Cart) int {
	fmt.Println("\n--- Detail Pesanan Anda ---")
	var total int
	for i, item := range cart {
		subtotal := item.Qty * item.Price
		total += subtotal
		fmt.Printf("%d. %s x%d - Rp %d\n", i+1, item.NameProduct, item.Qty, subtotal)
	}
	fmt.Printf("Total belanja: Rp %d\n", total)
	fmt.Println("Pesanan Anda berhasil diproses. Terima kasih telah berbelanja!")
	return total
}

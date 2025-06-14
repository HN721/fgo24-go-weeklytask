package product

import "fmt"

func Order(cart []Cart) {
	fmt.Println("\n--- Detail Pesanan Anda ---")
	var total int
	for i, item := range cart {
		subtotal := item.Qty * item.Price
		total += subtotal
		fmt.Printf("%d. %s x%d - Rp %d\n", i+1, item.NameProduct, item.Qty, subtotal)
		HistoryTransaction(item.NameProduct, item.Qty, total)
	}
	fmt.Printf("Total belanja: Rp %d\n", total)
	DeleteCart()
	fmt.Println("Pesanan Anda berhasil diproses. Terima kasih telah berbelanja!")
}

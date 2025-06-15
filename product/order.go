package product

import (
	"fmt"
	"sync"
	"time"
)

func Order(cart []Cart, handler CartHandler, resultChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() { resultChan <- true }()
	fmt.Println("\n--- Detail Pesanan Anda ---")
	var total int
	for i, item := range cart {
		subtotal := item.Qty * item.Price
		total += subtotal
		fmt.Printf("%d. %s x%d - Rp %d\n", i+1, item.NameProduct, item.Qty, subtotal)
		HistoryTransaction(item.NameProduct, item.Qty, total)
	}
	fmt.Printf("Total belanja: Rp %d\n", total)
	handler.DeleteCart()
	fmt.Println("Tunggu Sebentar Sedang melakukan Verifikasi Pesanan......")

	time.Sleep(3 * time.Second)

}

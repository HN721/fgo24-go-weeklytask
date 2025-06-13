package product

import (
	"fmt"
	"strings"
)

func Search(items []*List) {
	var keyword string
	fmt.Print("Masukkan nama produk yang ingin dicari: ")
	fmt.Scanln(&keyword)

	resultChan := make(chan *List)

	go searchByName(keyword, items, resultChan)
	// go searchByCategory(keyword, items, resultChan)

	result := <-resultChan // ini reciver
	if result != nil {
		fmt.Printf("✅ Ditemukan: %s - Rp %d (%s)\n", result.Name, result.Price, result.Category)
		var chose string
		fmt.Println("Tambahkan Keranjang? y/n")
		fmt.Scanln(&chose)
		if chose == "y" {
			AddCart(1, result.Price, result.Name)
			fmt.Println("Berhasil menambah kan data ke keranjang")

		}

	} else {
		fmt.Println("❌ Produk tidak ditemukan.")
	}
}

func searchByName(name string, items []*List, resultChan chan *List) {
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Name), name) {
			resultChan <- item // ini sender
			return
		}
	}
	resultChan <- nil
}

// func searchByCategory(name string, items []*List, resultChan chan *List) {
// 	for _, item := range items {
// 		if strings.Contains(strings.ToLower(item.Category), name) {
// 			resultChan <- item
// 			return
// 		}
// 	}
// 	resultChan <- nil
// }

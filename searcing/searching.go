package searcing

import (
	"fmt"
	"homework/product"
	"strings"
)

func Search(items []*product.List) {
	var keyword string
	fmt.Print("Masukkan nama produk yang ingin dicari: ")
	fmt.Scanln(&keyword)

	resultChan := make(chan *product.List)

	// go searchByName(keyword, items, resultChan)
	go searchByCategory(keyword, items, resultChan)

	result := <-resultChan
	if result != nil {
		fmt.Printf("✅ Ditemukan: %s - Rp %d (%s)\n", result.Name, result.Price, result.Category)
	} else {
		fmt.Println("❌ Produk tidak ditemukan.")
	}
}

func searchByName(name string, items []*product.List, resultChan chan *product.List) {
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Name), name) {
			resultChan <- item
			return
		}
	}
	resultChan <- nil
}
func searchByCategory(name string, items []*product.List, resultChan chan *product.List) {
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Category), name) {
			resultChan <- item
			return
		}
	}
	resultChan <- nil
}

package product

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ClearScreen() {
	print("\033[H\033[2J")
}
func Search(items []*List) {
	var keyword string
	fmt.Print("Masukkan nama produk yang ingin dicari: ")
	fmt.Scanln(&keyword)

	resultChan := make(chan []*List)
	go searchByName(keyword, items, resultChan)

	results := <-resultChan
	if len(results) == 0 {
		fmt.Println("❌ Produk tidak ditemukan.")
		return
	}

	const pageSize = 5
	totalPages := (len(results) + pageSize - 1) / pageSize
	currentPage := 0

	reader := bufio.NewReader(os.Stdin)

	for {
		start := currentPage * pageSize
		end := start + pageSize
		if end > len(results) {
			end = len(results)
		}
		ClearScreen()
		fmt.Printf("\n✅ Hasil Pencarian (Halaman %d/%d):\n", currentPage+1, totalPages)
		for i := start; i < end; i++ {
			fmt.Printf("%d. %s - Rp %d (%s)\n", i+1, results[i].Name, results[i].Price, results[i].Category)
		}

		if totalPages == 1 {
			break
		}

		fmt.Println("\n[n] Next | [p] Prev | [s <nomor>] Select Produk | [q] Quit")
		fmt.Print("Pilihan: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "n" && currentPage < totalPages-1 {
			currentPage++
		} else if input == "p" && currentPage > 0 {
			currentPage--
		} else if strings.HasPrefix(input, "s ") {
			var pilih int
			_, err := fmt.Sscanf(input, "s %d", &pilih)
			if err == nil && pilih > 0 && pilih <= len(results) {
				selected := results[pilih-1]
				AddCart(1, selected.Price, selected.Name)
				fmt.Printf("✅ '%s' ditambahkan ke keranjang.\n", selected.Name)
				return
			} else {
				fmt.Println("❌ Nomor produk tidak valid.")
			}
		} else if input == "q" {
			fmt.Println("❌ Batal memilih produk.")
			return
		} else {
			fmt.Println("❗ Perintah tidak dikenali.")
		}
	}
}

func searchByName(name string, items []*List, resultChan chan []*List) {
	var results []*List
	keyword := strings.ToLower(name)

	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Name), keyword) {
			results = append(results, item)
		}
	}

	resultChan <- results
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

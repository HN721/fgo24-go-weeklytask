package product

import (
	"bufio"
	"fmt"
	"homework/utils"
	"os"
	"strings"
	"sync"
	"time"
)

func ClearScreen() {
	print("\033[H\033[2J")
}
func Search(items []*List, cart CartHandler) {
	var keyword string
	fmt.Print("Masukkan nama produk yang ingin dicari: ")
	fmt.Scanln(&keyword)
	var wg sync.WaitGroup
	wg.Add(1)
	resultChan := make(chan []*List)
	go searchByName(keyword, items, resultChan, &wg)
	fmt.Println("🔍 Mencari produk...")
	time.Sleep(3 * time.Second)
	results := <-resultChan
	wg.Wait()

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
		end := min(start+pageSize, len(results))
		ClearScreen()
		fmt.Printf("\n✅ Hasil Pencarian (Halaman %d/%d):\n", currentPage+1, totalPages)
		for i := start; i < end; i++ {
			fmt.Printf("%d. %s - Rp %d (%s)\n", i+1, results[i].Name, results[i].Price, results[i].Category)
		}

		if totalPages == 1 {
			fmt.Println("[no] Produk ")
			fmt.Print("Pilihan: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			strings.HasPrefix(input, "no ")
			pilih := utils.Input
			_, err := fmt.Sscanf(input, "no %d", &pilih)
			if err == nil && pilih > 0 && pilih <= len(results) {
				selected := results[pilih-1]
				var qty int
				fmt.Printf("Berapa jumlah '%s' yang ingin dibeli? ", selected.Name)
				fmt.Scanln(&qty)
				cart.AddCart(qty, selected.Price, selected.Name)
				fmt.Printf("✅ '%s' ditambahkan ke keranjang.\n", selected.Name)
				return
			}
		}

		fmt.Println("\n[n] Next | [p] Prev | [no <nomor>]  Produk | [q] Quit")
		fmt.Print("Pilihan: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "n" && currentPage < totalPages-1 {
			currentPage++
		} else if input == "p" && currentPage > 0 {
			currentPage--
		} else if strings.HasPrefix(input, "no ") {
			var pilih int
			_, err := fmt.Sscanf(input, "no %d", &pilih)
			if err == nil && pilih > 0 && pilih <= len(results) {
				selected := results[pilih-1]
				cart.AddCart(1, selected.Price, selected.Name)
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

func searchByName(name string, items []*List, resultChan chan []*List, wg *sync.WaitGroup) {

	defer wg.Done()
	var results []*List
	keyword := strings.ToLower(name)
	if len(items) == 0 {
		resultChan <- results
		return
	}
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

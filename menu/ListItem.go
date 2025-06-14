package menu

import (
	"bufio"
	"fmt"
	"homework/product"
	"os"
	"strconv"
	"strings"
)

func ListMenuByCategory(items []*product.List, category string) []*product.List {

	pageSize := 4

	filtered := []*product.List{}

	for _, item := range items {
		if item.Category == category {
			filtered = append(filtered, item)
		}
	}

	if len(filtered) == 0 {
		fmt.Println("Tidak ada item pada kategori ini.")
		return nil
	}
	totalPages := (len(filtered) + pageSize - 1) / pageSize
	currentPage := 0
	reader := bufio.NewReader(os.Stdin)

	for {
		ClearScreen()
		start := currentPage * pageSize
		end := min(start+pageSize, len(filtered))
		fmt.Println(end)
		fmt.Printf("\n======= Daftar %s (Halaman %d dari %d) =======\n", category, currentPage+1, totalPages)
		fmt.Println(len(filtered))
		for i := start; i < end; i++ {
			fmt.Printf("%d. %s - Rp %d\n", i+1, filtered[i].Name, filtered[i].Price)
		}

		fmt.Println("\n[n] Next | [p] Prev | [no <nomor>] Select Produk | [q] Quit")
		fmt.Print("Pilih: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "n" && currentPage < totalPages-1 {
			currentPage++
		} else if input == "p" && currentPage > 0 {
			currentPage--
		} else if strings.HasPrefix(input, "no ") {
			numStr := strings.TrimPrefix(input, "no ")
			index, err := strconv.Atoi(numStr)
			if err != nil || index < 1 || index > len(filtered) {
				fmt.Println("❌ Nomor tidak valid.")
				continue
			}
			choseItem := filtered[index-1]
			var qty int
			fmt.Printf("Berapa jumlah '%s' yang ingin dibeli? ", choseItem.Name)
			fmt.Scanln(&qty)
			product.AddCart(qty, choseItem.Price, choseItem.Name)
			fmt.Println("✅ Ditambahkan ke keranjang.")
			break
		} else if input == "q" {
			break
		} else {
			fmt.Println("❌ Input tidak valid.")
		}
	}

	return filtered
}

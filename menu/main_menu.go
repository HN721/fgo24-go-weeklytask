package menu

import (
	"bufio"
	"fmt"
	"homework/product"
	"os"
	"strings"
)

var Choose int

func ClearScreen() {
	print("\033[H\033[2J")
}
func Menu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		ClearScreen()
		fmt.Println("======= Selamat Datang di HN CafeSHOP =======")
		fmt.Println("Pilih Menu:")
		fmt.Println("1. Lihat Menu Makanan")
		fmt.Println("2. Lihat Menu Minuman")
		fmt.Println("3. Lihat Menu Snack")
		fmt.Println("4. Cari Produk?")
		fmt.Println("5. Lihat Cart?")
		fmt.Println("6. EXIT")
		fmt.Print("Pilih : ")
		fmt.Scanln(&Choose)

		switch Choose {
		case 1:
			ClearScreen()
			ListMenuByCategory(product.Items, "Makanan")

		case 2:
			ClearScreen()

			ListMenuByCategory(product.Items, "Minuman")
		case 3:
			ClearScreen()

			ListMenuByCategory(product.Items, "Snack")
		case 4:
			ClearScreen()

			product.Search(product.Items)
		case 5:
			ClearScreen()
			product.GetCart()

		case 6:
			fmt.Println("Terima kasih telah berkunjung!")
			os.Exit(0)

		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
		fmt.Print("\nIngin pesan lagi? (y/n): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer != "y" {
			fmt.Println("Terima kasih telah berkunjung!")
			return
		}
		fmt.Println()

	}
}

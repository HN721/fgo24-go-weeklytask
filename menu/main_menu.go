package menu

import (
	"bufio"
	"fmt"
	"homework/product"
	"homework/utils"
	"os"
	"strings"
)

func ClearScreen() {
	print("\033[H\033[2J")
}
func Menu(cart product.CartHandler) {
	reader := bufio.NewReader(os.Stdin)

	for {
		ClearScreen()
		menuUser()
		input := utils.Input
		fmt.Scanln(&input)

		switch input {
		case 1:
			ClearScreen()
			ListMenuByCategory(product.Items, "Makanan", cart)

		case 2:
			ClearScreen()

			ListMenuByCategory(product.Items, "Minuman", cart)
		case 3:
			ClearScreen()

			ListMenuByCategory(product.Items, "Snack", cart)
		case 4:
			ClearScreen()

			product.Search(product.Items, cart)
		case 5:
			ClearScreen()
			cart.GetCart()
		case 6:
			product.ShowHistory()
		case 7:
			fmt.Println("Terima kasih telah berkunjung!")
			os.Exit(0)

		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
		fmt.Print("\n Kembali ke menu? (y/n): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer != "y" {
			fmt.Println("Terima kasih telah berkunjung!")
			return
		}
		fmt.Println()

	}
}

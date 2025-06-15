package menu

import (
	"fmt"
	"homework/product"
	"homework/user"
	"homework/utils"
)

func Menu_user() {
	for {
		ClearScreen()
		layoutMenu()
		Input := utils.Input
		fmt.Scanln(&Input)
		switch Input {
		case 1:
			user.Register()
		case 2:
			if user.Login() {
				cart := product.NewCartService()

				Menu(cart)
				return
			} else {
				fmt.Println("❌ Email atau Password salah")
			}
		case 3:
			fmt.Println("Keluar dari aplikasi...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

package menu

import (
	"fmt"
	"homework/user"
)

func Menu_user() {
	for {
		ClearScreen()
		fmt.Println("======= Login/Register di HN CafeSHOP =======")
		fmt.Println("Pilih Menu:")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. EXIT")
		fmt.Print("Pilih : ")
		fmt.Scanln(&Choose)

		switch Choose {
		case 1:
			user.Register()
		case 2:
			if user.Login() {
				Menu()
				return // ⬅️ TAMBAHKAN INI!
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

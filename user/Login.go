package user

import (
	"fmt"
)

func Login() bool {
	var email, password string

	fmt.Println("Masukan Email :")
	fmt.Scanln(&email)
	fmt.Println("Masukan password :")
	fmt.Scanln(&password)
	for _, user := range Data {
		if user.email == email && user.password == password {
			fmt.Println("✅ Login Berhasil")
			return true
		}
	}

	fmt.Println("❌ Email atau password salah")
	return false
}

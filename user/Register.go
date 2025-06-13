package user

import "fmt"

func Register() {
	var username, email, password string

	fmt.Println("Masukan Username :")
	fmt.Scanln(&username)
	fmt.Println("Masukan Email :")
	fmt.Scanln(&email)
	fmt.Println("Masukan password :")
	fmt.Scanln(&password)
	AddRegister(username, email, password)
}

func AddRegister(username string, email string, password string) User {
	userData := User{
		email:    email,
		username: username,
		password: password,
	}
	Data = append(Data, userData)
	fmt.Println("Berhasil Register", Data)
	return userData
}

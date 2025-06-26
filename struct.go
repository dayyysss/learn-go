package main

import "fmt"

type User struct {
	Nama string
	Umur int
}

func main() {
	user1 := User{
		Nama: "Ferdi",
		Umur: 18,
	}

	fmt.Println(user1.Nama, "berumur", user1.Umur)
}
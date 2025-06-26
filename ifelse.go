package main

import "fmt"

func main() {
	umur := 18

	if umur >= 19 {
		fmt.Println("Lu udah dewasa bray.")
	} else {
		fmt.Println("Masih bocil nikmati dulu masa kecillu.")
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("Angka ke-", i)
	}
}
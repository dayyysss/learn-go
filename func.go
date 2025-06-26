package main

import "fmt"

func sapa(nama string) {
	fmt.Println("Halo", nama)
}

func tambah(a int, b int) int{
	return a + b
}

func main() {
	sapa("Ferdi")
	hasil := tambah(5, 3)
	fmt.Println("Hasil penjumlahan:", hasil)
}
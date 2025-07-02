package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var todos []string
var fileName = "todos.txt"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	loadTugas()

	for {
		fmt.Println("\n=== TODO LIST ===")
		fmt.Println("1. Lihat semua tugas")
		fmt.Println("2. Tambah tugas")
		fmt.Println("3. Hapus tugas")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih Menu: ")

		scanner.Scan()
		pilihan := scanner.Text()

		switch pilihan {
		case "1":
			tampilkanTugas()
		case "2":
			tambahTugas(scanner)
		case "3":
			hapusTugas(scanner)
		case "0":
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Menu ga valid!")
		}
	}
}

func tampilkanTugas() {
	if len(todos) == 0 {
		fmt.Println("Belum ada tugas.")
		return
	}
	for i, t := range todos {
		fmt.Printf("%d. %s\n", i+1, t)
	}
}

func tambahTugas(scanner *bufio.Scanner) {
	fmt.Print("Masukkan tugas: ")
	scanner.Scan()
	tugas := scanner.Text()
	todos = append(todos, tugas)
	saveTugas()
	fmt.Println("Tugas ditambahkan!")
}

func hapusTugas(scanner *bufio.Scanner) {
	tampilkanTugas()
	if len(todos) == 0 {
		return
	}
	fmt.Print("Nomor tugas yang mau dihapus: ")
	scanner.Scan()
	var index int
	fmt.Sscan(scanner.Text(), &index)

	if index < 1 || index > len(todos) {
		fmt.Println("Nomor ga valid.")
		return
	}
	todos = append(todos[:index-1], todos[index:]...)
	saveTugas()
	fmt.Println("Tugas dihapus.")
}

func loadTugas() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	for _, line := range lines { 
		if line != "" {
			todos = append(todos, line)
		}
	}
}

func saveTugas() { 
	data := strings.Join(todos, "\n")
	err := os.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		fmt.Println("Gagal Simpan:", err)
	}
}
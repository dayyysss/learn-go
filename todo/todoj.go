package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
)

type Todo struct {
    Judul   string
    Selesai bool
}


var todos []Todo
var fileName = "todos.txt"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	loadTugas()

	for {
		fmt.Println("\n=== TODO LIST ===")
		fmt.Println("1. Lihat semua tugas")
		fmt.Println("2. Tambah tugas")
		fmt.Println("3. Hapus tugas")
		fmt.Println("4. Tandai tugas selesai")
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
		case "4":
			markTugasSelesai(scanner)
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

	fmt.Println("Daftar Tugas:")
	for i, t := range todos {
		status := "Belum"
		if t.Selesai {
			status = "Selesai"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, t.Judul, status)
	}
}

func tambahTugas(scanner *bufio.Scanner) {
	fmt.Print("Masukkan tugas: ")
	scanner.Scan()
	judul := scanner.Text()

	newTodo := Todo{
		Judul: judul,
		Selesai: false,
	}

	todos = append(todos, newTodo)
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
        fmt.Println("Nomor nggak valid.")
        return
    }

    todos = append(todos[:index-1], todos[index:]...)
    saveTugas()
    fmt.Println("Tugas dihapus.")
}

func markTugasSelesai(scanner *bufio.Scanner) {
	tampilkanTugas()
	if len(todos) == 0 {
		return
	}

	fmt.Print("Nomor tugas yang mau ditandai selesai: ")
	scanner.Scan()
	var index int
	fmt.Sscan(scanner.Text(), &index)

	if index < 1 || index > len(todos) {
		fmt.Println("Nomor nggak valid.")
		return
	}

	if todos[index-1].Selesai {
		fmt.Println("Tugas sudah selesai sebelumnya!")
		return
	}

	todos[index-1].Selesai = true
	saveTugas()
	fmt.Println("Tugas ditandai selesai!")
}


func loadTugas() {
	file, err := os.Open("todos.json")
	if err != nil { 
		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&todos)
	if err != nil {
		fmt.Println("Gagal decode JSON:", err)
	}
}

func saveTugas() { 
	file, err := os.Create("todos.json")
	if err != nil {
		fmt.Println("Gagal buat file:", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(todos)
	if err != nil { 
		fmt.Println("Gagal encode JSON:", err)
	}
}
package main

import (
	"fmt"
	"os"
)

type Teman struct {
	ID    int
	Nama     string
	Alamat   string
	Pekerjaan string
	Alasan   string
}

func main() {
	temanList := []Teman{
		{1, "Thomas", "Jakarta", "Developer", "Ingin mempelajari Golang"},
		{2, "Dio", "Balikpapan", "Developer", "Ingin mempelajari Python"},
	}

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Gunakan: go run biodata.go [nama]")
		return
	}

	namaTeman := args[0]

	tampilkanDataTeman(namaTeman, temanList)
}

func tampilkanDataTeman(nama string, temanList []Teman) {
	for _, teman := range temanList {
		if teman.Nama == nama {
			fmt.Printf("Data Teman:\nAbsen: %d\nNama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s\n", teman.ID, teman.Nama, teman.Alamat, teman.Pekerjaan, teman.Alasan)
			return
		}
	}
	fmt.Printf("Tidak ada data teman dengan nama %s\n", nama)
}
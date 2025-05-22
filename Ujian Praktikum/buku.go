package main

import "fmt"

type buku struct {
	kode    int
	judul   string
	penulis string
	tahun   int
}

type arrBuku [100]buku

func isiData(T *arrBuku, n *int) {
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Println("Buku ke-", i+1)
		fmt.Print("Kode: ")
		fmt.Scan(&T[i].kode)

		fmt.Print("Judul (pisahkan kata dengan _): ")
		fmt.Scan(&T[i].judul)

		fmt.Print("Penulis (pisahkan kata dengan _): ")
		fmt.Scan(&T[i].penulis)

		fmt.Print("Tahun terbit: ")
		fmt.Scan(&T[i].tahun)
	}
}

func urutkanByTahun(T *arrBuku, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if T[i].tahun < T[j].tahun {
				T[i], T[j] = T[j], T[i]
			}
		}
	}
}

func urutkanByKode(T *arrBuku, n int) {

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if T[i].kode > T[j].kode {
				T[i], T[j] = T[j], T[i]
			}
		}
	}
}

func cariBuku(T arrBuku, n, kodeCari int) {
	ditemukan := false
	for i := 0; i < n; i++ {
		if T[i].kode == kodeCari {
			fmt.Println("\nBuku ditemukan:")
			fmt.Printf("Kode: %d\nJudul: %s\nPenulis: %s\nTahun: %d\n",
				T[i].kode, T[i].judul, T[i].penulis, T[i].tahun)
			ditemukan = true
			break
		}
	}
	if !ditemukan {
		fmt.Println("Buku tidak ditemukan.")
	}
}

func tampilkanBuku(T arrBuku, n int) {
	fmt.Println("Daftar Buku:")
	for i := 0; i < n; i++ {
		fmt.Printf("Kode: %d | %s - %s (%d)\n", T[i].kode, T[i].judul, T[i].penulis, T[i].tahun)

	}
}

func main() {
	var T arrBuku
	var n, kodeCari int

	isiData(&T, &n)

	fmt.Println("\nMengurutkan berdasarkan tahun terbit (desc)...")
	urutkanByTahun(&T, n)
	tampilkanBuku(T, n)

	fmt.Println("\nMengurutkan kembali berdasarkan kode buku (asc)...")
	urutkanByKode(&T, n)
	tampilkanBuku(T, n)

	fmt.Print("\nMasukkan kode buku yang ingin dicari: ")
	fmt.Scan(&kodeCari)

	cariBuku(T, n, kodeCari)
}

// Menruut saya sudah sangat bagus perkuliahan ini, hanya saja beberapa momen sepertinya terlalu terburu-buru dalam mengajar materi sehingga kami kurang memahami dengan baik, seperti teorinya yang terlalu cepat dan juga prakteknya yang kurang. Sekian terimakasih.

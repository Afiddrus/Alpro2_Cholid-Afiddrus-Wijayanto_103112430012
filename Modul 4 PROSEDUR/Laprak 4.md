
# <h1 align="center">Laporan Praktikum Modul 4 <br> Prosedur</h1>
<p align="center">Cholid Afiddrus Wijayanto - 103112430012</p>

## Dasar Teori

Prosedur adalah bagian dari program yang berisi sekumpulan instruksi untuk melakukan tugas tertentu tanpa mengembalikan nilai. Prosedur membantu menyederhanakan kode dengan memecah program menjadi bagian-bagian yang lebih kecil dan mudah dipahami. Dalam penggunaannya, prosedur dapat dipanggil di dalam program utama untuk menjalankan fungsinya. Dengan menggunakan prosedur, kode menjadi lebih modular, terorganisir, dan dapat digunakan kembali tanpa perlu menulis ulang instruksi yang sama berkali-kali​

## Unguided

### Soal 1

Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) Masukan terdiri dari empat buah bilangan asli 𝑎, 𝑏, 𝑐, dan 𝑑 yang dipisahkan oleh spasi, dengan syarat 𝑎 ≥ 𝑐 dan 𝑏 ≥ 𝑑. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi 𝒂 terhadap 𝑐, sedangkan baris kedua adalah hasil permutasi dan kombinasi 𝑏 terhadap 𝑑. Catatan: permutasi (P) dan kombinasi (C) dari 𝑛 terhadap 𝑟 (𝑛 ≥ 𝑟) dapat dihitung dengan menggunakan persamaan berikut!
$$𝑃(𝑛, 𝑟) = 𝑛! / (𝑛−𝑟)! ,  sedangkan 𝐶(𝑛, 𝑟) = 𝑛! / 𝑟!(𝑛−𝑟)!$$

```go
package main

  

import "fmt"

  

func faktorial(n int, hasil *int) {

    *hasil = 1

    for i := 2; i <= n; i++ {

        *hasil *= i

    }

}

  

func permutasi(n, r int, hasil *int) {

    var faktN, faktNR int

    faktorial(n, &faktN)

    faktorial(n-r, &faktNR)

    *hasil = faktN / faktNR

}

  

func kombinasi(n, r int, hasil *int) {

    var faktN, faktR, faktNR int

    faktorial(n, &faktN)

    faktorial(r, &faktR)

    faktorial(n-r, &faktNR)

    *hasil = faktN / (faktR * faktNR)

}

  

func main() {

    var angka1, angka2, angka3, angka4 int

    fmt.Scan(&angka1, &angka2, &angka3, &angka4)

  

    var perm1, komb1, perm2, komb2 int

    permutasi(angka1, angka3, &perm1)

    kombinasi(angka1, angka3, &komb1)

    permutasi(angka2, angka4, &perm2)

    kombinasi(angka2, angka4, &komb2)

  

    fmt.Println(perm1, komb1)

    fmt.Println(perm2, komb2)

}
```

> Output
![[Modul 4/output/soal1.jpg]]


Program ini menghitung permutasi dan kombinasi berdasarkan input dari user. Pertama, user diminta memasukkan empat angka, yaitu `angka1`, `angka2`, `angka3`, dan `angka4`, yang masing-masing berfungsi sebagai nilai `n` dan `r` dalam perhitungan permutasi dan kombinasi. Program menggunakan prosedur `faktorial(n int, hasil *int)` untuk menghitung faktorial dengan cara mengalikan angka dari 2 hingga `n`, lalu menyimpannya dalam variabel `hasil`. Prosedur `permutasi(n, r int, hasil *int)` menghitung permutasi dengan membagi faktorial `n` dengan faktorial `(n-r)`, sedangkan prosedur `kombinasi(n, r int, hasil *int)` menghitung kombinasi dengan membagi hasil permutasi dengan faktorial `r`, sesuai dengan rumus `C(n, r) = P(n, r) / r!`. Setelah membaca input, `main()` memanggil prosedur `permutasi` dan `kombinasi` dua kali, pertama untuk pasangan `angka1` dan `angka3`, lalu untuk `angka2` dan `angka4`. Hasilnya disimpan dalam variabel `perm1`, `komb1`, `perm2`, dan `komb2`, lalu ditampilkan dalam dua baris output, masing-masing berisi hasil permutasi dan kombinasi dari pasangan angka yang sesuai.
### Soal 2

Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. prosedure hitungSkor(in/out soal, skor : integer) Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.

Keterangan: Astuti menyelesaikan 6 soal dalam waktu 287 menit, sedangkan Bertha 7 soal dalam waktu 294 menit. Karena Bertha menyelesaikan lebih banyak, maka Bertha menang. Jika keduanya menyelesaikan sama banyak, maka pemenang adalah yang menyelesaikan dengan total waktu paling kecil.

```go
package main

  

import "fmt"

  

func hitungSkor(soal *int, skor *int) {

    *soal = 0

    *skor = 0

    batasWaktu := 301

  

    for i := 0; i < 8; i++ {

        var waktu int

        fmt.Scan(&waktu)

        if waktu < batasWaktu {

            *soal++

            *skor += waktu

        }

    }

}

  

func main() {

    var pemenang string

    var maxSoal, minSkor int = 0, 1<<31 - 1

  

    for {

        var nama string

        fmt.Scan(&nama)

        if nama == "Selesai" {

            break

        }

  

        var soal, skor int

        hitungSkor(&soal, &skor)

  

        if soal > maxSoal || (soal == maxSoal && skor < minSkor) {

            maxSoal = soal

            minSkor = skor

            pemenang = nama

        }

    }

  

    fmt.Println(pemenang, maxSoal, minSkor)

}
```

> Output
![[Modul 4/output/soal2.jpg]]

Program ini digunakan untuk menentukan pemenang dalam kompetisi pemrograman berdasarkan jumlah soal yang berhasil diselesaikan dan total waktu pengerjaan. Pertama, ada prosedur `hitungSkor(soal *int, skor *int)` yang menghitung jumlah soal yang dikerjakan dan total waktu yang dibutuhkan. Program membaca delapan angka satu per satu, di mana setiap angka mewakili waktu pengerjaan soal. Jika waktu pengerjaan lebih dari 301 menit, soal dianggap tidak dikerjakan. Dalam `main()`, program meminta user memasukkan nama peserta dan waktu pengerjaan setiap soal. Jika nama yang dimasukkan adalah "Selesai", program berhenti menerima input. Setiap peserta dibandingkan berdasarkan jumlah soal yang diselesaikan, dan jika jumlahnya sama, maka peserta dengan total waktu terkecil akan dipilih sebagai pemenang. Program menyimpan informasi pemenang dalam variabel `pemenang`, `maxSoal`, dan `minSkor`, lalu mencetak output berupa nama pemenang, jumlah soal yang diselesaikan, dan total waktu pengerjaan.

### Soal 3

Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah ½n, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1. Sebagai contoh jika dimulai dengan n=22, maka deret bilangan yang diperoleh adalah: 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1 Untuk suku awal sampai dengan 1000000, diketahui deret selalu mencapai suku dengan nilai 1. Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal. prosedure cetakDeret(in n : integer ) Masukan berupa satu bilangan integer positif yang lebih kecil dari 1000000. Keluaran terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang dan dipisahkan oleh sebuah spasi.

```go
package main

  

import (

    "fmt"

)

  

func cetakDeret(n int) {

    for n != 1 {

        fmt.Print(n, " ")

        if n%2 == 0 {

            n /= 2

        } else {

            n = 3*n + 1

        }

    }

    fmt.Println(n)

}

  

func main() {

    var angka int

    fmt.Scan(&angka)

    cetakDeret(angka)

}
```

> Output
![[Modul 4/output/soal3.jpg]]

Program ini mencetak deret bilangan berdasarkan input dari user. Pertama, user diminta memasukkan sebuah angka yang kemudian diproses dalam prosedur `cetakDeret(n int)`. Di dalam prosedur ini, angka akan dicetak terlebih dahulu, lalu diperiksa apakah angka tersebut genap atau ganjil. Jika genap, angka akan dibagi 2, sedangkan jika ganjil, angka akan dikalikan 3 lalu ditambah 1. Proses ini terus berulang hingga angka mencapai nilai 1, yang kemudian dicetak sebagai angka terakhir dalam deret. Setelah menerima input angka, fungsi `main()` memanggil prosedur `cetakDeret()` untuk menampilkan hasilnya dalam satu baris output.
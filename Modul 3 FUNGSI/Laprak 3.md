
# <h1 align="center">Laporan Praktikum Modul 3 <br> Fungsi</h1>
<p align="center">Cholid Afiddrus Wijayanto - 103112430012</p>

## Dasar Teori

Fungsi adalah bagian kode yang bisa dipanggil untuk menjalankan tugas tertentu. Dengan menggunakan fungsi, Maka tidak perlu lagi menulis ulang kode yang sama berulang kali, sehingga program jadi lebih rapi dan mudah diperbaiki. Fungsi biasanya punya nama, bisa menerima input (parameter), dan bisa mengembalikan hasil.

## Unguided

### Soal 1

Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) Masukan terdiri dari empat buah bilangan asli 𝑎, 𝑏, 𝑐, dan 𝑑 yang dipisahkan oleh spasi, dengan syarat 𝑎 ≥ 𝑐 dan 𝑏 ≥ 𝑑. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi 𝒂 terhadap 𝑐, sedangkan baris kedua adalah hasil permutasi dan kombinasi 𝑏 terhadap 𝑑. Catatan: permutasi (P) dan kombinasi (C) dari 𝑛 terhadap 𝑟 (𝑛 ≥ 𝑟) dapat dihitung dengan menggunakan persamaan berikut!

```go
package main

  

import "fmt"

  

func faktorial(n int) int {

    if n == 0 || n == 1 {

        return 1

    }

    hasil := 1

    for i := 2; i <= n; i++ {

        hasil *= i

    }

    return hasil

}

  

func permutasi(n, r int) int {

    return faktorial(n) / faktorial(n-r)

}

  

func kombinasi(n, r int) int {

    return faktorial(n) / (faktorial(r) * faktorial(n-r))

}

  

func main() {

    var a, b, c, d int

    fmt.Scan(&a, &b, &c, &d)

  

    if a >= c && b >= d {

        fmt.Println(permutasi(a, c), kombinasi(a, c))

        fmt.Println(permutasi(b, d), kombinasi(b, d))

    }

}
```

> Output
> ![[Modul%203%20FUNGSI/output/soal1.jpg]]


Program ini digunakan untuk menghitung permutasi dan kombinasi berdasarkan input dari user. Pertama, ada fungsi faktorial(n int) int yang menghitung faktorial dari n. Jika n adalah 0 atau 1, maka langsung mengembalikan 1. Jika lebih dari 1, program akan menghitung faktorial dengan mengalikan angka dari 2 hingga n.

Fungsi permutasi(n, r int) int menghitung permutasi dengan membagi faktorial n dengan faktorial (n-r). Sedangkan kombinasi(n, r int) int hampir serupa, tetapi hasilnya juga dibagi dengan faktorial r. Di dalam main(), program meminta user memasukkan empat angka (a, b, c, d). Jika memenuhi syarat, program akan menghitung dua nilai permutasi dan dua kombinasi, lalu menampilkan hasilnya di layar.

### Soal 2

Diberikan tiga buah fungsi matematika yaitu 𝑓 (𝑥) = 𝑥 2 , 𝑔 (𝑥) = 𝑥 − 2 dan ℎ (𝑥) = 𝑥 + 1. Fungsi komposisi (𝑓𝑜𝑔𝑜ℎ)(𝑥) artinya adalah 𝑓(𝑔(ℎ(𝑥))). Tuliskan 𝑓(𝑥), 𝑔(𝑥) dan ℎ(𝑥) dalam bentuk function. Masukan terdiri dari sebuah bilangan bulat 𝑎, 𝑏 dan 𝑐 yang dipisahkan oleh spasi. Keluaran terdiri dari tiga baris. Baris pertama adalah (𝑓𝑜𝑔𝑜ℎ)(𝑎), baris kedua (𝑔𝑜ℎ𝑜𝑓)(𝑏), dan baris ketiga adalah (ℎ𝑜𝑓𝑜𝑔)(𝑐)!

```go
package main

  

import "fmt"

  

func f(x int) int {

    return x * x

}

  

func g(x int) int {

    return x - 2

}

  

func h(x int) int {

    return x + 1

}

  

func fogoh(x int) int {

    return f(g(h(x)))

}

  

func gohof(x int) int {

    return g(h(f(x)))

}

  

func hofog(x int) int {

    return h(f(g(x)))

}

  

func main() {

    var a, b, c int

    fmt.Scan(&a, &b, &c)

  

    fmt.Println(fogoh(a))

    fmt.Println(gohof(b))

    fmt.Println(hofog(c))

}
```

> Output
![[Modul 3 FUNGSI/output/soal2.jpg]]

Program ini menghitung hasil dari tiga fungsi matematika yang dikombinasikan, yaitu f(x) = x², g(x) = x - 2, dan h(x) = x + 1. Fungsi-fungsi ini digabung dalam bentuk komposisi, seperti fogoh(x) = f(g(h(x))), yang berarti h(x) dijalankan dulu, hasilnya masuk ke g(x), lalu ke f(x). Begitu juga dengan gohof(x) = g(h(f(x))) dan hofog(x) = h(f(g(x))), yang masing-masing memiliki urutan perhitungan berbeda.

Di dalam fungsi main(), program meminta user memasukkan tiga angka (a, b, c). Setelah itu, program menghitung hasil dari tiga fungsi komposisi tersebut dan menampilkannya di layar. Dengan cara ini, program menunjukkan bagaimana angka yang dimasukkan berubah saat diproses oleh beberapa fungsi secara berurutan.

### Soal 3

Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥, 𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥, 𝑦) berdasarkan dua lingkaran tersebut. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

```go
package main

  

import (

    "fmt"

    "math"

)

  

func jarak(a, b, c, d int) float64 {

    return math.Sqrt(float64((a-c)*(a-c) + (b-d)*(b-d)))

}

  

func didalam(cx, cy, r, x, y int) bool {

    return jarak(cx, cy, x, y) <= float64(r)

}

  

func main() {

    var cx1, cy1, r1 int

    var cx2, cy2, r2 int

    var x, y int

  

    fmt.Scan(&cx1, &cy1, &r1)

    fmt.Scan(&cx2, &cy2, &r2)

    fmt.Scan(&x, &y)

  

    dalam1 := didalam(cx1, cy1, r1, x, y)

    dalam2 := didalam(cx2, cy2, r2, x, y)

  

    if dalam1 && dalam2 {

        fmt.Println("Titik di dalam lingkaran 1 dan 2")

    } else if dalam1 {

        fmt.Println("Titik di dalam lingkaran 1")

    } else if dalam2 {

        fmt.Println("Titik di dalam lingkaran 2")

    } else {

        fmt.Println("Titik tidak di dalam lingkaran 1 atau 2")

    }

}
```

> Output
![[Modul 3 FUNGSI/output/soal3.jpg]]

![[soal3(2).jpg]]

Program ini digunakan untuk menentukan apakah sebuah titik berada di dalam satu atau dua lingkaran berdasarkan koordinat pusat dan jari-jari lingkaran. Pertama, ada fungsi jarak(a, b, c, d int) float64 yang menghitung jarak antara dua titik menggunakan rumus Pythagoras. Kemudian, fungsi didalam(cx, cy, r, x, y int) bool memeriksa apakah titik (x, y) berada di dalam lingkaran dengan cara membandingkan jaraknya dengan jari-jari lingkaran.

Di dalam fungsi main(), program meminta user untuk memasukkan koordinat pusat dan jari-jari dari dua lingkaran, lalu titik yang akan diperiksa. Kemudian program mengecek apakah titik tersebut berada dalam lingkaran 1, lingkaran 2, keduanya, atau di luar keduanya. Hasil akhirnya akan ditampilkan dalam bentuk teks, seperti "Titik di dalam lingkaran 1 dan 2" atau "Titik tidak di dalam lingkaran 1 atau 2".

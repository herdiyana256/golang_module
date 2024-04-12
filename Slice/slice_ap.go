// Slice (Append Function) =digunakan untuk menambahkan satu atau beberapa elemen ke dalam slice. Jika kapasitas array yang mendasari tidak mencukupi, append() akan membuat array baru yang lebih besar, menyalin elemen-elemen lama, dan menambahkan elemen baru. Ini memungkinkan slice untuk "tumbuh" sesuai kebutuhan.



package main // Mendefinisikan package utama

import "fmt" // Mengimpor package fmt untuk formatting dan output

func main() { // Mendefinisikan fungsi main sebagai titik masuk program

	var fruits = make([]string, 4) // Membuat slice fruits dengan panjang awal 4 menggunakan fungsi make
	_ = fruits // Mengabaikan nilai kembalian dari make untuk menghindari error saat kompilasi

	fruits[0] = "srikaya"    // Menetapkan nilai "srikaya" ke indeks 0 slice fruits
	fruits[1] = "mango"      // Menetapkan nilai "mango" ke indeks 1 slice fruits
	fruits[2] = "banana"     // Menetapkan nilai "banana" ke indeks 2 slice fruits
	fruits[3] = "watermelon" // Menetapkan nilai "watermelon" ke indeks 3 slice fruits

	fmt.Printf("%#v", fruits) // Mencetak slice fruits dengan menggunakan format %#v yang menghasilkan representasi string dari nilai slice tersebut
}


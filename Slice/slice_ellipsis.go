/*
Slice (append function with ellipsis)Jika kita ingin memasukkan seluruh element-element padasuatu array ke dalam array lainnya, maka kita dapatmenggunakan tanda ellipsis (...) atau tanda titik tiga berurut.Contohnya seperti pada gambar pertama di sebelah kanan.Terdapat 2 buah variable bernama fruits1 dan fruits2 yangmasing-masing memiliki tipe data slice of string danmenyimpan nama-nama buah.Lalu variable fruits1 mencoba untuk menambahkan seluruhelement yang terdapat pada variable fruits2, dan memakaitanda ellipsis untuk mengambil seluruh elementnya.Jika dijalankan pada terminal, maka hasilnya akan  sepertipada gambar kedua disebelah kanan.

*/

package main // Ini adalah deklarasi package, yang menandakan bahwa file ini adalah bagian dari package "main".

import "fmt" // Import package "fmt" untuk digunakan dalam program ini, fmt digunakan untuk format input dan output.

func main() { // Ini adalah fungsi utama dari program Go.

	var fruits1 = []string{"Salak", "Avokado", "Tomate"} // Mendeklarasikan slice of string 'fruits1' dan menginisialisasinya dengan beberapa buah.

	var fruits2 = []string{"Durian", "Kiwi", "Pear"} // Mendeklarasikan slice of string 'fruits2' dan menginisialisasinya dengan beberapa buah.

	fruits1 = append(fruits1, fruits2...) // Menggunakan fungsi 'append' untuk menggabungkan isi slice 'fruits2' ke slice 'fruits1'.

	fmt.Printf("%#v", fruits1) // Mencetak isi dari slice 'fruits1' dengan format %#v (format untuk menampilkan representasi nilai Go yang sesuai dengan sintaks Go).
}

// Program ini menggabungkan dua slice string 'fruits1' dan 'fruits2', dan mencetak slice 'fruits1' yang sudah digabungkan.


//[]string{"Salak", "Avokado", "Tomate", "Durian", "Kiwi", "Pear"}
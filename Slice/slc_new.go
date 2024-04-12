//Slice (Creating a new backing array)

package main 

import "fmt"

func main() {
	// Membuat slice 'cars' yang berisi beberapa merek mobil.
	cars := []string{"Ford", "Honda", "Mustibisa", "Hyundai"}

	// Membuat slice kosong 'newCars'.
	newCars := []string{}

	// Menyalin elemen dari slice 'cars' yang berada di indeks 0 hingga 1 (inklusif)
	// ke dalam slice 'newCars' menggunakan fungsi append.
	newCars = append(newCars, cars[0:2]...)

	// Mengubah nilai elemen di indeks 0 dari slice 'cars' menjadi "Nissan".
	cars[0] = "Nissan"

	// Mencetak isi slice 'cars' setelah perubahan.
	fmt.Println("cars:", cars)

	// Mencetak isi slice 'newCars' setelah ditambahkan elemen dari 'cars'.
	fmt.Println("newCars:", newCars)
}


/*
Penjelasan :
Ketika kita ingin mendapatkan element-element dari slice yang sudah ada, namun kita juga ingin membuat backing array yang baru, maka kita dapat menggunakan fungsi append untuk melakukannya. Contohnya seperti pada gambar pertama di sebelah kanan. Variable newCars ingin mendapatkan element-element dimulai dari index ke-0 hingga index ke-1 dari variable cars. Namun newCars mendapatkan element-elementnya dengan menggunakan fungsi append walaupun masih menggunakan slicing di dalam fungsi append nya. Ketika index ke-0 dari cars dirubah, maka newCars tidak ikut terubah dikarenakan mereka tidak memiliki backing array yang sama. Jika dijalankan di terminal kita, maka hasilnya akan seperti pada gambar kedua.
Ini adalah program sederhana dalam bahasa Go. Ini membuat dua slice dari string (cars dan newCars) dan menunjukkan bagaimana slice dan fungsi append bekerja di Go.

Pada awalnya, slice cars dibuat dengan beberapa merek mobil.
Kemudian, slice kosong newCars dibuat.
Elemen dari indeks 0 hingga 1 (inklusif) dari slice cars (yaitu "Ford" dan "Honda") disalin ke slice newCars menggunakan fungsi append.
Kemudian, nilai elemen di indeks 0 dari slice cars diubah menjadi "Nissan".
Setelah itu, isi slice cars dan newCars dicetak untuk melihat hasilnya.


*/
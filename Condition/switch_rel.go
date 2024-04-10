/*
Switch with relational operators
Kita juga dapat menggunakan switch dengan operator perbandinganseperti layaknya pada sebuah kondisional dengan keyword if, else ifdan else.Contohnya seperti pada gambar disebelah kanan. Kondisional padagambar disebelah kanan akan menghasilkan kalimat “not bad”karena variable score memiliki angka 6 dan angka 6 memang lebihkecil dari angka 8 dan lebih besar daripada angka 3.Lalu kita juga bisa melihat bahwa scope block default pada gambardisebelah kanan menggunakan kurung kurawal.Hal ini sangat bagus diterapkan jika kita ingin membuat lebih darisatu statement pada scope block default agar syntax kita lebih rapidan mudah di maintain.

*/

package main

import "fmt"

func main() {
	// Deklarasi variabel score dan inisialisasi dengan nilai 6
	var score = 6 

	// Penggunaan switch tanpa ekspresi switch itu sendiri
	switch {
	// Jika nilai score sama dengan 8, cetak "Perfect"
	case score == 8:
		fmt.Println("Perfect")
	// Jika nilai score kurang dari 8 dan lebih dari 3, cetak "Not Bad"
	case (score < 8) && (score > 3):
		fmt.Println("Not Bad") // Not Bad

	// Jika tidak memenuhi kondisi di atas, jalankan blok default
	default:
		{
			fmt.Println("Study Harder")
			fmt.Println("You Need to Learn More!")
		}
	}
}
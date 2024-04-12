/*
Loopings (break and continue keywords)
Sebelumnya kita telah menerapkan keyword break untuk menghentikan sebuah proses looping, dan sekarang kita jugaakan memakai keyword continue yang digunakan untuk melanjutkan suatu proses looping. Contohnya seperti padagambar pertama dibawah.Pada looping di bawah, terdapat dua kondisional yang dimana pada kondisional pertama digunakan untuk memeriksajika variable i memiliki nilai ganjil, maka proses loopingnya dipaksa berlanjut dan akan mengacuhkan syntax yang adadibawah keyword continue (pada kasus kita sekarang berarti kondisional kedua dan fungsi fmt.Println pada loopingdibawah).

*/

package main 

import "fmt"

func main () {
	for i := 1; i <=10; i++ { // Membuat loop dari 1 hingga 10.
		if i%2 == 1 { // Jika i adalah bilangan ganjil, lewati iterasi ini.
			continue // Melanjutkan ke iterasi berikutnya dalam loop.
		}
		if i > 8 { // Jika i lebih besar dari 8,
			break // Keluar dari loop.
		}

		fmt.Println("Angka", i) // Cetak nilai i.
	}
}

/*
Penjelasan :
Lalu pada kondisional kedua digunakan untuk memeriksa jika variable i sudah memiliki nilai diatas angka 8, makakeyword break akan terpanggil dan proses looping akan berhenti. Jika kita jalankan pada terminal kita, maka hasilnyaakan seperti pada gambar kedua di bawah.


Angka 2
Angka 4
Angka 6
Angka 8
*/
/*
Switch (fallthrough keyword)
Ketika kita ingin switch pada bahasa Go melanjutkan pengecekankepada case selanjutnya walaupun suatu case telah terpenuhikondisinya, maka kita bisa menggunakan keyword fallthrough.Jika kita perhatikan pada gambar disebelah kanan, kondisionalswitch tersebut akan menghasilkan 2 kalimat yaitu “not bad” dan “Itis ok, but you need to study harder” karena ketika kondisi pada casekedua terpenuhi, maka dia pengecekan akan berlanjut kepada caseketiga karena case kedua menggunakan keyword fallthrough.Ketika berlanjut ke case ketiga, kondisinya terpenuhi karena variablescore memiliki nilai berupa angka 6 dan angka 6 memang lebih besardari angka 5.

*/

package main 

import "fmt"

func main() {
	// Deklarasi variabel score dan inisialisasi dengan nilai 6
	var score = 6

	// Penggunaan switch tanpa ekspresi switch itu sendiri
	switch {
	// Jika nilai score sama dengan 8, cetak "perfect"
	case score == 8:
		fmt.Println("perfect")
	// Jika nilai score kurang dari 8 dan lebih dari 3, cetak "not bad" dan lanjut ke case selanjutnya dengan fallthrough
	case (score < 8) && (score > 3):
		fmt.Println("not bad") // not bad
		fallthrough
	// Jika nilai score kurang dari 5, cetak "It is ok, but please study harder"
	case score < 5:
		fmt.Println("It is ok, but please study harder") // it is ok, but please study harder
	// Jika tidak memenuhi kondisi di atas, jalankan blok default
	default:
		{
			fmt.Println("study harder")
			fmt.Println("You dont have a good score yet")
		}
	}
}
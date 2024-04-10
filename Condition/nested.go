/*
Nested Conditions
Kita juga dapat melakukan kondisional bersarang seperti pada gambar disebelah kanan. Kondisional bersarang merupakan sebuah proses kondisional yang didalamnya terdapat proses kondisional kembali. Kita dapat menggunakan if, else if, else, switch ataupun menggabungkanseluruhnya.
*/

package main 

import "fmt"

func main() {
	// Deklarasi variabel score dan inisialisasi dengan nilai 10
	var score = 10 

	// Kondisi utama: Jika score lebih besar dari 7
	if score > 7 {
		// Switch case di dalam kondisi utama
		switch score {
		// Jika score sama dengan 10, cetak "Perfect"
		case 10:
			fmt.Println("Perfect") // Perfect
		// Jika tidak memenuhi case di atas, cetak "Nice!"
		default:
			fmt.Println("Nice!")
		}
	} else {
		// Jika score tidak lebih besar dari 7
		// Kondisi dalam else
		if score == 5 {
			fmt.Println("Not Bad")
		} else if score == 3 {
			fmt.Println("Keep Trying")
		} else {
			// Jika score tidak sama dengan 5 atau 3
			fmt.Println("You can do it")
			// Kondisi bersarang dalam else
			if score == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
// Perfect 
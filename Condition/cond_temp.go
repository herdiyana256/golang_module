/*
Conditions
Kondisional pada pemrograman dapat digunakan untuk mengatur alur dari suatu program. Analogi yang bisa kita pakaipada kehidupan sehari-hari adalah analogi lampu merah, yang dimana lampu merah digunakan untuk mengatur lalulintas. Misalkan kita mendapat lampu merah maka kendaraan kita harus berhenti, misalkan kita mendapat lampu kuningmaka kendaraan kita harus bersiap-siap dan jika kita mendapatkan lampu hijau maka kita harus menjalankan kendaraankita. Sama halnya di dalam pemrograman, yang dimana dengan menggunakan conditions maka kita dapat mengaturprogram kita untuk mengetahui kapan harus mengeksekusi suatu kode.Pada bahasa Go, terdapat 2 macam kondisional yang dapat kita pakai yaitu if else dan switchhttps://www.flocabulary.com/unit/coding-conditionals/

*/

package main

import "fmt"

func main() {
	var currentYear = 2024 

	if age := currentYear - 1996; age < 25 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}
}
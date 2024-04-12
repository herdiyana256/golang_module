// Memeriksa keberadaan key 

package main 

import "fmt"

// Deklarasi fungsi main sebagai entry point dari program.
func main() {
	// Deklarasi variabel myMap sebagai map yang memetakan string ke integer.
	var myMap map[string]int
	
	// Inisialisasi map myMap menggunakan fungsi make untuk alokasi memori.
	myMap = make(map[string]int)

	// Menambahkan entry ke map myMap dengan key "Herdiyan" dan value 1.
	myMap["Herdiyan"] = 1
	// Menambahkan entry ke map myMap dengan key "Adam" dan value 2.
	myMap["Adam"] = 2
	// Menambahkan entry ke map myMap dengan key "Putra" dan value 3.
	myMap["Putra"] = 3

	// Mengecek apakah key "Founder" ada di dalam map myMap.
	// _ adalah variabel dummy untuk menampung value yang tidak digunakan dalam pemanggilan fungsi.
	// ok adalah variabel boolean yang menyatakan apakah key ditemukan atau tidak.
	_, ok := myMap["Founder"]
	// Jika key "Founder" ditemukan, cetak pesan bahwa key tersebut ada.
	if ok {
		fmt.Println("Key 'Founder' exists.")
	} else {
		// Jika key "Founder" tidak ditemukan, cetak pesan bahwa key tersebut tidak ada.
		fmt.Println("Key 'Founder' does not exist")
		// Pesan yang dicetak jika key "Founder" tidak ditemukan.
	}
}


// Key 'Founder' does not exist
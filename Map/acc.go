// Mengakses elemen map 

package main 

import "fmt"

func main() {
	var myMap map[string]int // Deklarasi variabel myMap sebagai map yang memetakan string ke integer.
    myMap = make(map[string]int) // Inisialisasi myMap menggunakan fungsi make untuk membuat map baru dengan tipe map[string]int.

    myMap["Herdiyan"] = 10 // Menambahkan elemen ke map myMap dengan key "Herdiyan" dan nilai 10.

    fmt.Println(myMap["Herdiyan"]) // Mencetak nilai yang terkait dengan key "Herdiyan" pada map myMap. Outputnya akan menjadi 10.
}

// Output : 10 
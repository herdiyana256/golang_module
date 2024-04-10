// Pernyataan switch digunakan untuk mengeksekusi kode berdasarkan nilai variabel.


package main 

import "fmt"


func main() {
	day := "Senin" // hari ini senin 

	switch day {
	case "Senin":
		fmt.Println("Hari ini Senin.")
	case "Selasa":
		fmt.Println("Hari ini Selasa.")
	case "Rabu":
		fmt.Println("Hari ini Rabu.")
	case "Kamis":
		fmt.Println("Hari ini Kamis.")
	case "Jumat":
		fmt.Println("Hari ini Jumat.")
	case "Sabtu":
		fmt.Println("Hari ini Sabtu.")
		default:
			fmt.Println("Hari tidak valid.")
	} // Kode di atas akan mencetak "Hari ini Senin." jika nilai day sama dengan "Senin".


}
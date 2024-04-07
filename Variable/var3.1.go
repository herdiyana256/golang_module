//Multiple variable declarations

/*
Another example

Jika kita ingin memberikan nilai dengan tipe data yang berbeda-beda, kita dapat menerapkan type inference dengan keyword var maupun dengan short declaration. Contohnya seperti pada gambar pertama di sebelah kanan.Jika kita jalankan pada terminal kita, maka hasilnya akan seperti pada gambar kedua.

*/

package main 


import "fmt"

func main () {
	var name, age, address = "Herdiyan", 28, "Jalan Jati Kelapa No.31"

	first, second, third := "1", 2, "3"

	fmt.Println(name, age, address) // Herdiyan 28 Jalan Jati Kelapa No.31

	fmt.Println( first, second, third) // 1 2 3 
}
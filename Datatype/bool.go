/*
Data Types - Sesi 1 BoolTipe 
data bool pada bahasa Go hanya terdiri dari 2 nilai yaitu false, dan true. Pada umumnya tipe data ini digunakan dalam perkondisian atau perulangan. Jika kita melihat pada gambar pertama dibawah, terdapat sebuah variable yang memiliki tipe data bool dengan nilai true. Lalu kita mencoba untuk mencetak nilai dengan tipe data bool tersebut menggunakan fungsi fmt.Printf dengan verb %t. Verb %t digunakan untuk memformat tipe data bool menjadi tipe data string. Dan jika kita jalankan pada terminal kita, maka hasilnya akan seperti pada gambar kedua dibawah.


*/

package main 


import "fmt"

func main () {
	var condition bool = true
	var situation bool = false 

	fmt.Printf("is it permitted? %t\n", condition)
	fmt.Printf("is it situation? %t\n", situation)
}

/*
is it permitted? true
is it situation? false
*/
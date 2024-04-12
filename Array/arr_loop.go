/*

Array - Sesi 2 Array(Loop through elements) Ada 2 cara penulisan agar kita dapat melakukan looping untukmengakses element-element yang terdapat pada sebuah array.Perhatikan pada gambar pertama disebelah kanan. Carapertamanya adalah dengan menggunakan range loop.Dengan menggunakan range loop kita bisa dengan mudahmengakses seluruh element pada sebuah array dengan bisa jugasekaligus mendapatkan indexnya.Lalu cara keduanya adalah kita bisa menggunakan looping biasadan kita perlu menggunakan bantuan fungsi len untukmendapatkan panjang dari array nya dengan cara penulisanlen(<nama array>). Jika kita jalankan syntax pada gambarpertama, maka hasilnya akan seperti pada gambar kedua disebelah kanan.
*/

package main // Mendeklarasikan bahwa ini adalah package utama

import (
	"fmt"     // Import package fmt untuk formatting dan mencetak output
	"strings" // Import package strings untuk bekerja dengan string
)

func main() { // Mendefinisikan fungsi main, fungsi utama dalam program Go
	var fruits = [3]string{"Apple", "Banana", "Mango"} // Mendeklarasikan array fruits dengan tipe data string dan panjang 3

	// cara pertama: menggunakan for range loop
	for i, v := range fruits { // Looping melalui array fruits dengan menggunakan perulangan for range
		fmt.Printf("Index: %d, Value:%s\n", i, v) // Mencetak indeks dan nilai dari setiap elemen array fruits
	}
	fmt.Println(strings.Repeat("#", 25)) // Mencetak deretan karakter "#" sebanyak 25 kali sebagai pemisah antara dua cara

	// Cara kedua: menggunakan for loop biasa
	for i := 0; i < len(fruits); i++ { // Looping melalui array fruits menggunakan perulangan for biasa dengan inisialisasi i = 0, kondisi i < panjang array, dan penambahan i++
		fmt.Printf("Index :%d ,Value:%s\n", i, fruits[i]) // Mencetak indeks dan nilai dari setiap elemen array fruits
	}
}



/*
Output ":
Index: 0, Value:Apple
Index: 1, Value:Banana
Index: 2, Value:Mango
#########################
Index :0 ,Value:Apple
Index :1 ,Value:Banana
Index :2 ,Value:Mango


Penjelasan :


*/

	
/*
3. fmt.printf():

Fungsi ini menyediakan format string yang memungkinkan Anda mengontrol bagaimana nilai argumen dicetak.
Menggunakan format specifier seperti %s untuk string, %d untuk integer, %f untuk floating-point, dll.
Opsional menambahkan baris baru dengan menyertakan \n di format string.
Digunakan untuk mencetak output dengan format yang lebih spesifik dan terkontrol.
Contoh:
*/

package main 

import "fmt"


func main () {
	name := "Herdiyan Adam"
	age  := 28

	fmt.Printf("Nama: %s, Usia: %d\n", name, age)
	fmt.Printf("Nama: %s, Usia: %d", name, age)
}
/*
Nama: Herdiyan Adam, Usia: 28
Nama: Herdiyan Adam, Usia: 28

*/


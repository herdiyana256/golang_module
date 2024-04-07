//Variable without data type 

// Variable without data type/ type inference 

// Kita juga dapat membuat suatu variable  dengan tidak memberikan tipe datanya secara ekplisit, hal ini dikenal dengan nama type inferrence yang dimana sistem dari golang yang akan menentukan sendiri tipe datanya secara otomatis

package main 

import "fmt"

func main () {
	var name = "Herdiyan"
	var age = 28 


	fmt.Printf("%T, %T", name, age)
	// %T adalah format specifier yang digunakan untuk mencetak tipe data.
	// kita sekarang menggunakan fungsi fmt.Printf. Fungsi fmt.Printf mempunyai kegunaan yang sama dengan fungsi fmt.Println.Hanya saja struktur output dari fungsi fmt.Printf ditentukan dari flag yang kita berikan. Hal ini akan kita bahas nanti. Jika kita coba menjalankan gambar pertama, maka kita akan mendapatkan keterangan tipe data kedua variable yang kita buat. Variable name akan secara otomatis memiliki tipe data string dan variable age akan mempunyai tipe data int.

}
// Output : string, int

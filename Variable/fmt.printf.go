/*
fmt.Printf Usage
Sebelumnya kita telah menggunakan fungsi fmt.Printf untuk mengetahui tipe data suatu variable dan kali ini akan kita bahas sedikit mengenai fungsi fmt.Printf tersebut. Struktur output yang dihasilkan oleh fungsi fmt.Printf akan bergantung dari flag yang kita pakai. Seperti contohnya jika kita ingin mengetahui tipe data dari suatu variable, maka kita dapat menggunakan verb% T. Contohnya seperti pada gambar pertama di sebelah kanan. Jika kita jalankan maka hasilnya akan seperti pada gambar kedua.Jika kita perhatikan pada gambar pertama, kita juga memakai simbol \n yang dimana simbol ini dapat kita pakai untuk membuat baris baru. Karena pada dasarnya, fmt.Printf tidak akan membuat baris baru tidak seperti fmt.Println yang menambah baris baru secara otomatis.


*/

package  main 

import "fmt"

func main () {

	first, second := 1, "2"
	fmt.Printf("Tipe data variable first adalah %T \n", first)
	fmt.Printf("Tipe data variable second adalah %T \n", second)
}

/*
Tipe data variable first adalah int 
Tipe data variable second adalah string 
*/
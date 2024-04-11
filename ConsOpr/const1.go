/*
Constants & Operators - Sesi 1 ConstantConstant (const) atau Konstanta merupakan jenis variable pada bahasa Go yang nilainya tidak dapat diubah. Contohnya jika kitamemiliki nilai-nilai tetap seperti PI, kecepatan cahaya, luas lingkaran dan lain-lain yang merupakan nilai-niali tetap.Maka kita bisa menyimpan nilai-nilai tersebut dengan variable const. Seperti yang kita lihat pada gambar pertama dibawah, kitamempunyai suatu variable const yang memiliki tipe data string. Perlu diingat disini bahwa ketika kita membuat variable denganconst, maka kita harus langsung memberikan nilai kepadanya. Karena jika tidak maka akan timbul error pada saat compile timeseperti pada gambar kedua dibawah.

*/

package main 

import "fmt"

func main () {
	const full_name string = "Herdiyan Adam Putra"

	fmt.Printf("Hello, %s", full_name)
}//Hello, Herdiyan Adam Putra

// atau bisa juga seperti ini 
fmt.Println(full_name)
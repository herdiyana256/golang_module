/* 
Slice
Slice merupakan suatu tipe data yang mirip dengan tipe data array, yang juga memiliki kegunaan untuk menyimpan satu ataulebih data. Namun tipe data slice dan array memiliki sifat yang berbeda. Slice tidak memiliki sifat fixed-length  yang berartipanjang dari slice tidak tetap sehingga kita bisa dengan leluasa menentukan panjang dari slice nya. Slice termasuk dalamkategori reference type yang dimana jika kita melakukan copy terhadap suatu slice, dan kita ubah element dari yang kitacopytersebut, maka slice semulanya juga akan ikut terubah.Cara membuat slice cukup mudah hampir mirip dengan jika kita membuat suatu array. Yang menjadi perbedaan adalah kitatidak perlu menuliskan panjang dari slice nya tidak seperti array. Contohnya seperti pada gambar dibawah ini.

*/

package main 
import "fmt"
func main () {
	var fruits = []string{"Apple", "Samsung", "Realme"}

	_ = fruits

	fmt.Println(fruits) // [Apple Samsung Realme]
}
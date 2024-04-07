/*
Underscore variableBahasa 
Golang memiliki satu aturan yang cukup strict/ketat, yang dimana tidak ada variable yang boleh menganggur ketika sudah kita buat. Contohnya pada gambar pertama di sebelah kanan. Jika kita perhatikan pada gambar pertama, terjadi error pada saat compile time yang ditandai dengan garis merah. Ini terjadi karena kita telah membuat dan mendeklarasikan variable tetapi variable-variable tersebut kita biarkan menganggur atau tidak dipakai. Maka dari itu cara menanggulanginya adalah dengan memakai variable underscore seperti pada gambar kedua. Dengan variable underscore maka kita dapat menghindari error yang akan terjadi jika kita mempunyai suatu variable yang menganggur atau tidak dipakai.


*/

// example saat variable ada yang tidak digunakan 
// package main 


// import "fmt"

// func main () {

// 	var firstVariable string 

// 	var name, age, address = "Herdiyan", 28, "Jalan Jati Kelapa No.31"

// }


// Solusi 
package main 

import "fmt"

func main () {
	var firstVariable string

	 name, age, address := "Herdiyan", 28, "Jalan Jati Kelapa"

	_,_,_,_ = firstVariable, name, age, address

	fmt.println(name, age, address)
}
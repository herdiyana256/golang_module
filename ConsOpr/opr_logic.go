/*
Operator Logika pada Golang
Operator logika dalam Golang digunakan untuk menggabungkan dua nilai boolean (true atau false) menjadi satu nilai boolean. Berikut adalah daftar operator logika:

Operator	Keterangan	Contoh
&&	Dan (AND)	true && true menghasilkan true
`	`	Atau (OR)
!	Negasi (NOT)	!true menghasilkan false

*/


package main 

import "fmt"

func main() {
	var a bool = true 
	var b bool = false 

	fmt.Println("Apakah a dan b sama-sama true?", a && b)
	fmt.Println("Apakah a atau b sama dengan true?", a || b)
	fmt.Println("Apakah a tidak sama dengan true?", !a)
}
/*
Apakah a dan b sama-sama true? false
Apakah a atau b sama dengan true? true
Apakah a tidak sama dengan true? false
*/


/*
Summary 
Operator logika dapat digunakan dalam berbagai situasi, seperti:

Percabangan (if-else)
Perulangan (for, while)
Pengujian kondisi
Operator Logika Lainnya:

Golang juga memiliki beberapa operator logika lain yang lebih kompleks, seperti:

^: XOR (Exclusive OR)
&: Bitwise AND
|: Bitwise OR
<<: Bitwise left shift
>>: Bitwise right shift
Anda dapat mempelajari lebih lanjut tentang operator logika di dokumentasi resmi Golang: https://www.konsistensi.com/2014/03/mengatasi-angkettidak-valid.html

Kesimpulan:

Operator logika merupakan alat penting dalam Golang untuk menggabungkan nilai boolean dan menghasilkan nilai boolean. Memahami cara menggunakan operator logika dengan baik akan membantu Anda menulis kode yang lebih logis dan mudah dipahami.
*/
/*
Operator Relasional pada Golang
Operator relasional dalam Golang digunakan untuk membandingkan dua nilai dan menghasilkan nilai boolean (true atau false). Berikut adalah daftar operator relasional:

Operator	Keterangan	Contoh
==	Sama dengan	10 == 10 menghasilkan true
!=	Tidak sama dengan	10 != 5 menghasilkan true
<	Kurang dari	5 < 10 menghasilkan true
<=	Kurang dari atau sama dengan	5 <= 10 menghasilkan true
>	Lebih dari	10 > 5 menghasilkan true
>=	Lebih dari atau sama dengan	10 >= 10 menghasilkan true

*/


package main 


import "fmt"


func main () {

	var a int = 10 
	var b int = 5


	fmt.Println("Apakah a sama dengan b?", a == b)
	fmt.Println("Apakah a tidak sama dengan b?", a != b)
	fmt.Println("Apakah a kurang dari b?", a < b)
	fmt.Println("Apakah kurang dari atau sama dengan b?", a<=b)
	fmt.Println("Apakah a lebih dari b?", a > b )
	fmt.Println("Apakah a lebih dari atau sama dengan b?", a >= b)
}
/*
Apakah a sama dengan b? false
Apakah a tidak sama dengan b? true
Apakah a kurang dari b? false
Apakah kurang dari atau sama dengan b? false
Apakah a lebih dari b? true
Apakah a lebih dari atau sama dengan b? true
*/


/*
SUMMARY : 

Operator relasional dapat digunakan dalam berbagai situasi, seperti:

Percabangan (if-else)
Perulangan (for, while)
Pengujian nilai
Operator Relasional Lainnya:

Golang juga memiliki beberapa operator relasional lain yang lebih kompleks, seperti:

==: Membandingkan nilai dan tipe data
!=: Membandingkan nilai dan tipe data, tidak sama
<: Membandingkan nilai, tidak sama
<=: Membandingkan nilai, tidak sama atau sama dengan
>: Membandingkan nilai, sama dengan
>=: Membandingkan nilai, sama dengan atau tidak sama
Anda dapat mempelajari lebih lanjut tentang operator relasional di dokumentasi resmi Golang: [URL yang tidak valid dihapus]

Kesimpulan:

Operator relasional merupakan alat penting dalam Golang untuk membandingkan nilai dan menghasilkan nilai boolean. Memahami cara menggunakan operator relasional dengan baik akan membantu Anda menulis kode yang lebih logis dan mudah dipahami.

*/



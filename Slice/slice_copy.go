/*
Slice (copy function)Kita juga bisa menggunakan fungsi copy untuk meng-copyseluruh element pada sebuah slice ke dalam slice lainnya.Perlu diingat disini bahwa ketika kita mencoba untukmeng-copy sebuah slice kedalam slice lainnya, maka seluruhelement pada slice lainnya tersebut akan ter-replace olehelement-element yang di copy kannya.Contohnya seperti pada gambar pertama di sebelah kanan.

*/

package main 

import "fmt"

func main() {
	// Deklarasi dua slice string, fruits1 dan fruits2, dengan elemen yang berbeda.
	var fruits1 = []string{"Blackberry", "Blueberry", "Cerry"}
	var fruits2 = []string{"Dragon", "pineapple", "strawberry"}

	// Menggunakan fungsi copy() untuk menyalin elemen dari fruits2 ke fruits1.
	// nn adalah jumlah elemen yang berhasil disalin.
	nn := copy(fruits1, fruits2)

	// Mencetak isi fruits1, fruits2, dan jumlah elemen yang disalin.
	fmt.Println("Fruit1 =>", fruits1)
	fmt.Println("Fruits2 =>", fruits2)
	fmt.Println("Copied element =>", nn)
	fmt.Println(fruits1) // Mencetak isi fruits1 lagi.
}

//Penjelasan : Pada kasus kita kali ini, variable fruits1 ingin meng-copyseluruh element yang ada pada variable fruits2.Argumen pertama yang diterima oleh fungsi copy  adalahdestinasi atau slice yang ingin meng-copy, lalu argumentkedua adalah source/sumber dari slice yang ingin di copy.Fungsi copy juga akan mengembalikan jumlah element yangberhasil ter-copy.Ketika kita jalankan pada terminal kita, maka hasilnya akanseperti pada gambar kedua disebelah kanan. Element padafruits1 sudah ter-replace oleh fruits2, dan terdapat 3 elementyang berhasil ter-copy.


/*
Fruit1 => [Dragon pineapple strawberry]
Fruits2 => [Dragon pineapple strawberry]
Copied element => 3
[Dragon pineapple strawberry]
*/
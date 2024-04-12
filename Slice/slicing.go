//Slice (Slicing)Ada cara lain lagi agar kita dapat mendapatkan element-element darisebuah slice dan kita juga bisa menentukan element dari index keberapa yang ingin kita dapatkan. Caranya adalah denganmenggunakan slicing.Contohnya seperti pada gambar disebelah kanan. Cara penulisandalam melakukan slicing adalah sama dengan [start:stop]. Start samadengan awal index yang ingin kita akses dan stop berarti indexakhirnya. Perhatikan hasil dari syntax pada gambar pertama di gambarkedua.


package main 

import "fmt"

func main() {
	// Deklarasi slice `fruits1` dengan elemen string.
	var fruits1 = []string{"Apple", "Watermelon", "Mango", "Pearl", "Sirsak"}

	// Membuat slice baru `fruits2` yang merupakan subset dari `fruits1` dari indeks 1 hingga sebelum indeks 4.
	var fruits2 = fruits1[1:4]
	fmt.Printf("%#v\n", fruits2) // prints [Watermelon Mango Pearl]

	// Membuat slice baru `fruits3` yang merupakan subset dari `fruits1` dari indeks 0 hingga akhir.
	var fruits3 = fruits1[0:]
	fmt.Printf("%#v\n", fruits3) // prints [Apple Watermelon Mango Pearl Sirsak]

	// Membuat slice baru `fruits4` yang merupakan subset dari `fruits1` dari awal hingga sebelum indeks 3.
	var fruits4 = fruits1[:3]
	fmt.Printf("%#v\n", fruits4) // prints [Apple Watermelon Mango]

	// Membuat slice baru `fruits5` yang merupakan salinan dari seluruh `fruits1`.
	// Ini juga dapat ditulis sebagai fruits1[:len(fruits1)].
	var fruits5 = fruits1[:]
	fmt.Printf("%#v\n", fruits5) // prints [Apple Watermelon Mango Pearl Sirsak]
}


/*
[]string{"Watermelon", "Mango", "Pearl"}
[]string{"Apple", "Watermelon", "Mango", "Pearl", "Sirsak"}
[]string{"Apple", "Watermelon", "Mango"}
[]string{"Apple", "Watermelon", "Mango", "Pearl", "Sirsak"}

*/


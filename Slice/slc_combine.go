//Slice ( Combine slicing and append )

//Kita juga dapat mengkombinasikan fungsi append dengan slicing.Contohnya seperti pada gambar pertama di sebelah kanan. Jika kitaperhatikan, variable fruits1 ingin me-replace index ketiga hinggaseterusnya dengan hanya buah “rambutan” saja. Jika kita jalankanpada terminal maka hasilnya akan seperti pada gambar kedua disebelah kanan.


package main 

import "fmt"

func main() {
	var fruits1 = []string{"Apple", "Huawei", "Samsung", "Realme", "Xiomi"} // Inisialisasi slice string 'fruits1' dengan beberapa elemen.

	fruits1 = append(fruits1[:3], "Blackberry") // Mengganti elemen ke-3 ('Samsung') dengan "Blackberry" menggunakan slicing dan fungsi append.

	fmt.Printf("%#v\n", fruits1) // Mencetak nilai slice 'fruits1'.
}


// []string{"Apple", "Huawei", "Samsung", "Blackberry"}
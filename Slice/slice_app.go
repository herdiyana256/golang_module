//Karena kita ingin agar variable fruits yang bertambahelementnya, maka dari itu kita me-reassign variable fruitsdengan fungsi append. Parameter pertama yang diberikanpada fungsi append adalah slice yang ingin ditambahkan, laluparameter setelahnya adalah element-element yang inginditambahkan dan jangan lupa dipisahkan dengan koma.

package main // Ini menyatakan bahwa file ini adalah bagian dari package "main", yang merupakan package yang dieksekusi pertama kali saat program dijalankan.

import "fmt" // Import package "fmt" untuk melakukan formatting dan output.

func main() { // Fungsi utama dari program.
	var fruits = make([]string, 3) // Mendeklarasikan variabel "fruits" sebagai slice (array dinamis) string dengan kapasitas awal 3 elemen.

	fruits = append(fruits, "Apple", "Banana", "Mango") // Menambahkan elemen "Apple", "Banana", dan "Mango" ke slice "fruits" menggunakan fungsi "append".

	fmt.Printf("%#v", fruits) // Mencetak nilai dari slice "fruits" menggunakan formatting verb "%#v" untuk mencetak nilai lengkap dari slice tersebut.
}


//[]string{"", "", "", "Apple", "Banana", "Mango"}


// Noted : %#v digunakan untuk mencetak representasi nilai lengkap dari suatu variabel,
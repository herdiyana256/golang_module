// fmt.Print()

// Function fmt.Print memiliki peran yang sama seperto function fmt.Println, Pembedanya function fmt.Print tidak menghasilkan baris baru di akhir outputnya.

// Perbedaan lain nya adalah, nilai pada parameter-parameter yang dimasukkan ke fungsi tersebut digabungkan tanpa pemisah. Tidak seperti pada function fmt.Println yang nilai parameternya digabung menggunakan penghubung spaci.

// contoh aritmatika code
package main 

import "fmt"

func main() {
	// Penjumlahan 
	fmt.Println(10 + 5) // Output = 15


	//Pengurangan 
	fmt.Println(10 * 5) // Output = 50 

	// Pembagian 
	fmt.Println(10 / 5) // Output : 2 

	//Sisa Pembagian 
	fmt.Println(10 % 5) // Output : 0 

	// Operasi 
	fmt.Println(10 + 5 * 2) // Output : 20 

	// Pangkat
	fmt.Println(2 ^ 3)// Output : 1

	// Tanpa baris baru 
	fmt.Print("Herdiyan Adam")

	//Tanpa baris baru 
	fmt.Print("Putra")
}
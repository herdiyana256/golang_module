//3. Pernyataan else if:

//Anda dapat menggunakan pernyataan else if untuk mengeksekusi kode jika kondisi pertama tidak terpenuhi, dan seterusnya.

package main 

import "fmt"

func main () {
	age := 28 

	if age >= 28 {
		fmt.Println("Anda sudah tua")
	} else if age >=25 {
		fmt.Println("Anda ABG")
	} else {
		fmt.Println("Anda masih anak-anak")
	}
}
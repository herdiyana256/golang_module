// Mengiterasi array 

package main

import "fmt"

func main() {
    //// Mendeklarasikan sebuah array dengan panjang 5 dan tipe data int

	var myArray [5]int

	// Melakukan iterasi terhadap setiap elemen dalam array
for i := 0; i < len(myArray); i++ {
	// Mencetak nilai dari setiap elemen array
	fmt.Println(myArray[i])
  }

}

/*


0
0
0
0
0

*/
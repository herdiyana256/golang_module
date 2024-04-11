/*
Number (float)-floatFloat
 merupakan tipe data numerik desimal pada bahasaGo. Secara umum float dibagi menjadi 2 jenis yaitufloat32 dan float34.Perbedaan kedua jenis float tersebut berada di lebarcakupan nilai desimal yang bisa ditampung. Untuk lebihjelasnya bisa merujuk ke spesifikasiIEEE-754 32-bitfloating-point numbershttps://www.geeksforgeeks.org/data-types-in-go/

*/

package main 

import "fmt"

func main () {
	var decimalNumber float32 = 3.63 

	fmt.Printf("decimal Number: %f\n", decimalNumber)
	fmt.Printf("decimal Number: %.3f\n", decimalNumber) // 3f itu jumlah digit yang ingin kita hasilkan, n itu adalah nilai 
	fmt.Printf("decimal Number: %.2f\n", decimalNumber) 
}

/*
decimal Number: 3.630000
decimal Number: 3.630
decimal Number: 3.63
*/
// deklarasikan array 

package main 

import "fmt"


func main() {
	var myArray [5]int // array dengan 5 element integer
	myArray[0] = 1
	myArray[1] = 2 
	myArray[2] = 3
	myArray[3] = 4
	myArray[4] = 5

		//[1 2 3 4 5]
	fmt.Println(myArray) // Mencetak seluruh elemen array 
}
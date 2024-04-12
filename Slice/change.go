// Mengubah Slice 

package main 

import 	"fmt"

func main() {
	mySlice := []int{1,2,3,4,5}

	mySlice = mySlice[:3] // Memotong slice hingga element ke -3 

	// mySlice = mySlice[:4] // Memotong slice hingga element ke - 4

	// mySlice = mySlice[:5] // Memotong slice hingga element ke  - 5

	fmt.Println(mySlice) // [1 2 3]
}
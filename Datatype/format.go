package main 

import "fmt"

func main () {

	x := 100 
	y := "Hello, Ganteng!"
	z := 3.15

	fmt.Printf("Nilai x: %d\n", x) // 100
	fmt.Printf("Tipe data x: %T\n", x) // Tipe data x: int
	fmt.Printf("Nilai y: %s\n", y) // Nilai y: Hello, Ganteng!
	fmt.Printf("Tipe data y :%T\n", y) // Tipe data y :string
	fmt.Printf("Tipe data z: %T\n", z) // Tipe data z: float64
	fmt.Printf("Decimal number z : %.2f\n", z) // Decimal number z : 3.15



}
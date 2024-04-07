package main 

import "fmt"

// fungsi untuk mengecek apakah sebuah bilangan ganjil atau genap
func checkOddEven(number int) string {
	if number%2 == 0 {
		return "even"
	}

	return "odd"
}

// Fungsi untuk menghitung faktorial dari sebuah bilangan 
func factorial(n int) int  {
	if n == 0 {
		return 1 
	}
	return n * factorial(n-1)
}


func main () {
	// contoh penggunaan fungsi checkOddEven 
	num := 5
	fmt.Printf ("%d is %s\n", num, checkOddEven(num))

	// Contoh penggunaan fungsi factorial 
	n := 5
	fmt.Printf("Factorial of %d is %d\n", n, factorial(n))
}

/*
Output : 
5 is odd
Factorial of 5 is 120


*/
// Mendeklarasikan map

package main

import "fmt"

func main() {
	// Mendefinisikan tipe map dengan key string dan value integer
	var myMap map[string]int

	// Menginisialisasi map menggunakan fungsi make
	myMap = make(map[string]int)

	// Mengisi map dengan beberapa pasangan kunci-nilai
	myMap["Apple"] = 1
	myMap["Banana"] = 2
	myMap["Mango"] = 3
	myMap["Srikaya"] = 4

	// Mencetak isi map
	fmt.Println(myMap) //map[Apple:1 Banana:2 Mango:3 Srikaya:4]
}

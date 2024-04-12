// Mengiterasi map : proses mengulangi serangkaian instruksi atau elemen data dalam sebuah struktur data


package main 

import "fmt"

func main () {
// Mendefinisikan variabel myMap sebagai sebuah map yang memiliki kunci bertipe string dan nilai bertipe int.
var myMap map[string]int 

// Menginisialisasi map myMap dengan menggunakan fungsi make() untuk membuat map baru.
myMap = make(map[string]int)

// Menetapkan nilai pada map myMap dengan kunci "Apple" yang memiliki nilai 1.
myMap["Apple"] = 1 
// Menetapkan nilai pada map myMap dengan kunci "Banana" yang memiliki nilai 2.
myMap["Banana"] = 2 
// Menetapkan nilai pada map myMap dengan kunci "Orange" yang memiliki nilai 3.
myMap["Orange"] = 3 
// Menetapkan nilai pada map myMap dengan kunci "Watermelon" yang memiliki nilai 4.
myMap["Watermelon"] = 4 

// Iterasi melalui setiap elemen pada map myMap menggunakan loop for range.
for key, value := range myMap {
    // Mencetak kunci dan nilai dari setiap elemen map.
    fmt.Println(key, ":", value)
}
}

/*
Output :
pple : 1
Banana : 2
Orange : 3
Watermelon : 4


*/
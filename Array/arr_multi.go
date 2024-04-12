package main 

import "fmt"

func main() {
    // Deklarasi array 2 dimensi dengan 2 baris dan 3 kolom
    var myArray [2][3]int

    // Mengisi elemen-elemen array dengan nilai tertentu
    myArray[0][0] = 1  // Mengisi elemen pada baris pertama, kolom pertama dengan nilai 1
    myArray[0][1] = 2  // Mengisi elemen pada baris pertama, kolom kedua dengan nilai 2
    myArray[0][2] = 3  // Mengisi elemen pada baris pertama, kolom ketiga dengan nilai 3
    myArray[1][0] = 4  // Mengisi elemen pada baris kedua, kolom pertama dengan nilai 4
    myArray[1][1] = 5  // Mengisi elemen pada baris kedua, kolom kedua dengan nilai 5
    myArray[1][2] = 6  // Mengisi elemen pada baris kedua, kolom ketiga dengan nilai 6

    // Mencetak seluruh elemen array
    fmt.Println(myArray)
}

//[[1 2 3] [4 5 6]]


//Array adalah struktur data yang penting dalam Golang. Memahami cara menggunakan array dengan baik akan membantu Anda menulis kode yang lebih efisien dan mudah dipahami.


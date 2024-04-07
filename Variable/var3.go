// Multiple variable declarations


/*
Pada bahasa Golang, kita dapat membuat lebih dari satu variable pada satu baris/line yang sama. Cara nya seperti gambar pertama. Jika kita perhatikan, kita telah membuat 3 variable pada baris pertama dan membuat 3 variable pada baris kedua. Pada baris pertama, ketiga variable tersebut memiliki tipe data string dan kita langsung memberikan nilai kepadanya. Sedangkan pada baris kedua, ketiga variable tersebut memiliki tipe data int dan kita memberikan nilai kepadanya dengan cara me-reassign nilai-nilai kepada ketiga variable pada baris kedua tersebut.Dan ketika kita jalankan pada terminal kita, maka hasil nya akan seperti pada gambar kedua.
*/

package main 

import "fmt"

func main () {
	var student1, student2, student3 string = "satu", "dua", "tiga"

	var first, second, third int 


	first, second, third = 1,2,3 

	fmt.Println(student1, student2, student3) // satu dua tiga 

	fmt.Println(first, second, third) // 1 2 3 
}

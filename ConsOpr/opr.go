package main 

import "fmt"

func main () {
	// Operator Matematika
var a int = 10
var b int = 5

fmt.Println("Jumlah:", a + b)
fmt.Println("Selisih:", a - b)
fmt.Println("Perkalian:", a * b)
fmt.Println("Pembagian:", a / b)
fmt.Println("Sisa pembagian:", a % b)

// Operator Perbandingan
fmt.Println("Apakah a sama dengan b?", a == b)
fmt.Println("Apakah a tidak sama dengan b?", a != b)
fmt.Println("Apakah a kurang dari b?", a < b)
fmt.Println("Apakah a kurang dari atau sama dengan b?", a <= b)
fmt.Println("Apakah a lebih dari b?", a > b)
fmt.Println("Apakah a lebih dari atau sama dengan b?", a >= b)

// Operator Logika
fmt.Println("Apakah a lebih dari 5 dan b kurang dari 10?", a > 5 && b < 10)
fmt.Println("Apakah a sama dengan 10 atau b sama dengan 5?", a == 10 || b == 5)
fmt.Println("Apakah a tidak sama dengan 10?", !(a == 10))
}

/*
Jumlah: 15
Selisih: 5
Perkalian: 50
Pembagian: 2
Sisa pembagian: 0
Apakah a sama dengan b? false
Apakah a tidak sama dengan b? true
Apakah a kurang dari b? false
Apakah a kurang dari atau sama dengan b? false
Apakah a lebih dari b? true
Apakah a lebih dari atau sama dengan b? true
Apakah a lebih dari 5 dan b kurang dari 10? true
Apakah a sama dengan 10 atau b sama dengan 5? true        
Apakah a tidak sama dengan 10? false

Summary : 
Konstanta dan operator merupakan elemen penting dalam Golang untuk mendefinisikan nilai yang tidak berubah dan melakukan operasi pada data. Memahami cara menggunakan konstanta dan operator dengan baik akan membantu Anda menulis kode yang lebih efektif dan mudah dibaca.




*/
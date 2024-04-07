// Me-reassign suatu variable dengan suatu nilai asalkan kita memberikan nilai dengan tipe data yang sama dengan tipe data yang sudah dimiliki oleh variable tersebut.
// Contohnya seperti pada gambar sebelah kanan 

package main 

import "fmt"


func main() {
	var name string 
	var age int = 28 

	name = "Herdiyan"
	age = 28

	fmt.Println("Ini adalah namanya :", name)
	fmt.Println("Ini adalah umurnya :", age)
}
/*
Ini adalah namanya : Herdiyan
Ini adalah umurnya : 28
*/

/*

Penjelasan : 
 Jika kita perhatikan pada gambar pertama, kita tidak memberikan nilai apapun pada variable nama pada pertama kali kita membuat varible tersebut.
 Hal ini dapat dilakukan asalkan kita sudah memberikan tipe datanya.
 Lalu kita me-reassign variable name ydan age dengan suatu nilai baru dan tidak ada error yang terjadi. 
 Hal ini dikarenakan kita masih meberikan nilai dengan nilai baru dan tidak ada error yang terjadi.

 Hal ini dikarenakan kita masih memberikan nilai dengan tipe data yang sama seperti pertama kali variable name dan age dibuat 
*/



//Namun jika kita tidak memberikan nilai dengan tipe data yang sama pada variable name dan age, maka akan terjadi error 
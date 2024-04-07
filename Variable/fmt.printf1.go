/*
fmt.Printf Usage
Kita juga dapat memakai lebih dari satu verb, misalkan kita ingin memasukkan berbagai macam variable dengan tipe data berbeda-beda ke dalam satu kalimat. Maka kita bisa melakukannya seperti pada gambar pertama disebelah kanan. Verb%s digunakan untuk memformat suatu nilai dengan tipe data string sedangkan verb %d digunakan untuk memformat tipe data int dengan base 10 atau angka 0 sampai 9. Jika kita jalankan maka akan terlihat seperti pada gambar dibawah ini. Jika teman-teman ingin mengetahui tentang verb yang lainnya, maka bisa mengunjungi link berikut https://pkg.go.dev/fmt.
*/

package main 

import "fmt"

func main () {
	var name = "Herdiyan"

	var age = 28 

	var address = "Jalan Jati Kelapa No.31"

	fmt.Printf("Hello my name %s, umur saya adalah %d, saya tinggal di %s", name, age, address)
}
/*

Hello my name Herdiyan, umur saya adalah 28, saya tinggal di Jalan Jati Kelapa No.31

*/
/*
Array(Modify Element Through Index) Kita juga dapat memodifikasi element-element yang terdapatdalam sebuah array dengan cara mengakses indexnya. Contohnyaseperti pada gambar pertama disebelah kanan kita.Terdapat sebuah variable dengan nama fruits yang memiliki arrayyang isinya adalah nama buah-buahan menggunakan bahasaIndonesia.Lalu buah-buahan yang terdapat pada array dalam variable fruitstersebut di modifikasi dengan cara mengakses indexnya danmengganti buah-buahannya ke dalam bahasa Inggris. Dan jikadijalankan hasilnya seperti pada gambar kedua.

*/

package main 

import "fmt"

func main () {
	// Mendefinisikan array 'fruits' dengan panjang 3 dan tipe data string
	var fruits = [3]string{"Apel Jin", "Sri Kaya", "Mau ngga"}

	// Mengubah nilai elemen pertama dari "Apel Jin" menjadi "apple"
	fruits[0] = "apple"

	// Mengubah nilai elemen kedua dari "Sri Kaya" menjadi "Srikaya"
	fruits[1] = "Srikaya"

	// Mengubah nilai elemen ketiga dari "Mau ngga" menjadi "Mango"
	fruits[2] = "Mango"


	// Mencetak isi array 'fruits' dengan format %#v untuk menampilkan representasi Go syntax dari array
	fmt.Printf("%#v\n", fruits)
}


/*
Penjelasan :
Array(Modify Element Through Index) Kita juga dapat memodifikasi element-element yang terdapatdalam sebuah array dengan cara mengakses indexnya. Contohnyaseperti pada gambar pertama disebelah kanan kita.Terdapat sebuah variable dengan nama fruits yang memiliki arrayyang isinya adalah nama buah-buahan menggunakan bahasaIndonesia.Lalu buah-buahan yang terdapat pada array dalam variable fruitstersebut di modifikasi dengan cara mengakses indexnya danmengganti buah-buahannya ke dalam bahasa Inggris. Dan jikadijalankan hasilnya seperti pada gambar kedua.




[3]string{"apple", "Srikaya", "Mango"}
go: remove C:\Users\herdi\AppData\Local\Temp\go-build3046653252\b001\exe\arr_modi.exe: The process cannot access the file because it is being used by another process

*/
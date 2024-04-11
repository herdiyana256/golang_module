/*
Loopings (first way of looping)
Looping atau perulangan merupakan suatu proses berulang yang dimana proses tersebut akan berhenti jika memenuhisuatu kondisi. Bahasa Go hanya memiliki satu looping yaitu looping dengan menggunakan keyword for atau yang kitakenal dengan sebutan for loop.Contohnya seperti pada gambar dibawah ini.

*/
package main

import "fmt"

func main() {
    // Loop dimulai dari i = 0; 
    // Loop akan terus berjalan selama i < 3;
    // Setiap iterasi, nilai i akan bertambah 1
	for i := 0; i < 3; i++ {
        // Mencetak teks "Angka" diikuti dengan nilai i pada setiap iterasi
		fmt.Println("Angka", i)
	}
}
/*
Penjelasan code :
Jika kita lihat pada gambar pertama dibawah, ini merupakan cara pertama agar kita dapat melakukan looping pada bahasaGo. Looping tersebut akan bekerja selama variable i memiliki nilai kurang dari angka 3.Jika kita jalankan maka hasilnya akan seperti pada gambar kedua dibawah. Perlu diingat disini bahwa kita tidak boleh lupauntuk menambahkan variable  i seperti pada gambar pertama yang dimana i++ sama saja dengan i = i + 1. Jika variable itidak ditambahkan maka akan menimbulkan infinite loop atau proses perulangan yang tidak akan berhenti.


Angka 0
Angka 1
Angka 2
*/
/*
Loopings - Sesi 2 Loopings (second way of looping)Cara kedua dalam melakukan looping pada bahasa Go adalah dengan menyelipkan kondisional seperti padalooping dengan menggunakan keyword while.Contohnya seperti pada gambar pertama dibawah. Looping pada gambar pertama dibawah menggunakankondisi yang dimana jika nilai yang dimiliki oleh variable i masih kurang dari angka 3, maka proses loopingtersebut akan terus berlanjut. Namun jika sudah melebihi dari angka 3, maka proses looping tersebut akanberhenti. Jika kita menjalankan pada terminal kita, maka hasilnya akan seperti pada gambar kedua dibawah.


*/

package main

import "fmt"

func main() {
    var i = 0 // Mendeklarasikan variabel i dengan nilai awal 0

    // Looping tak terbatas (infinite loop) menggunakan for tanpa kondisi
    for {
        fmt.Println("Angka", i) // Mencetak nilai i

        i++ // Increment nilai i
        if i == 3 { // Jika nilai i sama dengan 3
            break // Keluar dari loop
        }
    }
}

/*

Angka 0
Angka 1
Angka 2


Loopings (third way of looping)
Cara ketiga dalam melakukan looping pada bahasa Go adalah dengan melakukan looping tanpa memberikan kondisiapapun. Contohnya seperti pada gambar pertama dibawah. Looping pada gambar dibawah memakai bantuan keywordbreak yang dimana dengan menggunakan keyword ini, maka dapat menghentikan suatu proses looping.

*/
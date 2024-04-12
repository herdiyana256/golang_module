/*

Loopings - Sesi 2 Loopings (Label)
Pada looping bersarang, keyword break dan continue akan berlaku pada blok perulangan dimana ia digunakan saja.Adacara agar kedua keyword ini bisa tertuju pada perulangan terluar atau perulangan tertentu, yaitu dengan memanfaatkanteknik pemberian label. Contohnya seperti pada gambar pertama dibawah.


*/

//menggunakan konsep label untuk mengontrol aliran eksekusi program dengan lebih detail.

package main

import "fmt"

func main() {
    // Label outerLoop digunakan untuk mengidentifikasi loop terluar.
    outerLoop:
    for i := 0; i < 3; i++ {
        fmt.Println("Perulangan ke - ", i+1)
        for j := 0; j < 3; j++ {
            // Ketika nilai i adalah 2, pernyataan berikut akan dieksekusi.
            if i == 2 {
                // Break outerLoop akan menghentikan iterasi loop terluar.
                break outerLoop
            }
            // Jika nilai i bukan 2, kode berikut akan dieksekusi.
            fmt.Print(j, " ")
        }
        // Setiap kali iterasi loop dalam selesai, baris baru akan dicetak.
        fmt.Print("\n")
    }
}


/*
Penjelasan :
ika kita perhatikan hasil dari looping pertama pada gambar kedua, seluruh looping pada perulangan ketiga terhenti karenaterdapat sebuah kondisional pada proses looping kedua yang dimana jika variable i sudah memiliki nilai dengan angkasama dengan 2, maka looping pertama atau looping terluar akan dihentikan atau sama saja dengan seluruh proses loopingterhenti.



Perulangan ke -  1
0 1 2 
Perulangan ke -  2
0 1 2 
Perulangan ke -  3
*/
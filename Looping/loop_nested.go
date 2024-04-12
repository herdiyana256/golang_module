/*
Loopings (Nested Looping)
Nested looping atau looping bersarang adalah suatu proses looping yang memiliki suatu proses looping di dalamnya.Contohnya seperti pada gambar pertama dibawah. Jika kita perhatikan, kita menggunakan fungsi fmt.Println dan tidakmemberikan argumen apapun kedalamnya. Ini bisa dilakukan ketika kita hanya ingin membuat baris baru. Hasilnya samasaja ketika kita memakai fungsi fmt.Print dengan argumen “\n” atau sama dengan ( fmt.Print("\n") ). Jika kita jalankan padaterminal kita, maka hasilnya akan seperti pada gambar kedua dibawah.


*/
package main // Deklarasi package utama

import "fmt" // Import package fmt untuk formatting dan output

func main() { // Deklarasi fungsi main

    // Looping pertama, iterasi dari 0 hingga 4
    for i := 0; i < 5; i++ {
        
        // Looping kedua, iterasi dari nilai i hingga 4
        for j := i; j < 5; j++ {
            fmt.Print(j, " ") // Mencetak nilai j diikuti dengan spasi
        }

        fmt.Println() // Pindah ke baris baru setelah mencetak baris ke-1
    }

}

/*
0 1 2 3 4 
1 2 3 4 
2 3 4 
3 4 
4 

*/
// Slice (cap function)

//Slice (Cap function)Fungsi cap dapat kita gunakan untuk mengetahui kapasitasdari sebuah array maupun slice.Ketika pertama kali kita membuat suatu slice, panjang dankapasitasnya dipastikan sama, namun dapat berubah seiringdengan slicing yang kita lakukan.Contohnya seperti pada gambar disebelah kanan. Variablefruits1 telah disiapkan di awal dengan panjang 4 dankapasitas 4.Lalu terdapat 2 variable baru bernama fruits2 dan fruits3 yangmelakukan slicing terhadap fruits1 untuk mendapatkan 3element dari fruits1.Lalu kenapa fruits2 memiliki panjang dan kapasitas yangberbeda sedangkan fruits3 panjang dan kapasitasnya samatetapi kapasitasnya berkurang? Mari kita bedah di halamanselanjutnya.


package main 

import (
	"fmt"
	"strings"
)


func main() {
    // Membuat slice baru 'fruits1' dengan elemen "apple", "strawberry", dan "banana"
    var fruits1 = []string{"apple", "strawberry", "banana"}

    // Mencetak kapasitas dan panjang slice 'fruits1'
    fmt.Println("Fruits1 cap:", cap(fruits1)) // 4
    fmt.Println("Fruits1 len:", len(fruits1)) // 3

    // Mencetak pemisah untuk kejelasan
    fmt.Println(strings.Repeat("#", 20), "\n")

    // Membuat slice baru 'fruits2' dengan mengambil sebagian dari slice 'fruits1'
    var fruits2 = fruits1[0:3]

    // Mencetak kapasitas dan panjang slice 'fruits2'
    fmt.Println("Fruits2 cap:", cap(fruits2)) // 4
    fmt.Println("Fruits2 len:", len(fruits2)) // 3

    // Mencetak pemisah untuk kejelasan
    fmt.Println(strings.Repeat("#", 20), "\n")

    // Membuat slice baru 'fruits3' dengan mengambil sebagian lainnya dari slice 'fruits1'
    var fruits3 = fruits1[:1]

    // Mencetak kapasitas dan panjang slice 'fruits3'
    fmt.Println("Fruits3 cap:", cap(fruits3)) // 3
    fmt.Println("Fruits3 len:", len(fruits3)) // 1
}


/*

Penjelasan :
Slice (Cap function)Fungsi cap dapat kita gunakan untuk mengetahui kapasitasdari sebuah array maupun slice.Ketika pertama kali kita membuat suatu slice, panjang dankapasitasnya dipastikan sama, namun dapat berubah seiringdengan slicing yang kita lakukan.Contohnya seperti pada gambar disebelah kanan. Variablefruits1 telah disiapkan di awal dengan panjang 4 dankapasitas 4.Lalu terdapat 2 variable baru bernama fruits2 dan fruits3 yangmelakukan slicing terhadap fruits1 untuk mendapatkan 3element dari fruits1.Lalu kenapa fruits2 memiliki panjang dan kapasitas yangberbeda sedangkan fruits3 panjang dan kapasitasnya samatetapi kapasitasnya berkurang? Mari kita bedah di halamanselanjutnya.



Fruits1 cap: 3
Fruits1 len: 3
####################

Fruits2 cap: 3
Fruits2 len: 3
####################

Fruits3 cap: 3
Fruits3 len: 1

*/

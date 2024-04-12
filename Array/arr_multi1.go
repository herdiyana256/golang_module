/*
Array(Multidimensional Array)
 Kita juga dapat membuat suatu array multidimensional yangberarti terdapat sebuah array di dalam sebuah array. Contohnyaseperti pada gambar pertama di sebelah kanan.Array yang dikandung oleh variable balances merupakan sebuahmultidimensional array.Tanda [2][3]int  memiliki arti bahwa array terluar nya memilikipanjang sama dengan 2, dan array yang berada di dalamnyamemiliki panjang sama dengan 3 dengan tipe data int. {5,6,7}dan {8,9,10} merupakan array yang berada di dalam danmasing-masing memiliki panjang sama dengan 3 dengan tipedata int. Jika kita jalankan pada terminal kita, maka hasilnya akanseperti pada gambar kedua di sebelah kanan.

*/

package main // Mendeklarasikan bahwa file ini adalah bagian dari package "main", yang berarti ini adalah program yang dapat dieksekusi.

import "fmt" // Mengimpor package "fmt", yang berisi fungsi-fungsi untuk melakukan formatting input-output.

func main() { // Mendefinisikan fungsi utama dari program.

    balances := [2][3]int{{5, 6, 7}, {8, 9, 10}} // Mendeklarasikan dan menginisialisasi array dua dimensi dengan tipe data int, di mana setiap array dalam array utama memiliki panjang tiga.

    for _, arr := range balances { // Looping melalui setiap array dalam array utama "balances".
        for _, value := range arr { // Looping melalui setiap elemen dalam array yang sedang diproses.
            fmt.Printf("%d ", value) // Mencetak nilai elemen saat ini dalam format integer dengan spasi diikuti oleh newline.
        }
        fmt.Println() // Mencetak newline setelah mencetak semua elemen dalam array.
    }
}


/*

Output :
5 6 7 
8 9 10 




*/
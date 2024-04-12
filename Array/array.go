/*
Array - Sesi 2 Array Array pada bahasa Go merupakan sebuah tipe data untukmenyimpan sebuah kumpulan dari data-data. Data-data yang kitasimpan pada sebuah array dalam bahasa Go harus memiliki tipedata yang sama, kecuali kita menyimpannya sebagai suatuinterface kosong yang akan kita bahas nanti.Array pada bahasa Go memiliki sifat fixed-length atau memilikipanjang yang tetap yang harus kita tentukan dari awal kitamembuat arraynya.
Array
*/


/*

Array Kita dapat membuat array dengan 2 macam cara yaitu denganmendeklarasikannya terlebih dahulu tanpa memberi nilai apapun,dan kita juga bisa mendeklarasikan dan langsungmenginisialisasikannya dengan memberikannya sebuah nilai.Contohnya seperti pada gambar pertama di sebelah kanan. Kitatelah membuat 2 variable bernama numbers dan strings.Variable numbers memiliki tipe data array of int atau array dengantipe data int dan memiliki panjang 4. Lalu variable strings memilikitipe data array of strings atau array dengan tipe data string.Lalu kita memakai verb %#v untuk memformat tipe data array agarkita juga dapat melihat panjang dari arraynya. Jika kita jalankanmaka hasilnya akan seperti pada gambar kedua.


*/
package main // Ini adalah pernyataan package yang menunjukkan bahwa kode ini merupakan bagian dari package "main".

import "fmt" // Ini adalah pernyataan import yang mengimpor paket "fmt" untuk digunakan dalam kode.

func main () { // Ini adalah fungsi utama (main) dari program. Program akan dijalankan dari sini.
    var numbers [4]int // Mendefinisikan array numbers dengan panjang 4 yang berisi bilangan bulat.
    numbers = [4]int{1,2,3,4} // Menginisialisasi array numbers dengan nilai-nilai tertentu.

    var strings = [3]string{"Herdiyan", "Adam", "Putra"} // Mendefinisikan array strings dengan panjang 3 yang berisi string.
    
    fmt.Printf("%#v\n", numbers) // Mencetak array numbers menggunakan format %#v.
    fmt.Printf("%#v\n", strings) // Mencetak array strings menggunakan format %#v.
}


/*
[4]int{1, 2, 3, 4}
[3]string{"Herdiyan", "Adam", "Putra"}



*/
/*
Slice (make function)

Kita juga bisa membuat sebuah slice dengan menggunakanfungsi make. Argumen pertama yang diberikan pada gambarpertama di sebelah kanan adalah tipe dari slice nya, danargumen keduanya adalah panjang dari slice nya.Jika kita jalankan maka hasilnya akan seperti pada gambarkedua di sebelah kanan. Bisa dilihat pada gambar kedua, slicepada variable fruits belum berisi nilai apapun dan variablefruits memiliki tipe data slice of string atau slice dengan tipedata string.Maka dari ikut ketika di print ke terminal, maka hasilnya adalahtiga string kosong.


*/

package main // Ini adalah deklarasi package utama yang mendefinisikan bahwa file ini adalah bagian dari paket "main".

import "fmt" // Ini adalah deklarasi import yang mengimpor paket "fmt", yang digunakan untuk format input dan output.

func main() { // Ini adalah fungsi utama yang akan dijalankan ketika program dieksekusi.
    var fruits = make([]string, 3) // Ini adalah deklarasi variabel "fruits" yang merupakan slice (potongan) dari string dengan panjang awal 3. Fungsi make() digunakan untuk membuat slice baru dengan panjang dan kapasitas yang ditentukan.

    _ = fruits // Ini adalah penggunaan variabel sementara "_" untuk menunjukkan bahwa variabel "fruits" digunakan, meskipun tidak langsung dalam kode ini.

    fmt.Printf("%#v\n", fruits) // Ini adalah pemformatan dan pencetakan menggunakan fmt.Printf(). "%#v" adalah format specifier yang digunakan untuk mencetak representasi "fruits" dalam format yang terperinci. "\n" adalah karakter baris baru yang digunakan untuk membuat output menjadi baris baru.
}
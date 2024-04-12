// Menggabungkan slice 

package main


import "fmt"

func main () {

	mySlice1 := []int{1,2,3}
	mySlice2 := []int{4,5,6}

	mySlice := append(mySlice1, mySlice2...) // Gabungkan 2 slice 

	fmt.Println(mySlice)// Mencetak slice yang digabungkan 

	//[1 2 3 4 5 6]
}

/*
Kesimpulan:

Slice adalah struktur data yang penting dalam Golang. Memahami cara menggunakan slice dengan baik akan membantu Anda menulis kode yang lebih efisien dan mudah dipahami.

Sumber belajar:

Dokumentasi resmi Golang: https://go.dev/
Tutorial Golang: https://www.konsistensi.com/2014/03/mengatasi-angkettidak-valid.html
Buku Golang: https://www.golang-book.com/
Sumber


*/
// Mencari substring 

package main 

import (
	"fmt"
	"strings"
)

func main () {

message := "Selamat Belajar Golang!"
substring := "Golang"
found := strings.Contains(message, substring)
if found {

	/*
	esan output "Substring ditemukan!" muncul karena pada baris found := strings.Contains(message, substring) terjadi pencarian substring dalam message, dan karena substring ("Golang") ditemukan di dalam message ("Selamat Belajar Golang!"), nilai variabel found akan menjadi true.

Jadi, ketika eksekusi masuk ke blok if found, karena nilai found adalah true, maka pesan "Substring ditemukan!" dicetak.
	*/
	fmt.Println("Substring ditemukan!")
} else {
	fmt.Println("Substring tidak ditemukan!")
}
// Substring ditemukan!
}
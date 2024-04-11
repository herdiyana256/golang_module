package main 

import "fmt"

func main() {
	message := "Pemograman Golang"
	message1 := "Sangat Menyenangkan"
	firstChar := message[0]
	secondChar := message1[0]
	fmt.Println("Karakter pertama:", string(firstChar), string(secondChar))
}
//Karakter pertama: P S



// lebih ringkas nya 

package main 

import "fmt"

func main () {
	message := "Pemograman Golang"
	message1 := "Sangat Menyenangkan"
	fmt.Println("Karakter pertama:", string(message[0]), string(message1[0]))
}

//Karakter pertama: P S
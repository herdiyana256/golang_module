package main 

import "fmt"

func main () {
	message := "Aku sayang andi andriyani"
	for i := 0; i < len(message); i++ {
		fmt.Printf("%c ", message[i])
	}
}
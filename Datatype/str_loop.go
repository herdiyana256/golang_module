// Melakukan looping pada string 
package main 

import "fmt"

func main() {
	message:= "Golang Indonesia"
	for i := 0; i < len(message); i++ {
		fmt.Printf("%c ", message[i]) // Mencetak tiap karakter 
	}
}
// G o l a n g   I n d o n e s i a 
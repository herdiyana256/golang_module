/*
Switch 
Selain menggunakan keyword if, else if dan else, kita jugabisa menggunakan keyword switch untuk membuat suatukondisional, seperti contohnya pada gambar disebelahkanan.Jika kita perhatikan, kondisional pada gambar disebelahkanan akan menghasilkan kalimat “perfect” karena casepertama sudah terpenuhi. Jika tidak ada case yangterpenuhi maka kondisional tersebut akan menghasilkankalimat “not bad” karena keyword default sama dengankeyword else yang dimana akan dieksekusi blocknya jikatidak ada kondisi yang terpenuhi.Switch pada bahasa Go tidak memerlukan keyword break,jadi jika suatu case  telah terpenuhi maka dia tidak akanberlanjut kepada case berikutnya.

*/


package main 

import "fmt"

func main() {
	var score = 8 

	switch score {
	case 8:
		fmt.Println("Perfect") // Perfect
	case 7:
		fmt.Println("Awesome") // Awesome
	case 6:
		fmt.Println("Not Bad") // Not Bad
	default:
		fmt.Println("stupid") // stupid if value 7-8 not fulfilled
	}
}
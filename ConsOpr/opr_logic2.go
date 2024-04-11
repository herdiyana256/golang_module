/*
Operators (Logical Operators)Seperti yang kita lihat pada gambar disebelah kanan,terdapat sebuah variable bernama right dengan nilaitrue lalu terdapat sebuah variable bernama wrongdengan nilai false. Lalu terdapat variable dengan nama wrongAndRightyang dimana kita melakukan pengecekan kombinasitipe data bool antart variable wrong dengan right.Variable wrongAndRight akan menghasilkan nilai falsekarena jika kita memakai operator &&, maka keduanilainya harus sama dengan true.

*/

package main 

import "fmt"

func main() {
	var right = true
	var wrong = false 

	var wrongAndRight = wrong && right 
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)
	/*
Baris ini menggunakan operator logika && (AND) untuk menggabungkan nilai wrong dan right.

Operator && akan menghasilkan true hanya jika kedua nilai true.

Karena wrong bernilai false, maka wrongAndRight juga akan bernilai false.
	*/


	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)
	/*
Baris ini menggunakan operator logika || (OR) untuk menggabungkan nilai wrong dan right.

Operator || akan menghasilkan true jika salah satu atau kedua nilai true.

Karena right bernilai true, maka wrongOrRight akan bernilai true.
	*/

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)
/*
Baris ini menggunakan operator logika ! (NOT) untuk membalikkan nilai wrong.

Operator ! akan menghasilkan true jika nilai awal false, dan sebaliknya.

Karena wrong bernilai false, maka wrongReverse akan bernilai true.
*/

}

/*
Operators (Logical Operators)Kemudian terdapat sebuah variable dengan namawrongOrRight yang menghasilkan nilai true karena jikakita memakai operator || dan ingin menghasilkan nilaitrue, maka salah satu dari nilainya boleh false.Dan yang terakhir terdapat sebuah variable bernamawrongReverse yang menghasilkan nilai true, karenamemang jika menggunakan operator dengan simbol !,maka operator tersebut akan membalikan nilai daribool nya, misalkan nilainya true maka akan menjadifalse dan begitu juga sebaliknya.

Penjelasan : 
Baris ini menggunakan fungsi fmt.Printf untuk mencetak teks dan nilai variabel.
wrong && right adalah ekspresi yang akan dievaluasi terlebih dahulu.
\t digunakan untuk membuat tab pada output.
%t adalah format specifier untuk mencetak nilai boolean (true atau false).
\n digunakan untuk pindah baris pada output.



Output : 
wrong && right  (false) 
wrong || right  (true) 
!wrong          (true) 
*/
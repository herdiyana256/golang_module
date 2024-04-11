/*
Operators (Relational Operators)
Jika kita perhatikan pada gambar disebelah kanan, kita telahmembuat 4 variable yang dimana semuanya akanmenghasilkan tipe data bool yang dihasilkan daripengecekan kesamaan nilai menggunakan operatorperbandingan.Jika kita jalankan pada terminal kita, maka kita akan melihathasilnya seperti pada gambar kedua. Variable bernamafirstCondition akan menghasilkan nilai true karena memangangka 2 lebih kecil dari angka 3.

*/

package main

import "fmt"

func main () {
	var firstCondition bool = 2 < 3
	var secondCondition bool = "Herdiyan" == "Herdiyan"
	/*
Variable bernama secondCondition akan menghasilkan nilaifalse karena kata Herdiyan tidak sama dengan kata Herdiyan.Variable bernama thirdCondition akan menghasilkan nilaitrue karena angka 10 tidak sama dengan angka 2.3.Dan yang terakhir variable bernama fourthCondition akanmenghasilkan nilai true karena angka 11 kurang dari atausama dengan angka 11.

	*/
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11


	fmt.Println("first condition:", firstCondition)
	fmt.Println("second condition:", secondCondition)
	fmt.Println("third condition:", thirdCondition)
	fmt.Println("fourth condition:", fourthCondition)
}
/*
first condition: true
second condition: true
third condition: true
fourth condition: true
*/
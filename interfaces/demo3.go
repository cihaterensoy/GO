package interfaces

import "fmt"

//tip doğrulama işlemini göreceğiz
/*
func dogrula(i interface{}) {
	sayi := i.(int)
	fmt.Println(sayi)
}
*/
func dogrula(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo3() {
	//dogrula(5)

	var sayi1 interface{} = "Engin"
	dogrula(sayi1)
	//böyle bir şey gönderirsek interface'de tip hatası alacağız
	//bunun olmaması için hata ayıklama yapıyoruz
	var sayi2 interface{} = 5
	dogrula(sayi2)
	/*
		0 false
		5 true
		böyle bir geri dönüt alıyoruz 0 yazmasının sebebi int'in default değerini ataması
	*/
}

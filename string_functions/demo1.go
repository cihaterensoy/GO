package string_functions

import (
	"fmt"
	s "strings" //alias kodu olarak s kullanacağız
)

//bir stringi içinde istediğimiz ifade geçiyor mu vs bakabiliriz
//bize düzensiz olarak gelen formatları düzeltebiliriz vs vs

func Demo1() {
	var isim string = "Cihat"
	//string fonksiyonları strings paketi üzerinden gelir
	fmt.Println(s.Count(isim, "h"))
	//ismin içersinde kaç tane h var diye say
	//büyük kücük harf duyarlıdır
	fmt.Println(s.Contains(isim, "C"))
	//ismin içersinde C harfi var mı yok mu ?
	fmt.Println(s.Index(isim, "C"))
	//aradığımız elemanın görüldüğü ilk indexi döndürür
	//bulamazsa -1 döndürür
	//Tolower(), ToUpper() vardır

}

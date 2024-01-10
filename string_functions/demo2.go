package string_functions

import (
	"fmt"
	s "strings" //alias kodu olarak s kullanacağız
)

func Demo2() {
	isim := "Cihat"
	fmt.Println(s.HasPrefix(isim, "Cih"))
	//ifadenin ne ile başladığını kontrol edebiliriz
	//true ya da false döndürür
	fmt.Println(s.HasSuffix(isim, "at"))
	//aradığımız ifade ile bitiyor mu diye kontrol eder

	harfler := []string{"a", "u", "r", "o", "r", "a"}
	sonuc := s.Join(harfler, "*")
	//dizi temelli bir elemanları bir araya getirir, ayrac olan strin elemanların arasına yerleştilir

	fmt.Println(s.Replace(sonuc, "*", "+", 3))
	//string içersindeki bir şeyi başka bir şeyle değiştirmek
	//-1 tüm gördüğün şeyleri değiştir 3 verirsek 3 tane değiştirir vs vs

	fmt.Println(s.Split(sonuc, "*"))
	//belli bir formata göre verilmiş şeyleri ayırmaya yarar Joinin tam tersi
	//bunlar sonuçları bir yere atamaz

	fmt.Println(s.Repeat(sonuc, 5))
	//verdiğimiz string kopyasını yanyana yazar
}

package functions

//multi return veren fonksiyonları göreceğiz
//bütün program dillerinde bu yapı yoktur

// 4 işlemin sonucunu aynı anda döndüren bir fonksyion yazacağız
// birden fazla dönüş olduğunda bunları parantez içersinde belirtiyoruz
// hepsi aynı tipte olmak zorunda değil
func DortIslem(sayi1, sayi2 int) (int, int, int, float32) {
	toplam := sayi1 + sayi2
	cikarma := sayi1 - sayi2
	carpim := sayi1 * sayi2
	bolum := float32(sayi1 / sayi2) //çevirme işlemi böyle yapılmaktadır

	return toplam, cikarma, carpim, bolum
	//sonuc1,sonuc2,sonuc3,sonuc4 := functions.DortIslem(10,6) böyle bir ifade döndürebilir
	//istemediğimiz parametresi yerine _ koyabiliriz

}

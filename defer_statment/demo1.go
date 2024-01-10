package defer_statment

import "fmt"

//defer ifadesi bir fonksiyonun sonunda başka bir fonksiyonun çalışmasını istiyorsak onu defer ifadesiyle çağırabiliriz
//defer ifadesiyle onun kesin çalışacağını garanti ediyourz
//uygulama hata verse bile çalışacak

func A() {
	//bu fonksiyonu defer işlemi icn kullanacağız
	fmt.Println("a fonksiyonu çalıştı")

}

/*
	func B() {
		fmt.Println("b fonksiyonu çalıştı")
		//A()
		//normal yapıda böyle çağırabiliriz
		defer A()
		//deferin üstünde ne olursa olsun B fonksiyonun bitiminden sonra çalışıyor anlamı veriyor

}
*/

func B() {
	defer A()
	defer C()
	//üste koysak bile fonksiyonun bitiminde çalışıyor
	//birden fazla fonksiyonu defer olarak çağırabiliriz
	//ancak birden fazla defer varsa ilk başta birinciyi belleğe alıyor sırasıyla üst üste bunları alıp
	//sondan başa doğru yani last in first out mantığıyla çalışıyor
	//ilk önce C deferi çalışır

	fmt.Println("b fonksiyonu çalıştı")

}

func C() {
	fmt.Println("C fonksiyonu Çalıştı")

}

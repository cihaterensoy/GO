package defer_statment

import "fmt"

//defer fonksiyon bir hata verse bile çalışacaktır

func Demo2(sayi int32) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		return "Çift sayıdır"
	}
	if sayi%2 != 0 {
		return "Tek sayıdır"
	}
	//defer DeferFunc()
	return ""

}

//iflerden biri return'u gördüğü zaman fonksiyon okuması bitiyor
//ama defer bir fonksiyonu bizim fonksiyonumuzun altına koyarsak return olduğu icin onu okumadan fonksiyonu kapatabilir
//bu yüzden defer fonksiyonlar her zaman en üsttedir
//bu yukarıdaki fonksiyonda ilk önce bitti yazdırmasının nedeni bir şeyi return ettiriyoruz ilk başta return edip diğer fonksiyonu gönderiliyor
//eğer return yerine ilk önce fmt.Println ile "Cift sayıdır" yazdırsaydık başta o cıkacaktı

func Test() {
	sonuc := Demo2(9)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti")
}

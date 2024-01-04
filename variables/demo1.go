package variables

import "fmt"
//kütüphaneleri import ifadesiyle sisteme ekleriz


/*
kodların çalışması için c dili gibi bir   fonksiyonu olmalı


func main() {
    // fmt formattan gelir ekrana bişey yazdırmak gibi şeylerde kullanılır
    fmt.Println("Cihat Erensoy")
}
//go dilinde tanımladığımız bir şeyi kullanmazsak sıkıntılar cıkıyor

*/

func Demo1(){
	var metin string = "Aurora Aksnes" //stringler çift tırnakta kullanılıyor, tek tırnakta değil
	fmt.Println(metin)
	//tanımladığımız değişkenleri kullanmazsak hata alırız
	// int-> tamsayı icin kullanılır, float -> virgüllü sayılar icin kullanılır
	//Go dilinde değişkenler. var kdv int = 10 şeklinde tanımlayabiliriz
	//kdv3 := 25 şeklinde yazarsak veri tipini belirtmeden bir değişken tanımlayabiliriz
	kdv := 25
	fmt.Println(kdv)
	fmt.Printf("Veri Tipi: %T",kdv)//FORMATLI YAZDIRMA 
	//% virgülden sonraki ifadeyi yazdır T ise typetan geliyor

	//boolean kullanımı -> var durum bool =false veya true olaraktır
	//iki değişkenin eşitlik olup olmadığını true,false olarak bir değişkene atayabiliriz 
	 // == -> eşit mi?
	 // != -> eşit değil mi ?
	 // mantıksal bir değerin önüne ! koyarsak o mantıksal değerin tersini yazar

}
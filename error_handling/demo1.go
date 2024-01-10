package error_handling

import (
	"fmt"
	"os"
)

//hata ayıklamayı göreceğiz
//örneğimiz bir tane metin dosyasının acılması hatası olacak
/*
func Demo1() {
	//os paketi işletim sistemi paketdedir
	//os.Open("demo1.txt")
	//demo1.txt'sini aç demektir bize iki tane dönüş verir
	//birincisi  eğer dosya varsa dosyanın kendisini verir, yoksa hata veriyor
	f, err := os.Open("demo1.txt")
	//dosya bulunrsa err nil değerini alır
	if err != nil {
		//yan hata var anlamına geliyor
		fmt.Println("Dosya Bulunamadı")
	} else {
		fmt.Println(f.Name())
	}
	//eror_handling klasörünün içersinde bir demo1.txt'de olssa calıstırdığımız
	//yeri main.go olduğu icin dosyayı göremeyecek
	//bu yüzden hata verecek bu gibi durumların önüne geçmek icin hata yakalama kullanılır
	//demo1.txt kök dizine koyarsak yani main.go ile aynı yere koyarsam sıkıntısız bir şekilde çalışır

}
*/
//şimdi hatayı daha ayrıntılı görmek icin hamleler yapacağız
//mesela bir veritabanına bağlanıyoruz server kapalı olabilir, veritabanı şifresi vs yanlış olabilir
//pythonda try-catch ile yapabiliyoruz burada yöntem farklıdır
//buradaki adı Type Assertion
func Demo1() {
	f, err := os.Open("demo1.txt")

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			//dosyanın yolu ile bir hata varsa burası çalışır
			//*os.PathError bu şekilde yazmak bellek performansını arttırır
			//pathErrorsa ok değişkeni true olur ve hatamız pErr içersine atılır
			//ve buraya girilir
			//* adres üstünden gidiyoruz boşuna yenisini oluşturup ram'de yer tutmasın

			fmt.Println("Dosya bulunamadı", pErr.Path)
			return
		} else {
			fmt.Println("Başka bir hata var")
		}
		fmt.Println("Dosya Bulunamadı")
	} else {
		fmt.Println(f.Name())
	}

}

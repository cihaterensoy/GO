package goroutines

import (
	"fmt"
	"time"
)

//asenkron programlama yapacağız
//ilk etapda go routinelerini göreceğiz
//bir tane mobil uygulamanın yüklendiğini düşünelim
//bu yüklenme devam ederken biz aynı anda başka işler yapabiliriz
//işletim sistemi bir program olduğu icin bizim aynı anda hem bişey indirebilmemiz hem de diğer işleri yapabilmemize asenkron programlama denir

func CiftSayilar() {
	for i := 0; i < 10; i += 2 {
		//10 a kadaki cift sayıları yazacak
		fmt.Println("Cift sayı ", i)
		time.Sleep(1 * time.Second)
	}

}
func TekSayilar() {

	for i := 1; i < 10; i += 2 {
		//10 a kadaki tek sayıları yazacak
		fmt.Println("Tek sayı ", i)
		time.Sleep(1 * time.Second)
	}

}

/*
func CiftSayilar() {
	for i := 0; i < 10; i += 2 {
		//10 a kadaki cift sayıları yazacak
		fmt.Println("Cift sayı ", i)

	}

}
func TekSayilar() {

	for i := 1; i < 10; i += 2 {
		//10 a kadaki tek sayıları yazacak
		fmt.Println("Tek sayı ", i)

	}

}
goroutines.CiftSayilar()
goroutines.TekSayilar()
Cift sayı  0
Cift sayı  2
Cift sayı  4
Cift sayı  6
Cift sayı  8
Tek sayı  1
Tek sayı  3
Tek sayı  5
Tek sayı  7
Tek sayı  9

olarak bir cıktı aldık, kodlar adım adım okunuyor bu yüzden ilk başta çift sayıları yazdırır işi bitince diğer koda geçer
işlemler sıralı bir şekilde çalışır

eğer biz iki işleminde paralel olarak aynı anda çalışmasını istiyorsak ne yapmamız gerekir ?
bunun icin go routinlerinden yararlanırız
go goroutines.CiftSayilar()
go goroutines.TekSayilar()
şeklinde go routinlerin şeklinde kullanabiliriz.
main fonksiyonunda bunları çağrıcında go diline göre bunları farklı bir "thread"e alıyor ve kodları okumaya devam ediyor o yüzden direkt böyle yazınca bir cıktı alamıyoruz

bu olayın ne olduğunu görmek amacıyla time paketinden yararlanacağız

time.Sleep(5 * time.Second)
bunu yazarsak main 5 saniye boyunca duracak ve sonrasında kodları okumaya devam edecek
bu sayede diğer threadlerin çalışmasına fırsat veriyoruz

//----
go goroutines.CiftSayilar()
	go goroutines.TekSayilar()
	time.Sleep(5 * time.Second)
	fmt.Println("5 saniye geçti")
Tek sayı  1
Tek sayı  3
Tek sayı  5
Tek sayı  7
Tek sayı  9
Cift sayı  0
Cift sayı  2
Cift sayı  4
Cift sayı  6
Cift sayı  8
5 saniye geçti

böyle yazılmasının sebebi, burada yapılan işlemler o kadar hızlı ki go dili bir tanesi için threadi acarken diğeri işlemi bitirmiş oluyor
bu yüzden bu tip bir saydırma işlemi icin cok hızlı

bunu engellemk icin fonksiyonlarımızın icinede time.sleep kullanacağız


func CiftSayilar() {
	for i := 0; i < 10; i += 2 {
		//10 a kadaki cift sayıları yazacak
		fmt.Println("Cift sayı ", i)
		time.Sleep(1*time.Second)
	}

}
func TekSayilar() {

	for i := 1; i < 10; i += 2 {
		//10 a kadaki tek sayıları yazacak
		fmt.Println("Tek sayı ", i)
		time.Sleep(1*time.Second)
	}

}

Tek sayı  1
Cift sayı  0
Tek sayı  3
Cift sayı  2
Cift sayı  4
Tek sayı  5
Tek sayı  7
Cift sayı  6
Cift sayı  8
Tek sayı  9
5 saniye geçti
bunların sayesinde go routinlerin birlikte calıştığını görmek mümkün



*/

package main

//paketlemeyle kütüphaneler oluşturabiliriz, birbiriyle ilişkili olan dosyaları bu paketlerin içersine koyarız
//bir paketi başka bir yerde kullanmak icin import kullanmalıyız

//modüle oluşturabilmek icin  "go mod init golesson" yazarız terminalimize
//modülleri kütüphanelere benzetebiliriz
//paketleri sınıflara benzetebiliriz

import (
	"fmt"
	//"golesson/variables"
	//"golesson/conditionals"
	//"golesson/loops"
	//"golesson/arrays"
	//"golesson/slices"
	//"golesson/functions"
	//"golesson/maps"
	//"golesson/for_range"
	//"golesson/pointer"
	//"golesson/structs"
	//"golesson/goroutines"
	//"time"
	//"golesson/gochannels"
	//"golesson/interfaces"
	//"golesson/defer_statment"
	//"golesson/error_handling"
	//"golesson/interfaces"
	//"golesson/error_handling"
	//"golesson/string_functions"
	//"golesson/restful"
	"golesson/project"
) //birden fazla import icin bu yapıyı kullanabiliyoruz

func main() {
	//variables.Demo1()
	fmt.Println()
	//conditionals.Demo1()
	//conditionals.Workshop1()
	//------
	//loops.Demo1()
	//loops.Workshop3()
	//-----
	//slices.Demo2()
	//-----
	//functions.SelamVer()//class yapısındaki gibi bir fonksiyonu cağırabiliyoruz
	//functions.Topla(2,3)
	//var sonuc = functions.Toplar(2,6)
	//fmt.Println(sonuc*2)
	/*
		sonuc1,sonuc2,sonuc3,sonuc4 := functions.DortIslem(10,6)
		fmt.Println(sonuc1)
		fmt.Println(sonuc2)
		fmt.Println(sonuc3)
		fmt.Println(float32(sonuc4))
		sonuc1,sonuc2,sonuc3,_ := functions.DortIslem(10,6)
		4.parametreye ihtiyacım yok diyoruz _ ile
	*/
	/*
		sonuc1 := functions.ToplaVariadic(2,3,4,5,6,7,7,6,5,4)
		fmt.Println(sonuc1)
		sayilarim := []int {4,5,6,7,5,4,3}
		fmt.Println(functions.ToplaVariadic(sayilarim...))
	*/
	//maps.Demo1()
	//for_range.Demo3()
	//----
	/*
		var sayi int = 20
		pointer.Demo1(&sayi)
		fmt.Println("Maindeki sayı : ",sayi)
	*/
	// sayilar:= []int{1,2,3}
	// pointer.Demo2(sayilar)
	// fmt.Println("Maindeki Sayı: ",sayilar[0])
	//-----
	//structs.Demo2()
	//-----
	/*
		go goroutines.CiftSayilar()
		go goroutines.TekSayilar()
		time.Sleep(5 * time.Second)
		fmt.Println("5 saniye geçti")
	*/
	//--------
	/*
		//burada bir değişken tanımlayıp fonksiyonlarımın icersindeki chanella erişmem gerekiyor
		ciftSayiCn := make(chan int)
		tekSayiCn := make(chan int) //bu sayede fonksiyonların içersindeki channeları atayabiliyoruz aynı dizilerdeki gibi
		go gochannels.CiftSayilar(ciftSayiCn)
		go gochannels.TekSayilar(tekSayiCn)
		//var carpim int = ciftSayiCn * tekSayiCn
		//buradaki değerler bizim icin asla ama asla bir sayı değil sadece sayı taşıyan bir channel struct gibi düşünebiliriz
		//bunları okuyabilmemiz için yapmamız gereken işlem

		ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
		var carpim int = ciftSayiToplam * tekSayiToplam
		//chanellın değerini okuyup ciftSayiCn'yi ciftSayiToplama aktarıyoruz
		fmt.Println("Çarpım : ", carpim)
	*/
	//interfaces.Demo2()
	//--------------
	//defer_statment.Demo3()
	//--------
	//error_handling.Demo1()
	//--------
	//fmt.Println(error_handling.TahminEt2(102))
	//----
	//string_functions.Demo2()
	//-----
	//restful.Demo2()
	//-------
	//project.AddProduct()
	products, _ := project.GetAllProducts()
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}

}

//fonksiyon isimleri PascalCase olmalıdır

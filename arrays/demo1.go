package arrays

import "fmt"

//araylari neden kullanmak isteriz ?
//aynı tipteki veri tiplerini bir arada tutmak icin kullanılır

func Demo1(){
	var sayilar [5]int
	//5 tane 0 elemanı tanımladı cünkü biz bir değer vermedik 
	sayilar[2] = 50//set işlemi gercekleştirdik
	//ücüncü elemana 50 verdik
	fmt.Println(sayilar)
	fmt.Println(sayilar[2])
}
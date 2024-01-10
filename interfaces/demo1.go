package interfaces

import (
	"fmt"
)

//interfaceler birer şablon olarak düşünebiliriz
//interfaceler bize metod imzaları getirir
//örneğin bir sürü alan hesaplma formülü var ama yöntemleri farklı
//bu noktada bizde interface örneği yapacağız

type shape interface {
	area() float64
	//birden fazla şey yazılabilir
}

//area'yı kullanacak structlarımızı yazacağız

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

//structlar icin metod yazacağız
//burada dikkat edeceğimiz şey metodların adını shape içersindeki yani interface içersindeki değişkenle aynı olmalıdır

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return (3.14) * c.radius * c.radius

}

func school(s shape) {
	//s shape /*shape şablonuna uyan bir girdi ver demektir*/
	//burada parametre olarak interface'i istiyoruz
	fmt.Println("Şeklin alanı : ", s.area())

}

func Demo1() {
	r := rectangle{width: 10, height: 6}
	school(r) //shape'in imzasına uyan bir tane struct yollamamız gerkiyor
	//rectangle'ın bellekteki yerini gönderiyoruz

	c := circle{radius: 10}
	school(c)

}

//mesela dairenin fazladan bir metodu olabilir ama interface'in icinde bulunanlardan biri eksik olamaz
//interfaceler programların sürdürülebilirliğiyle doğru orantılıdır

//yukarıdaki structlara birer tane metod ekliyoruz bunlar area adında metodlar bunları icermesi gerektiğini düşündüğümüz için
//go programlama dilinde interface'i uyguladığımız anlatmanın tek yolu metoduda aynı isimde vermektir
//interface'imizin birden fazla metodu olabilir structlarımızın bunları icermesini bekliyoruz ancak structlarımızın her metodunu interface içermeyebilir

//school fonksiyonunda shape interface'ini istiyoruz, pointerlarla otomatik çalışıyor burası

//avantajı şudur sonradan bir şekil eklersek mevcut kodlara dokunmadan özelliklerini ekleriz
//çok bicimliliği interfaceler ile yaparız, go gibi dillerde

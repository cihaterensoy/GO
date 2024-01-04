package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

//bunun icin yazdığımız funclar operasyon denir
/*
func save() {
	fmt.Println("Eklendi")
} //bu eklendiyi customer'ın altında çağırmak istiyoruz diyelim
*/
//struct fonksiyonu oluşturmak icin yapmamız gereken
func (c customer) save() {
	fmt.Println("Eklendi: ", c.firstName)
	//bu fonksiyon save adındadır. c adında struct bir nesne göndereceğim adındadır
	//metod oluşturduk kısacası
}

func Demo2() {
	n := customer{firstName: "Cihat", lastName: "Erensoy", age: 23}
	//n. nokta diyince sadece değişkenlerimizi görüyoruz
	n.save()
	//ancak eğer fonksiyonu struct fonksiyonuna dönüştürürsek o fonksiyonuda n. yapınca görebiliriz

}

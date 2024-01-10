package interfaces

import "fmt"

//kredi alma uygulaması yapacağız
//kredi alırken ödeme yüzdeleri değişebilir mortgage yüzde 10 olsun car'da yüzde 15 olsun

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

// mortgage ve car icin creditCalculator'u implemente edeceğiz

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
	//hesaplamalar tamamen uydurma
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

//aşağıdaki fonksiyonda birden cok kredimiz olabileceği için devreye çok biçimlilik girecektir
//bundan önceki örnekte 1 tane shape göndermiştik burada birden fazla kredisi olabileceği icin dizi gönderiyoruz

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	var paymentTotal float64 = 0.0 // Corrected the data type to float64
	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal
}

//bir dizi var elimizde o dizide tek tek dolaşıyoruz
//i0da 2 kredisi varsa her biri için paymentTotali o anki kredinin calculati ile topluyoruz

func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, address: "Norveç Oslo"}
	credit2 := Mortgage{rate: 12, creditPaymentTotal: 50000, address: "Norveç Bergen"}
	credit3 := Car{rate: 15, creditPaymentTotal: 60000, carInfo: "Tesla CyberTruck"}
	credits := []CreditCalculator{credit1, credit2, credit3} //creditCalculator temelli bir yapı yollar
	total := CalculateMonthlyPayment(credits)                //parametre olarak bir kredi arrayi yollayacağım
	//daha önce string arrayi felan yaptık ama interfaces arrayide yapabiliriz

	fmt.Println("Toplam ödeme: ", total)

}

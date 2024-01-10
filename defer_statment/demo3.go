package defer_statment

import "fmt"

//defer ifadeleri structların metodları içinde kullanıyoruz

type product struct {
	productName string
	unitPrice   int
}

// structın metodu
func (p product) save() {
	fmt.Println("Ürün Kaydedildi", p.productName)
	defer Log()
	//böyle de loglama ifadeleri kullanabiliriz

}

func Log() {
	fmt.Println("loglandı")
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 5000}
	defer p.save()
	p = product{productName: "Mouse", unitPrice: 4000}
	//defer 2.sırada atadığımızda ilk pnin değerini alırdı eğer en alttaki
	//bir yerde olsaydı defer bize mouse'u döndürürdü
	//deferin nereye koyulduğu önemlidir bu yüzden
	fmt.Println("İşlem Başarılı")
	fmt.Println(p.productName)
}

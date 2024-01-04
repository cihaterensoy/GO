package structs

import "fmt"

//pythondaki claslara benzetebiliriz
//değer güdümlü çalışır yani int vs gibi değişkentipleriyle aynı mantıkta calısır

func Demo1() {
	fmt.Println(product{"Bilgisayar", 500, "XYZ", 0}) //PRODUCTtan bir nesne oluşturudk
	//structaki bütün değerleri eksiksiz yazmamızı istiyor
	//belirtmeyip diğer alanları defualt olarak bırakmak istiyorsak parametreleri isimli vermemiz gerekiyor

}

type product struct {
	//product olarak bir structımız var
	//bunun icine örneğin bir ürün özelliğini koyacağız
	name         string
	unitPrice    float64
	brand        string
	discountRate int
}

//product structımızda özellikleri koyduk ama çeşitli operasyonlarada ihtiyac duyabiliriz
//ekleme,cıkarma,silme vs vs
//structların icersine func oluşturamıyoruz
//demo2.go'da buna bakacağım

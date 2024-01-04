package pointer
import "fmt"
//ilk önce bir fonksiyon oluşturup o fonksiyona bir parametre yollayacağız


func Demo1(sayi *int){
	*sayi = *sayi + 1
	fmt.Println("Demodak sayı : ",*sayi)
	
}

/*
Demodak sayı :  21
Maindeki sayı :  20
*/
//biz mainden fonksiyona sayıyı göndersekte maindeki sayı değişkenimiz etkilenmez
//bunun nedeni değişkenlerimiz değeri tutan veri tipleridir, buradaki sayi değişkeni bellekte yepyeni bir değer olarak oluşturuluyor
//Demoya sadece değeri aktarıyoruz ramdeki yerini değil, sadece kopyalıyoruz
//sayının adresini göndermiyoruz yani

//ancak bu fonksiyonun yaptığı hesap maindeki değişkeni etkilesin istiyorsak
//o zaman pointerlerdan yararlanıyoruz
//ramdeki yerini işaret edenlerden yararlanıyoruz yani
//*int diye tanımlama yaparsak değerden öte artık onun bellekteki adresiyle çalışacağımızı belirtiriz
/*
func Demo1(sayi *int){
	*sayi = *sayi + 1
	fmt.Println("Demodak sayı : ",sayi)
	
}
kullanım böyledir
*sayi = *sayi + 1
adres olan intle değer olan int'i karıştırmaması icin ikisinde de * kullanıyoruz


Mainde böyle bir kod yazmamız gerkeiyor
var sayi int = 20
	pointer.Demo1(&sayi)
	fmt.Println("Maindeki sayı : ",sayi)

	çıktısı böyledir

Demodak sayı :  0x14000110020 //değişkenin bellekteki adresidir

Maindeki sayı :  21

&sayi  bunun anlamı sayının adresiyle beraber sayiyi yolladığımızı anlatıyoruz

fmt.Println("Demodak sayı : ",*sayi) yaparsak bellekteki konumu değil bellekteki değeri yazdırır

yani sayısal veriler değerleriyle kabul görürler, pointersız gönderirsek sadece değerini göndeririz adresimizle işimiz olmaz

**arrayler ve array bazlı yapılar değerleri üzerinden çalışmaz, bellekteki adresleriyle gönderilir ve onlarla çalışıyor

*/

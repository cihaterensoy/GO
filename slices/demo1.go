package slices
import "fmt"

func Demo1(){
	//değişken adı:= make(type,0)
	//tipini ve kaç elemandan oluşacağını yazıyoruz
	isimler:= make([]string,3)//slice tanımlama

	fmt.Println(isimler)
	//[  ] yazdırınca cıkan budur
	isimler[0] = "Cihat"
	isimler[1] = "Aurora"
	isimler[2] = "Taylor"

	fmt.Println(isimler)
	//isimler[3] = "Hüseyin" yeni elemanı böyle ekleyemiyoruz append fonksiyonu için ekliyoruz
	isimler = append(isimler,"Hüseyin")
	//bu şekilde yeni elemanımızı ekliyoruz
	//appendle yepyeni bir array oluşturuluyor olarak düşünebiliriz

	fmt.Println(isimler)
}
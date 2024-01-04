package arrays
import "fmt"

func Demo2(){
	var sehirler [5]string
	sehirler[0] = "Ankara"
	sehirler[1] = "İstanbul"
	sehirler[2] = "Karabük"
	sehirler[3] = "Trabzon"
	sehirler[4] = "Zonguldak"

	fmt.Println(sehirler)
	//dizi boyutunun dışına cıkamıyoruz python gibi kendi kendine ayarlamıyor
	for i:=0;i<len(sehirler);i++{
		fmt.Println(sehirler[i])
	}
}
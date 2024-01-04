package slices
import "fmt"

func Demo2(){
	//bir slice'ı kopyalamak icin yapmak gerekenler
	sehirler:=[]string{"Ankara","İstanbul","Karabük"}
	//bu şekilde de bir slice oluşturabiliriz
	fmt.Println(sehirler)
	sehirlerKopya := make([]string,len(sehirler))//şehirlerin boyunda bellekte bir yer ayırdık

	copy(sehirlerKopya,sehirler)//bu şekilde şehirleKopya slice'nın içine şehirleri kopyaladım
	//bunlar artık iki ayrı arraydir ve bellekteki refransları farklıdır aşağıdaki örnekteki gibi 
	fmt.Println(sehirlerKopya)

	sehirler = append(sehirler,"Adana")//ana slice değişse bile kopya olan sonradan değişmiyor
	fmt.Println(sehirler)
	fmt.Println(sehirler[2:3])//şehirler slice'nın 2.indisinden 3.indisine kadar al ama 3.dahil değildir
	fmt.Println(sehirler[2:])//2.indisten sonuna kadar al
	fmt.Println(sehirler[:3])//baştan 3.indise kadar alır


}
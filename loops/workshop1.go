package loops
import "fmt"


func Workshop1(){
	aklimdakiSayi := 80
	var tahminEdilenSayi int
	var sayac int = 0 
	for tahminEdilenSayi !=aklimdakiSayi {
			fmt.Println("Tahmin ettiğiniz sayıyı girin:")
			fmt.Scanf("%d",&tahminEdilenSayi)	
			//kullanıcıdan aldığımız değeri değişkene atamak icin & sembolünü kullanıyoruz
		if aklimdakiSayi>tahminEdilenSayi{
			fmt.Println("Tahmin ettiğiniz sayı aklımdaki sayıdan küçük:")
			sayac = sayac + 1

		}else if aklimdakiSayi<tahminEdilenSayi{
			fmt.Println("Tahmin ettiğiniz sayı aklımdaki sayıdan büyük:")
			sayac = sayac + 1
		}
	}
	fmt.Println("Aklımdaki Sayıyı Buldun")
	if sayac <= 3 {
		fmt.Println("Süper")
	}else if sayac >=4 && sayac <6 {
		fmt.Println("Güzel")
	}else if sayac >6 {
		fmt.Println("fena Değil")
	} 
}
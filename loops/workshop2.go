package loops
//kullanıcıdan bir sayı girmesini isteyeceğiz
//örneğin 23 girecek sayı asal mı değil mi bulacağız
import "fmt"
func Workshop2(){
	var kullaniciSayi int
	fmt.Println("Lütfen Bir Sayı Giriniz")
	fmt.Scanf("%d",&kullaniciSayi)

	asalMi:= true
	for i:=2;i<kullaniciSayi;i++{
		if kullaniciSayi % i == 0{
			asalMi= false
		}
	}
	if asalMi == true{
		fmt.Println("Asaldır")
	}else{
		fmt.Println("Asal Değildir")
	}
}
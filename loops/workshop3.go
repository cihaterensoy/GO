package loops

//arkadaş sayılar sorusu
//eğer iki sayı birbirinin kendisi haric bölenlerinin toplamı karşısındakine eşitsa bu iki sayı arkadaş sayıdır
//en kücük iki arkadaş sayıya örnek 220 ve 284
import "fmt"
func toplam(sayi int) int {
	var toplam int = 0
	for i := 1; i < sayi; i++ {
		if sayi%i == 0 {
			toplam = toplam + i
		}
	}
	return toplam
}

func Workshop3(){
	var sayi1,sayi2 int
	
	fmt.Println("denemek istediğiniz birinci sayıyı giriniz")
	fmt.Scanf("%d",&sayi1)
	fmt.Println("denemek istediğiniz ikinci sayıyı giriniz")
	fmt.Scanf("%d",&sayi2)

	if toplam(sayi1) ==sayi2 && sayi1 == toplam(sayi2){
		fmt.Printf("%d ve %d, arkadaş sayılardır.", sayi1, sayi2)
	} else {
		fmt.Printf("%d ve %d, arkadaş sayılar değildir.", sayi1, sayi2)
	}
}


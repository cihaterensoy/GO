package functions

//bazen fonksiyonlarda aynı tipte olan ama sayısı belli olmayan parametrelere ihtiyac duyabiliriz
//bunlara VARİADİC  fonksiyonlar denir
func ToplaVariadic(sayilar ...int) int{
	//bizim gönderdiğimiz parametreleri bir dizi olarak ele alır
	//yazım şekli aynen böyledir sayilar ...int
	//dizi ile ilgili ne yapabiliyorsak burada da onu yapabiliriz
	var toplam int = 0
	for i:=0; i<len(sayilar); i++{
		toplam = toplam + sayilar[i]
	}
	return toplam
	//bu fonksiyona bir arrayi parametre olarak verebiliriz
	//sayilarim := []int {4,5,6,7,5,4,3}
	//fmt.Println(functions.ToplaVariadic(sayilarim...))

}
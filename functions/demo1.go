package functions
import "fmt"
//bir kodu defalarca yazmamak icin kullandığımız yapıdır
func Topla(sayi1 int,sayi2 int){
	//parametreli bir fonksiyon oluşumudur
	fmt.Println(sayi1+sayi2)
	//bu sadece işlem yapıyor geri döndürmüyor

}
func Toplar(sayi1 int,sayi2 int) int{
	//bir sonuc geri döndüren fonksiyon tipidir
	var Toplam int 
	Toplam = sayi1 + sayi2
	return Toplam
}
func SelamVer(){
	fmt.Println("Merhaba")
}
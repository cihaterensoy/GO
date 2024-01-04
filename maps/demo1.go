package maps
import "fmt"
func Demo1(){
	//dictionary ve hash gibi yapıların karşılığıdır
	//anahtar değer gibi bir yapıdır yani

	sozluk := make(map[string]string)//sozluk tanımlanması bu şekildedir anahtar string- key string
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman sayısı. ",len(sozluk))
	delete(sozluk,"book")
	fmt.Println("Eleman sayısı. ",len(sozluk))
	fmt.Println("Sözlük: ",sozluk)

	//bir değerin sözlüğün icersinde var mı yok mu nasıl ararız?
	deger, varMi := sozluk["table"]
	//bu yapı bize 2 değer döndürüyor bi değerin kendi
	//ikinci olarakta değerin olup olmadığı bir boolean değer

	fmt.Println(deger)
	fmt.Println("Listede olma durumu: ",varMi)

	//sözlükteki elemanları bir for döngüsüyle döndürebiliyoruz
	//sözlüklerin başka bir oluşturulma şekli daha vardır
	//sozluk2:=map[string]string{"glass":"bardak"}

}
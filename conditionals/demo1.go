package conditionals

import "fmt"
func Demo1(){
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
        fmt.Print("Hesaptaki bakiye yetersiz")
    } else {
        fmt.Print("Paranız Veriliyor")
		hesap = hesap - cekilmekIstenen
    }
	fmt.Println(" Bitti. Hesaptaki para: " + fmt.Sprintf("%v",hesap))
	//fmt.Sprintf("%f"hesap)) formatlı yazdırıyor f dersek float şeklinde yazdırır, v dersek direkt value'yu yazdırır
}
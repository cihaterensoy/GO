package arrays

import "fmt"

func Demo3() {
	var sayilar [5]int = [5]int{70, 30, 40, 8000, 60}
	fmt.Println(sayilar)
	var enBuyukSayi int = sayilar[0]

	for i := 0; i < len(sayilar); i++ {
		if sayilar[i]>enBuyukSayi{
			enBuyukSayi = sayilar[i]
		}
	}
	fmt.Printf("En büyük sayı: %d", enBuyukSayi)
}

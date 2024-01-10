package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklımdakiSayi := 50
	//error go programlama dili tarafından tanınan bir kelimedir
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı griiniz.")
		//errors adında bir paket vardır
	}
	if tahmin > aklımdakiSayi {
		return "Daha kücük bir sayı giriniz", nil
	}
	if tahmin < aklımdakiSayi {
		return "Daha büyük bir sayı giriniz", nil
	}
	return "Bildiniz", nil
	//nil dediğimiz için hata yok döndürüyoruz

}

func Demo2() {
	mesaj, _ := TahminEt(50)
	fmt.Println(mesaj)
	fmt.Println(TahminEt(51010))
	fmt.Println(TahminEt(-10))

}

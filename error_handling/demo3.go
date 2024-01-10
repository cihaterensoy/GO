package error_handling

import (
	"fmt"
)

//kendi hata tiplerimizi oluşturabiliriz

// mesela sayı girilme sınırı olsun
type borderException struct {
	parameter int
	message   string
}

//go dilinde error diye bir interface var onun icersinde de error diye bir metod vardır

func (b *borderException) Error() string {
	return fmt.Sprintf("%d %s", b.parameter, b.message)
}

func TahminEt2(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &borderException{tahmin, "Sınırların dışında kaldı"}
		//burada hata olarka direkt borderException structını döndük
	}
	return "Bildiniz", nil

}

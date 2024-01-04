package for_range
import "fmt"

func Demo3() {
	sozluk := map[string]string{"book":"kitap","table":"masa"}

	for _,value:=range sozluk{
		fmt.Println(value)
	}
}
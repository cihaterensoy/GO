package for_range
import "fmt"

//1-10 arasındaki sayılardan tek olanları toplayan eleman
func Demo2(){
	sayilar:=[]int{1,2,3,4,5,6,7,8,9,10}
	var toplam int
	for _,sayi := range sayilar{
		toplam = sayi + toplam
		
	}
	fmt.Printf("Sayilar %v",toplam)
	
}
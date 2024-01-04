package conditionals
import "fmt"

func Workshop1(){
	
	var i,j,k int = 1,2,3
	/*
	if i > j && i > k {
        fmt.Printf("i en büyük değer %v", i)
    } else if j > k && j > i {
        fmt.Printf("j en büyük değer %v", j)
    } else if k > j && k > i {
        fmt.Printf("k en büyük değer %v", k)
    } else if k == j && j == i {
        fmt.Println("Bütün değerler eşit")
    }
	*/
	
	var enbuyukDeger int = i 
	if j > enbuyukDeger{
		enbuyukDeger = j
	}
	if k > enbuyukDeger{
		enbuyukDeger = k
	}
	fmt.Printf("En büyük değer %v ",enbuyukDeger)
}
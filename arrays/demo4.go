package arrays

import "fmt"
func Demo4(){
	//çok boyutlu dizi tanımlama

	var sayilar1 [2][3]int
	//birincisi satır sayısı 2. sütun sayısıdır

	sayilar1[0][0] = 1;
	sayilar1[0][1] = 3;
	sayilar1[0][2] = 5;
	sayilar1[1][0] = 2;
	sayilar1[1][1] = 4;
	sayilar1[1][2] = 6;


	for i:=0;i<len(sayilar1);i++{
		for j:=0;j<len(sayilar1[i]);j++{
			fmt.Println(sayilar1[i][j])
		}
	}
	//dizilerdeki temel problem yeni eleman eklemek ve dizinin eleman sayısını dinamik olarak attırmaktır
	//go dilinde altında array calısan ama kendisi yepyeni bir yapı olan slice kullanılacak
	//yeni bir şey eklenince yeniden boyutlandırmadan kurtuluruz
	

}
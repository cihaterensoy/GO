package for_range

//for döngüsünde range yapısını kullanmak
//range yapısıyla dizi temelli yapılarda elemanları tek tek dolaşabiliyoruz
//map yapılarında range kullanmak daha mantıklıdır

import "fmt"

func Demo1(){
	sehirler:=[]string{"Ankara","İstanbul","İzmir"}

	for i:=0;i<len(sehirler);i++{
		fmt.Println(sehirler[i])
	}

	for i,sehir:= range sehirler{
		fmt.Println(sehir)//dizideki o andaki elemanımız
		fmt.Println(i)//dolaşırken kaçıncı indeksteyiz kullanmayacaksak onun yerine _ koyabiliriz
	}
	//pythondaki gibi düşünebiliriz

}
package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"` //restful yapısına benzetmek icin yaptık ikinci userId API'dekine karşılık gelmektedir
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"` //bu syntax apideki neyi maplediğimizi gösterir
}

// json to go yazarsak apıdeki yazılara go formatına uygun bir şekilde dönüştürür
func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	//yukarıdaki kod siteye okuma isteğinde bulunucak
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close() //hatada olsa response body'i kapat demektir

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	//hata ve bize bilgiler döndürüyor byte olarak

	bodyString := string(bodyBytes) //elimizde bir string var bunu json yapısına dönüştüreceğiz
	fmt.Println(bodyString)
	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)
}

/*
{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}
{1 1 delectus aut autem false}
biz genelde 2. secenek olarak okuruz
s
*/

func Demo2() {
	//post operasyonu bir datayı yollama üzerine kuruludur
	todo := Todo{UserId: 1, Id: 2, Title: "Alışveriş yapılacak", Completed: false}

	jsonTodo, err := json.Marshal(todo) // todo'yu jsona dönüştürdük, err değişkeni döndürüyor
	if err != nil {
		fmt.Println(err)
		return // Hata durumunda işlemi sonlandır
	}

	// Gönderilecek olan datayı ve content type'ı belirterek HTTP POST isteği yapılıyor
	resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))
	if err != nil {
		fmt.Println(err)
		return // Hata durumunda işlemi sonlandır
	}
	defer resp.Body.Close() // İşlem tamamlandığında yanıtın kapatılması için defer kullanılıyor

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return // Hata durumunda işlemi sonlandır
	}

	// Hata ve bize bilgiler döndürüyor byte olarak
	bodyString := string(bodyBytes) // Elimizde bir string var bunu json yapısına dönüştüreceğiz
	fmt.Println(bodyString)

	var todo2 Todo
	err = json.Unmarshal(bodyBytes, &todo2)
	if err != nil {
		fmt.Println(err)
		return // Hata durumunda işlemi sonlandır
	}

	fmt.Println(todo2)
}

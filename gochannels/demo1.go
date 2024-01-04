package gochannels

/*
go rountilerinde time.sleep kullanarak senkronize çalışmalarını sağladık
buradaki problem şudur;
işlemci cok hızlı işlemler yaptığı icin bizim kafamıza göre verdiğimiz sürenin yeterli olduğunu düşünüyoruz
ama uzak bir sunucuda olsaydı kafamıza göre bir süre veremezdik
time.sleep yerine farklı bir cözüm bulmamız gerekiyor
time.sleep koymasakta bu işlemler asenkron çalışıyor biz sadece görebilmek icin yapıyoruz

*/

func CiftSayilar(ciftSayiCn chan int) {
	var toplam int = 0
	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i

	}
	//işlemin en sonuna bunu yapıyoruz
	//bu aşağıdaki kod baştan sona çalıştıran bu kişi ciftSayiCn ile karşılaşacktır
	ciftSayiCn <- toplam

}
func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
	}

	tekSayiCn <- toplam

}

/*
işlemlerin asenkron şekilde bittiğini sağlıklı bir şekilde takibini channeldan ypaıyoruz

diyelim ki iki fonksiyondan gelen sayıları toplayacağız
iki fonksiyonda işlemini bitirsin ve biz bunları carpacağımızı veya toplayacağımızı varsayalım
yaptığımız bir işlem sürerken sistemin kitlenmemesi gerekiyor

bunu nasıl yacapacağız ?
func CiftSayilar() {
	for i := 0; i < 10; i += 2 {
		//10 a kadaki cift sayıları yazacak
		fmt.Println("Cift sayı ", i)

	}

}
func TekSayilar() {

	for i := 1; i < 10; i += 2 {
		//10 a kadaki tek sayıları yazacak
		fmt.Println("Tek sayı ", i)

	}

}
bu fonksiyonlara parametre oalrak bir tane channel parametresi göndereceğiz

func CiftSayilar(ciftSayiCn chan int) {
	var toplam int = 0
	for i := 0; i < 10; i += 2 {
		//10 a kadaki cift sayıları yazacak
		fmt.Println("Cift sayı ", i)

	}

}
buradaki channel bir kanal görevi görür ve bununla bu işlemin bittiğini garanti edeceğiz
func CiftSayilar(ciftSayiCn chan int) {
	var toplam int = 0
	for i := 0; i < 10; i += 2 {
		toplam = toplam + i

	}
	//işlemin en sonuna bunu yapıyoruz
	//bu aşağıdaki kod baştan sona çalıştıran bu kişi ciftSayiCn ile karşılaşacktır
	ciftSayiCn <- toplam

}
func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i < 10; i += 2 {
		toplam = toplam + i
	}

	tekSayiCn <- toplam
}
//bu işlemlerin bittiğini yakalıyor olmamız gerektiğinden bu şekilde kullandık

//---
//main kısmı
//burada bir değişken tanımlayıp fonksiyonlarımın icersindeki chanella erişmem gerekiyor
	ciftSayiCn := make(chan int)
	tekSayiCn := make(chan int) //bu sayede fonksiyonların içersindeki channeları atayabiliyoruz aynı dizilerdeki gibi
	go gochannels.CiftSayilar(ciftSayiCn)
	go gochannels.TekSayilar(tekSayiCn)
	//var carpim int = ciftSayiCn * tekSayiCn
	//buradaki değerler bizim icin asla ama asla bir sayı değil sadece sayı taşıyan bir channel struct gibi düşünebiliriz
	//bunları okuyabilmemiz için yapmamız gereken işlem

	ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
	var carpim int = ciftSayiToplam * tekSayiToplam
	//chanellın değerini okuyup ciftSayiCn'yi ciftSayiToplama aktarıyoruz
	fmt.Println("Çarpım : ", carpim)

	burada iki fonksiyonda aynı anda çalıştı ve işlemler bittiği zamana atama yaptı
	carpımlarını yazdırabildik

	ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
	bu işlem yapıldığı icin mainde direkt aşağıya devam etmiyor time.sleep kullanmamıza gerek kalmıyor yani
	bu işlemler ne zaman bitiyorsa ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn bu kod çalışıyor


	channellarla yaşam döngüsünü kullanabiliyoruz

*/

package loops

import "fmt"

//kullanıcıdan sayı girmesini iste
//23 : Asaldır
func WorkShop2() {
	sayi := 0

	fmt.Println("Bir sayı giriniz:")
	fmt.Scanln(&sayi)

	asalMi := true

	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
			break
		}
	}

	if asalMi {
		fmt.Println(sayi, "Asaldır")
	} else {
		fmt.Println(sayi, "Asal değildir")
	}

	fmt.Println("Programdan çıkılıyor...")
}

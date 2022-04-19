package loops

import "fmt"

func WorkShop1() {
	aklimdakiSayi := 80
	tahminEdilenSayi := 0
	count := 1

	fmt.Println("Tahmin edilecek sayı:")
	fmt.Scanln(&tahminEdilenSayi)

	for tahminEdilenSayi != aklimdakiSayi {
		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Daha küçük bir sayı tahmin ediniz.")
		} else {
			fmt.Println("Daha büyük bir sayı tahmin ediniz.")
		}

		fmt.Scanln(&tahminEdilenSayi)
		count++
	}

	basariDurumu := ""

	if count > 0 && count <= 3 {
		basariDurumu = "Süper"
	} else if count > 3 && count <= 5 {
		basariDurumu = "Çok İyi"
	} else if count > 5 && count <= 7 {
		basariDurumu = "İyi"
	} else if count > 7 && count <= 9 {
		basariDurumu = "Kötü"
	} else {
		basariDurumu = "Çok Kötü"
	}

	fmt.Printf("Tebrikler, doğru tahmin ettiniz. %v Tahmin : %v\n", count, basariDurumu)
}

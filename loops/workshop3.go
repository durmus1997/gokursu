package loops

import "fmt"

func WorkShop3() {
	// arkadaş sayı olup olmadığını kontrol edin

	number1 := 0
	number2 := 0

	sumOfNumber1 := 0
	sumOfNumber2 := 0

	fmt.Println("Birinci sayıyı giriniz:")
	fmt.Scanln(&number1)
	fmt.Println("İkinci sayıyı giriniz:")
	fmt.Scanln(&number2)

	for i := 1; i < number1; i++ {
		if number1%i == 0 {
			sumOfNumber1 = sumOfNumber1 + i
		}
	}

	for i := 1; i < number2; i++ {
		if number2%i == 0 {
			sumOfNumber2 = sumOfNumber2 + i
		}
	}

	if sumOfNumber1 == number2 && sumOfNumber2 == number1 {
		fmt.Println(number1, "ve", number2, "Arkadaş sayısıdır.")
	} else {
		fmt.Println(number1, "ve", number2, "Arkadaş sayısı değildir.")
	}

}

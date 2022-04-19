package conditionals

import "fmt"

func WorkShop1() {
	// üç int değişken tanımla
	// Ekrana en büyük değeri yazdırınız

	var number1 int = 1
	var number2 int = 2
	var number3 int = 3

	// if number1 > number2 {
	// 	if number1 > number3 {
	// 		fmt.Println(number1)
	// 	} else {
	// 		fmt.Println(number3)
	// 	}
	// } else {
	// 	if number2 > number3 {
	// 		fmt.Println(number2)
	// 	} else {
	// 		fmt.Println(number3)
	// 	}
	// }

	var max int = number1

	if number2 > max {
		max = number2
	}

	if number3 > max {
		max = number3
	}

	fmt.Printf("En büyük sayı: %v", max)
}

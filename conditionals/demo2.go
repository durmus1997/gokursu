package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 1000

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki para yetersiz")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Paranız hazırlanıyor...")
		fmt.Println("Dikkat, hesapta para kalmadı...")
		hesap = hesap - cekilmekIstenen
	} else {
		fmt.Println("Paranız hazırlanıyor...")
		hesap = hesap - cekilmekIstenen
	}

	fmt.Printf("Bitti. Hesaptaki para: %v\n", hesap)
}

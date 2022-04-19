package loops

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dünya!"

	i := 1
	//infinite loop
	for i <= 5 {
		fmt.Println(metin)
		i = i + 1
	}

	fmt.Println("Loop bitmiştir.")
}

package arrays

import "fmt"

func Demo2() {
	var cities [5]string
	cities[0] = "Los Angeles"
	cities[1] = "New York"
	cities[2] = "Chicago"
	cities[3] = "Houston"
	cities[4] = "İstanbul"
	fmt.Println(cities)

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

}

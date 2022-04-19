package arrays

import "fmt"

func Demo3() {
	numbers := [5]int{20, 30, 50, 10, 2}
	fmt.Println(numbers)

	max := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	fmt.Println("En Büyük:", max)
}

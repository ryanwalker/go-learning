package topics

import (
	"fmt"
)

func SlicesBasics() {
	fmt.Println("Slices Basics")

	var numbersSlice []int
	slice := []int{1, 3, 5}
	fmt.Println(slice)

	numbersSlice = append(numbersSlice, 3)
	numbersSlice = append(numbersSlice, 12)

	fmt.Print(numbersSlice)

	for index, value := range slice {
		fmt.Println(index, value)
	}

}

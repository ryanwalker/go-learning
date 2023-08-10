package topics

import (
	"fmt"
	"os"
	"strconv"
)

func Cla() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need more args for", arguments[0])
		return
	}

	var min, max float64 = 0, 0

	for i := 1; i < len(arguments); i++ {
		val, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
		fmt.Println(val)
	}

	fmt.Println("Min:", min, " , Max:", max)
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	even := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range even {
		if even[i]%2 == 0 {
			fmt.Println(strconv.Itoa(even[i]) + " is even")
		} else {
			fmt.Println(strconv.Itoa(even[i]) + " is odd")
		}
	}

}

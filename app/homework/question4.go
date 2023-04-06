package main

import "fmt"

func main() {
	numbers := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	for maxIndex := len(numbers) - 1; 0 < maxIndex; maxIndex-- {

		for i := 0; i < maxIndex; i++ {

			if numbers[i] > numbers[i+1] {
				index := numbers[i]
				numbers[i] = numbers[i+1]
				numbers[i+1] = index

			}
		}

	}

	fmt.Println(numbers) 
}

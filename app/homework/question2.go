package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Println("j=", i*j)
		}
		fmt.Println(i,"の段")
	}
	fmt.Println("終了")
}

package main

import (
	"fmt"
)

func if_test(i int) {
	if i < 60 {
		fmt.Println("不合格です。")
	} else if i < 80 {
		fmt.Println("合格です。")
	} else {
		fmt.Println("大金星！！")
	}
}

func switch_test(i int) {
	switch {
	case i < 60:
		fmt.Println("不合格")
	case i < 80:
		fmt.Println("合格")
	default:
		fmt.Println("大金星")
	}
}

func main() {
	i := 70
	if_test(i)
	switch_test(i)
}

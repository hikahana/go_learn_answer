package main

import (
	"fmt"
)

func main() {

	i := map[string]string{"apple": "120", "banana": "100","orange": "150"}
    for k, v := range i {
        fmt.Printf("%sは%s円\n", k, v)
    }

    for k := range i {
        fmt.Println("key:", k)
    }


}

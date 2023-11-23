package main

import "fmt"

func main() {
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		fns = append(fns, func() {
			fmt.Printf("hello, this is %d\n", i)
		})

	}
	for _, fn := range fns {
		fn()
	}

}

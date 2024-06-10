package main

import (
	"fmt"
)

func main() {
	for i := 10; i < 100; i++ {
		condition := (i / 10) * (i % 10) * 2
		if condition == i {
			fmt.Print(condition, ",")
		}
	}
}

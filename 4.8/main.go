package main

import (
	"fmt"
)

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	counter := 1
	for {
		var str string
		_, _ = fmt.Scan(&str)
		if s == str {
			fmt.Println(counter)
			break
		} else {
			counter++
		}
	}
}

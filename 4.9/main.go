package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int
	_, _ = fmt.Scan(&a, &b, &c, &d, &e)
	var count int
	for x := 0; x <= 1000; x++ {
		if (x - e) == 0 {
			continue
		} else if (a*x*x*x+b*x*x+c*x+d)/(x-e) == 0 {
			count++
			if count == 3 {
				break
			}
		}
	}
	fmt.Println(count)
}

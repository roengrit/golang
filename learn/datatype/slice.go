package main

import (
	"fmt"
)

func main() {
	
	s := make([]int, 5)
	s[0] = 10
	s = append(s, 5)
	fmt.Println(s)
}

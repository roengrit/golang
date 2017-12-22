package main

import (
	"fmt"
)

func main() {
	var t [10]int
	t[0] = 1
	t[1] = 2
	t[2] = 3
	t[3] = 4
	t[4] = 5
	t[5] = 100

	var ta = [3]int{1, 2, 3}

	var t2d = [2][2]int{{1, 2}, {10, 2}}
	t2d[0][0] = 100

	for index := 0; index < 2; index++ {
		fmt.Println(t[index])
	}

	for _, p := range t {
		fmt.Println(p)
	}
	fmt.Println("----->")
	for _, p := range ta {
		fmt.Println(p)
	}
	fmt.Println("----->")
	fmt.Println(t2d)
	fmt.Println("----->")
	fmt.Println(t[:4])  //print index 0--->3
	fmt.Println(t[4:5]) //print index 4,4
	fmt.Println(t[4:6]) //print index 4-->5
	fmt.Println(t[3:])  //print index 3-->end
}

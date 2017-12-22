package main

import (
	"fmt"
)

func main() {

	fmt.Print("Input : ")
	var inputKey string
	fmt.Scan(&inputKey)
	if inputKey == "" {
		fmt.Println("Please input : ")
		fmt.Scan(&inputKey)
	}
	switch inputKey {
	case "roengrit":
		fmt.Println("I,m ", inputKey)
	case "roengrit1":
		fmt.Println("I,m ", inputKey, 485)
	}
}

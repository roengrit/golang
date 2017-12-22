package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	if _, ok := m["k3"]; ok {
		fmt.Println("has key k3")
	} else {
		fmt.Println("not has key k3")
	}

	if val, ok := m["k1"]; ok {
		fmt.Println("has key k1 value : ", val)
	} else {
		fmt.Println("not has key k1")
	}

	delete(m, "k3")
	if _, ok := m["k3"]; ok {
		fmt.Println("has key k3")
	} else {
		fmt.Println("delete key k3")
	}

	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	var nmap = map[string]float64{
		"solo": 1585.80,
		"mana": 4585.30,
	}
	for key, value := range nmap {
		fmt.Println("Key:", key, "Value:", value)
	}
}

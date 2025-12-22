package main

import "fmt"

func main() {
	mapa := map[string]int{"alice": 25, "bob": 30}
	for key, value := range mapa {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
}

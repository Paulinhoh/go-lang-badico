package main

import "fmt"

func main() {
	// criando com make
	mapa := make(map[string]int)
	mapa["alice"] = 25
	mapa["bob"] = 30
	fmt.Println(mapa)
}

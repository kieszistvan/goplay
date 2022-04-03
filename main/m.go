package main

import (
	"fmt"

	"github.com/kieszistvan/goplay/util/ints"
	"github.com/kieszistvan/goplay/util/strings"
)

func main() {
	fmt.Println(strings.Sup("pista"))
	fmt.Println(ints.MaxInt([]int{10, 1, 1000, 100}...))
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "23"
	b := 23
	c, _ := strconv.Atoi(a)
	fmt.Println(c, b)
}

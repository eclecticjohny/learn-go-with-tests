package main

import (
	"fmt"
	"strings"
)

func Repeat(c string, r int) string {
	return strings.Repeat(c, r)
}

func main() {
	fmt.Println(Repeat("a", 5))
}

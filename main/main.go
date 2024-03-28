package main

import (
	"fmt"
	"onehoax/go_sample/more_strings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello world!!")

	fmt.Println(more_strings.ReverseRunes(("!oG ,olleH")))

	fmt.Println(cmp.Diff("Hello world", "Hello go"))
}

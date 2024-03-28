package main

import (
	"fmt"
	"onehoax/go_sample/morestrings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello world!!")

	fmt.Println(morestrings.ReverseRunes(("!oG ,olleH")))

	fmt.Println(cmp.Diff("Hello world", "Hello go"))
}

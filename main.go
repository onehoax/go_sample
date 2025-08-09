package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/onehoax/go_sample/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello world", "hello go"))
}

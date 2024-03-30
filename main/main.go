package main

import (
	"fmt"
	"onehoax/go_sample/morenumbers"
	"onehoax/go_sample/morestrings"
	"onehoax/go_sample/vars"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("########### strings ###########")
	fmt.Println(morestrings.ReverseRunes(("!oG ,olleH")))
	fmt.Println(cmp.Diff("Hello world", "Hello go"))
	morestrings.IllegalChars()
	fmt.Println()

	fmt.Println("########### numbers ###########")
	fmt.Println(morenumbers.MB.String())
	fmt.Println(morenumbers.GB.String())
	morenumbers.LeftShitf(2, 5)
	morenumbers.RightShift(64, 5)
	fmt.Println()

	fmt.Println("########### vars ###########")
	vars.PrintVars()
	fmt.Println()
}

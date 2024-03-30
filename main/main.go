package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"onehoax/go_sample/morenumbers"
	"onehoax/go_sample/morestrings"
	"onehoax/go_sample/server"
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

	fmt.Println("########### server ###########")
	flag.Parse()
	http.Handle("/", http.HandlerFunc(server.QR))
	log.Println("Server runnin at http://localhost:1718")
	err := http.ListenAndServe(*server.Addr, nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

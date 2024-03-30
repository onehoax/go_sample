package vars

import (
	"fmt"
	"os"
)

var someVar string = "some var"

var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func PrintVars() {
	var1, var2, var3 := "var1", "var2", "var3"

	fmt.Println(someVar)
	fmt.Println(home)
	fmt.Println(user)
	fmt.Println(gopath)
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)

}

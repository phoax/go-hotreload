package main

import (
	"fmt"

	"github.com/phoax/go-hotreload/mypackage/mytest"
)

func main() {
	message := "Call Package:"
	fmt.Println(message, mytest.GetMyTest())
}

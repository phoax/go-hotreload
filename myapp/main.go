package main

import (
	"fmt"

	"github.com/phoax/go-hotreload/mypackage/mytest"
)

func main() {
	fmt.Println("Call Package:", mytest.GetMyTest())
}

package main

import (
	"ISC/Checks"
	"fmt"
)

func main() {
	result := Checks.RunCheck()
	fmt.Println(result)
}

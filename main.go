package main

import (
	"fmt"
	"strconv"
)

var letter = "a"
var number uint64 = 18446744073709551615

func main() {
	fmt.Print("number is: ", strconv.FormatUint(number, 16), "\n", "letter is: ", "letter", "\n")
}

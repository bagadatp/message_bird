package main

import (
	"fmt"
	"github.com/bagadatp/message_bird/pkg/sample"
)
func main() {

	input := 10
	result := sample.IntToString(input)
	fmt.Printf("Converted input int %v to string '%v'\n", input, result)
}

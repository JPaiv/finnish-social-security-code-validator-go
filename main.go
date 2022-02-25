package main

import (
	"fmt"

	"github.com/JPaiv/finnish-social-security-code-validator-go/src/Validate"
)

func main() {
	result := Validate.ValidatesocialSecurityCode("270694-0836")
	fmt.Printf("%t", result)
}

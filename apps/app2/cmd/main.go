package main

import (
	"fmt"

	someshareddep "github.com/DarylvdBerg/mono-release-me-poc/shared/someshareddep"
)

func main() {
	someshareddep.Hello()
	fmt.Print("HELLo")
	fmt.Print("HELLo")
}
package main

import (
	"fmt"

	someshareddep "github.com/DarylvdBerg/mono-release-me-poc/shared/some-shared-dep"
)

func main() {
	someshareddep.Hello()
	fmt.Print("HELLo")
}
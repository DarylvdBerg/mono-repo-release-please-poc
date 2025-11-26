package someshareddep

import "fmt"

func Hello() string {
	fmt.Print("Hello")
	return "hello"
}
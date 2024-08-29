package main

import (
	"fmt"

	"github.com/goJwt/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}
func main() {
	fmt.Println("Hello World")

}

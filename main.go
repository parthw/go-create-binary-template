package main

import (
	"fmt"

	"github.com/parthw/go-create-binary-template/cmd"
)

var Version = "development"

func main() {
	fmt.Println("Version: ", Version)
	cmd.Execute()
}

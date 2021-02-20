package main

import (
	"fmt"

	"xqtr/cmd"

	_ "github.com/go-git/go-git/v5/_examples"
	// "github.com/simeonpoot/xqtr/cmd"
)

func main() {
	fmt.Println("start app from main")
	cmd.Execute()
}

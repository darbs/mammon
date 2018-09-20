package main

import (
	"fmt"
	"os"
	"github.com/darbs/mammon"
	)

func main() {
	fmt.Println("Hello, world.")
	dial(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
}

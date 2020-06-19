package main

import (
	"fmt"
	"runtime"

	"github.com/carlescere/scheduler"
)

func main() {
	scheduler.Every(5).Seconds().Run(printSuccess)
	runtime.Goexit()
}

func printSuccess() {
	fmt.Printf("success!! \n")
}

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/carlescere/scheduler"
)

var count = 0

func main() {
	scheduler.Every(3).Seconds().Run(printSuccess)
	runtime.Goexit()
}

func printSuccess() {
	count++
	if count > 5 {
		fmt.Println("finish...")
		os.Exit(0)
	} else {
		fmt.Println("success!!")
	}
}

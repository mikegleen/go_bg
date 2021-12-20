package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf(
		"Hello go from %s/%s\n",
		runtime.GOOS,
		runtime.GOARCH,
	)
}

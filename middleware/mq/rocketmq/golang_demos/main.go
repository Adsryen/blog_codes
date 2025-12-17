package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	mode := flag.String("mode", "loop", "mode: loop or simple")
	flag.Parse()
	// MainSimple()
	switch *mode {
	case "loop": // let container keep running
		for {
			time.Sleep(time.Second)
		}
	case "simple":
		MainSimple()
	default:
		fmt.Println("Invalid mode")
	}

}

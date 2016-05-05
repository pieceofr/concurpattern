package main

import (
	"flag"
	"fmt"
)

func main() {
	pattern := cmdInput()

	switch pattern {
	case 1:
		Simple()
	case 2:
		TimerP()
	case 3:
		PingPon()
	case 4:
		FadeIn()
	case 5:
		FadeOut()
	case 6:
		WorkerSubWorker()
	default:
		Simple()
	}

	fmt.Println("Leaving Main")
}

func cmdInput() int {
	descr := "1:Simple 2:Timer 3:PingPon 4:FadeIn 5:Worker 6:WorkerSubWorker"
	patternPtr := flag.Int("pattern", 1, descr)
	flag.Parse()
	return *patternPtr
}

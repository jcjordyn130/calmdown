package main

import (
	"bufio"
	"os"
	"time"
	"flag"
)

func main() {
	var delay = flag.Duration("delay", 0, "Delay between bytes")
	flag.Parse()
	stdin := bufio.NewReader(os.Stdin)
	for {
		b, err := stdin.ReadByte()
		if err != nil {
			os.Exit(0)
		}
		os.Stdout.Write([]byte{b})
		time.Sleep(*delay)
	}
}

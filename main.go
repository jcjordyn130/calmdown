package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:")
		fmt.Println("	calmdown [delay between bytes in ms]")
		os.Exit(1)
	}
	delayint, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Argument 1 must be an Integer.")
		os.Exit(1)
	}
	delay := time.Duration(int64(delayint)) * time.Millisecond
	stdin := bufio.NewReader(os.Stdin)
	var b []byte
	b = append(b, 0x00)
	c := true
	for c {
		b[0], _ = stdin.ReadByte()
		if b[0] == 0 {
			os.Exit(0)
		}
		os.Stdout.Write(b)
		time.Sleep(delay)
	}
}

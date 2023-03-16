package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	defX = 2
	defY = 2
)

var startWith = flag.Int("start", 1, "Start with this number")

func init() { flag.Usage = Usage }

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	prefix := flag.Arg(0)
	x, y := defX, defY
	if flag.NArg() > 1 {
		x = a2iOrDie(flag.Arg(1))
	}
	if flag.NArg() > 2 {
		y = a2iOrDie(flag.Arg(2))
	}

	n := *startWith
	for col := 0; col < x; col++ {
		for row := 0; row < y; row++ {
			fmt.Printf(":%s-%d:", prefix, n)
			n++
		}
		fmt.Println()
	}
}

func a2iOrDie(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func Usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] <prefix> [x] [y]\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "Where x is the number of emojis in the line (default:  %d)\n", defX)
	fmt.Fprintf(flag.CommandLine.Output(), "      y is the number of lines              (default:  %d)\n\n", defY)
	fmt.Fprintf(flag.CommandLine.Output(), "Flags:")
	flag.PrintDefaults()
}

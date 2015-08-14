package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
)

var frequency = flag.Float64("f", 0.1, "frequency of line selection")

func main() {
	var exitStatus int

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		args = []string{"-"}
	}
	for _, f := range args {
		func() {
			var err error
			var in *os.File

			if f == "-" {
				// XCU7: If a file is '-', the cat utility shall read from the standard input at that point in the sequence
				in = os.Stdin
			} else {
				in, err = os.Open(f)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					exitStatus = 2
					return
				}
				defer in.Close()
			}
			err = frequencyCopy(os.Stdout, in)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				exitStatus = 2
				return
			}
		}()
	}
	os.Exit(exitStatus)
}

func frequencyCopy(w io.Writer, r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l := scanner.Text()
		if rand.Float64() < *frequency {
			fmt.Fprintln(w, l)
		}
	}
	err := scanner.Err()
	return err
}

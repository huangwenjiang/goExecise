package main
// print the file name of dup lines
import (
	"bufio"
	"fmt"
	"os"
)

type filenames []string

func main() {
	counts := make(map[string]filenames)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "STDOUT")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v", err)
			}
			countLines(f, counts, arg)
		}
	}
	for line, n := range counts {
		if len(n) > 1 {
			fmt.Printf("%v\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]filenames, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		v := counts[input.Text()]
		counts[input.Text()] = append(v, filename)
	}
}

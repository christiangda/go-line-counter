package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/christiangda/go-line-counter/algorithms"
)

// CountLines helper function to count the number of lines in a file
// using the given algorithm
func CountLines(c algorithms.Counter, r io.Reader) (int, error) {
	tStart := time.Now()
	defer func() {
		fmt.Printf("CountLines took %s  -> ", time.Since(tStart))
	}()

	return c.Count(r)
}

// algos is a map of algorithms.Counter
var algos = map[string]algorithms.Counter{
	"ScannerLineCounter":        &algorithms.ScannerLineCounter{},
	"JimBLineCounter":           &algorithms.JimBLineCounter{},
	"DanielCastilloLineCounter": &algorithms.DanielCastilloLineCounter{},
	"FuzLineCounter":            &algorithms.FuzLineCounter{},
}

func main() {
	file, err := os.Open("testdata/Subnational-period-life-tables-2017-2019-CSV.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for name, algo := range algos {

		// Reset the file pointer to the beginning of the file
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			fmt.Println(err)
		}

		lines, err := CountLines(algo, file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Algorithm: %s, lines: %d\n", name, lines)
	}
}

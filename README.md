# go-line-counter

I was curious about the fastest and efficient way to count the number of lines in a file using Go, so, I found a StackOverflow post that had a few different ways to do it.

I decided to try them out and see which one was the fastest.

The StackOverflow post [Golang: How do I determine the number of lines in a file efficiently?](https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently)

## Conclusion

The fastest way to count the number of lines in a file is to use the `bytes.Count()` of a buffer, the algorithm provided by [JimB](https://stackoverflow.com/users/32880/jimb).

[The implementation:](internal/algorithms.go)

```go
type JimBLineCounter struct {
	Size int    // Size of the buffer
	Sep  string // End of line character
}

// source: https://stackoverflow.com/a/24563853/3925852
func (b *JimBLineCounter) Count(r io.Reader) (int, error) {
	defaultSize := 32 * 1024
	defaultEndLine := "\n"

	if b.Size == 0 {
		b.Size = defaultSize
	}

	if b.Sep == "" {
		b.Sep = defaultEndLine
	}

	buf := make([]byte, b.Size)
	var count int

	for {
		n, err := r.Read(buf)
		count += bytes.Count(buf[:n], []byte(b.Sep))

		if err != nil {
			if err == io.EOF {
				return count, nil
			}
			return count, err
		}

	}
}
```

## Usage

```bash
go run main.go
```

Output:

```bash
CountLines took 10.056041ms  -> Algorithm: ScannerLineCounter, lines: 255361
CountLines took 2.215916ms  -> Algorithm: JimBLineCounter, lines: 255361
CountLines took 6.099542ms  -> Algorithm: DanielCastilloLineCounter, lines: 255361
CountLines took 11.25625ms  -> Algorithm: FuzLineCounter, lines: 255361
```

__Testing:__

```bash
go test ./...
```

__Benchmarking:__

```bash
go test -bench . -benchmem -count=5 -benchtime=1000x | tee benchmarking_stats.txt
```

Output:

```bash
goos: darwin
goarch: arm64
pkg: github.com/christiangda/go-line-counter
BenchmarkScannerLineCounter-10                      1000              9673 ns/op            4096 B/op          1 allocs/op
BenchmarkScannerLineCounter-10                      1000             10290 ns/op            4096 B/op          1 allocs/op
BenchmarkScannerLineCounter-10                      1000             10065 ns/op            4096 B/op          1 allocs/op
BenchmarkScannerLineCounter-10                      1000             10238 ns/op            4096 B/op          1 allocs/op
BenchmarkScannerLineCounter-10                      1000             10059 ns/op            4096 B/op          1 allocs/op
BenchmarkJimBLineCounter-10                         1000              5063 ns/op           32770 B/op          1 allocs/op
BenchmarkJimBLineCounter-10                         1000              4935 ns/op           32768 B/op          1 allocs/op
BenchmarkJimBLineCounter-10                         1000              4866 ns/op           32768 B/op          1 allocs/op
BenchmarkJimBLineCounter-10                         1000              4755 ns/op           32768 B/op          1 allocs/op
BenchmarkJimBLineCounter-10                         1000              4665 ns/op           32768 B/op          1 allocs/op
BenchmarkDanielCastilloLineCounter-10               1000             11395 ns/op           65536 B/op          1 allocs/op
BenchmarkDanielCastilloLineCounter-10               1000             11601 ns/op           65536 B/op          1 allocs/op
BenchmarkDanielCastilloLineCounter-10               1000             11119 ns/op           65536 B/op          1 allocs/op
BenchmarkDanielCastilloLineCounter-10               1000             11360 ns/op           65536 B/op          1 allocs/op
BenchmarkDanielCastilloLineCounter-10               1000             11074 ns/op           65536 B/op          1 allocs/op
BenchmarkFuzLineCounter-10                          1000             11791 ns/op            8192 B/op          1 allocs/op
BenchmarkFuzLineCounter-10                          1000             11560 ns/op            8192 B/op          1 allocs/op
BenchmarkFuzLineCounter-10                          1000             11575 ns/op            8192 B/op          1 allocs/op
BenchmarkFuzLineCounter-10                          1000             11607 ns/op            8192 B/op          1 allocs/op
BenchmarkFuzLineCounter-10                          1000             11636 ns/op            8192 B/op          1 allocs/op
PASS
ok      github.com/christiangda/go-line-counter 0.553s
```

## Challenge

Do you know a faster way to count the number of lines in a file? If so, please, let me know doing a PR or opening an issue.

## Credits

+ StackOverflow post: <https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently>
+ StackOverflow user: <https://stackoverflow.com/users/786821/sunsparc>
+ StackOverflow user: <https://stackoverflow.com/users/32880/jimb>
+ StackOverflow user: <https://stackoverflow.com/users/8440412/daniel-castillo>
+ StackOverflow user: <https://stackoverflow.com/users/417501/fuz>

## Test Data

File: [www.stats.govt.nz-> Subnational-period-life-tables-2017-2019-CSV.csv](https://www.stats.govt.nz/assets/Uploads/National-and-subnational-period-life-tables/National-and-subnational-period-life-tables-2017-2019/Download-data/Subnational-period-life-tables-2017-2019-CSV.csv)
Number of lines: 1,713,361
Size:  22,551,849 bytes (22.6 MB on disk)


## License

[Apache License 2.0](LICENSE)

package algorithms

import (
	"bufio"
	"bytes"
	"io"
)

// Counter is the interface that wraps the basic Count method.
// and any type that implements this interface can be used as a counter
type Counter interface {
	Count(io.ReadSeeker) (int, error)
}

// ScannerLineCounter counts the number of lines in a file using bufio.Scanner
type ScannerLineCounter struct{}

func (s *ScannerLineCounter) Count(r io.ReadSeeker) (int, error) {
	defer func() {
		_, err := r.Seek(0, io.SeekStart)
		if err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines) // this was my modification

	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines, scanner.Err()
}

// JimBLineCounter counts the number of bytes in a file using io.Reader
// proposed by https://stackoverflow.com/users/32880/jimb
type JimBLineCounter struct {
	Size int    // Size of the buffer
	Sep  string // End of line character
}

func (b *JimBLineCounter) Count(r io.ReadSeeker) (int, error) {
	defer func() {
		_, err := r.Seek(0, io.SeekStart)
		if err != nil {
			panic(err)
		}
	}()

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

// DanielCastilloLineCounter counts the number of bytes in a file using io.Reader
// proposed by https://stackoverflow.com/users/8440412/daniel-castillo
type DanielCastilloLineCounter struct {
	Size int    // Size of the buffer
	Sep  string // End of line character
}

func (b *DanielCastilloLineCounter) Count(r io.ReadSeeker) (int, error) {
	defer func() {
		_, err := r.Seek(0, io.SeekStart)
		if err != nil {
			panic(err)
		}
	}()

	defaultSize := bufio.MaxScanTokenSize
	defaultEndLine := "\n"

	if b.Size == 0 {
		b.Size = defaultSize
	}

	if b.Sep == "" {
		b.Sep = defaultEndLine
	}

	sepByte := []byte(b.Sep)[0]

	buf := make([]byte, b.Size)
	var count int

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], sepByte)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}

// FuzLineCounter counts the number of bytes in a file using io.Reader
// proposed by https://stackoverflow.com/users/417501/fuz
type FuzLineCounter struct {
	Size int    // Size of the buffer
	Sep  string // End of line character
}

// Count counts the number of lines in a file using io.Reader
func (b *FuzLineCounter) Count(r io.ReadSeeker) (int, error) {
	defer func() {
		_, err := r.Seek(0, io.SeekStart)
		if err != nil {
			panic(err)
		}
	}()

	defaultSize := 8 * 1024
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
		c, err := r.Read(buf)
		if err != nil {
			if err == io.EOF && c == 0 {
				break
			} else {
				return count, err
			}
		}

		for _, b := range buf[:c] {
			if b == '\n' {
				count++
			}
		}
	}

	return count, nil
}

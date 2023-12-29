package main

import (
	"os"
	"testing"

	"github.com/christiangda/go-line-counter/algorithms"
)

func TestScannerLineCounter_Count(t *testing.T) {
	t.Run("with_Subnational-period-life-tables-2017-2019-CSV.csv", func(t *testing.T) {
		file, err := os.Open("testdata/Subnational-period-life-tables-2017-2019-CSV.csv")
		if err != nil {
			t.Error(err)
		}
		defer file.Close()

		// extracted: wc -l testdata/Subnational-period-life-tables-2017-2019-CSV.csv
		want := 255361

		s := algorithms.ScannerLineCounter{}
		got, err := s.Count(file)
		if err != nil {
			t.Errorf("ScannerCounter.Count() error = %v", err)
			return
		}
		if got != want {
			t.Errorf("ScannerCounter.Count() = %v, want %v", got, want)
		}
	})
}

func BenchmarkScannerLineCounter(b *testing.B) {
	file, err := os.Open("testdata/Subnational-period-life-tables-2017-2019-CSV.csv")
	if err != nil {
		b.Error(err)
	}
	defer file.Close()

	s := algorithms.ScannerLineCounter{}
	for n := 0; n < b.N; n++ {
		s.Count(file)
	}
}

func TestJimBLineCounter_Count(t *testing.T) {
	t.Run("with_Subnational-period-life-tables-2017-2019-CSV.csv", func(t *testing.T) {
		file, err := os.Open("testdata/Subnational-period-life-tables-2017-2019-CSV.csv")
		if err != nil {
			t.Error(err)
		}
		defer file.Close()

		// extracted: wc -l testdata/Subnational-period-life-tables-2017-2019-CSV.csv
		want := 255361

		s := algorithms.JimBLineCounter{}
		got, err := s.Count(file)
		if err != nil {
			t.Errorf("ScannerCounter.Count() error = %v", err)
			return
		}
		if got != want {
			t.Errorf("ScannerCounter.Count() = %v, want %v", got, want)
		}
	})
}

func BenchmarkJimBLineCounter(b *testing.B) {
	file, err := os.Open("testdata/Subnational-period-life-tables-2017-2019-CSV.csv")
	if err != nil {
		b.Error(err)
	}
	defer file.Close()

	s := algorithms.JimBLineCounter{}
	for n := 0; n < b.N; n++ {
		s.Count(file)
	}
}

func TestDanielCastilloLineCounter(t *testing.T) {
	t.Run("with_Subnational-period-life-tables-2017-2019-CSV.csv", func(t *testing.T) {
		file, err := os.Open("testdata/Subnational-period-life-tables-2017-2019-CSV.csv")
		if err != nil {
			t.Error(err)
		}
		defer file.Close()

		// extracted: wc -l testdata/Subnational-period-life-tables-2017-2019-CSV.csv
		want := 255361

		s := algorithms.DanielCastilloLineCounter{}
		got, err := s.Count(file)
		if err != nil {
			t.Errorf("ScannerCounter.Count() error = %v", err)
			return
		}
		if got != want {
			t.Errorf("ScannerCounter.Count() = %v, want %v", got, want)
		}
	})
}

func BenchmarkDanielCastilloLineCounter(b *testing.B) {
	file, err := os.Open("testdata/Subnational-period-life-tables-2017-2019-CSV.csv")
	if err != nil {
		b.Error(err)
	}
	defer file.Close()

	s := algorithms.DanielCastilloLineCounter{}
	for n := 0; n < b.N; n++ {
		s.Count(file)
	}
}

func TestFuzLineCounter(t *testing.T) {
	t.Run("with_Subnational-period-life-tables-2017-2019-CSV.csv", func(t *testing.T) {
		file, err := os.Open("testdata/Subnational-period-life-tables-2017-2019-CSV.csv")
		if err != nil {
			t.Error(err)
		}
		defer file.Close()

		// extracted: wc -l testdata/Subnational-period-life-tables-2017-2019-CSV.csv
		want := 255361

		s := algorithms.FuzLineCounter{}
		got, err := s.Count(file)
		if err != nil {
			t.Errorf("ScannerCounter.Count() error = %v", err)
			return
		}
		if got != want {
			t.Errorf("ScannerCounter.Count() = %v, want %v", got, want)
		}
	})
}

func BenchmarkFuzLineCounter(b *testing.B) {
	file, err := os.Open("testdata/Subnational-period-life-tables-2017-2019-CSV.csv")
	if err != nil {
		b.Error(err)
	}
	defer file.Close()

	s := algorithms.FuzLineCounter{}
	for n := 0; n < b.N; n++ {
		s.Count(file)
	}
}

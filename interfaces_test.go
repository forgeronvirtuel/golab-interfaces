package golabinterfaces

import (
	"bytes"
	"os"
	"testing"
)

const filesize = 1_500_000

func readDataFromFile() string {
	f, err := os.Open("./testdata.txt")
	if err != nil {
		panic(err)
	}

	buf := make([]byte, filesize)
	if _, err := f.Read(buf); err != nil {
		panic(err)
	}
	return string(buf)
}

func BenchmarkConvertData(b *testing.B) {
	txtData := readDataFromFile()
	buf := bytes.NewBuffer(make([]byte, filesize))

	for i := 0; i < b.N; i++ {
		if err := ConvertData(buf, txtData); err != nil {
			panic(err)
		}
	}
}

func BenchmarkConvertInterface(b *testing.B) {
	txtData := readDataFromFile()
	buf := bytes.NewBuffer(make([]byte, filesize))

	for i := 0; i < b.N; i++ {
		if err := ConvertInterface(buf, txtData); err != nil {
			panic(err)
		}
	}
}

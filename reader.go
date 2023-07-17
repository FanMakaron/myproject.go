package main

import (
	"fmt"
	"golang.org/x/tour/reader"
	"io"
	"os"
	"strings"
)

type MyReader struct{}

// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}
	return 1, nil
}

type rot13Reader struct {
	r io.Reader
}

// https://go.dev/tour/methods/23
func (rot rot13Reader) Read(b []byte) (int, error) {
	return rot.r.Read(b)
}

func main() {

	/**
	lesson example https://go.dev/tour/methods/21
	*/
	newReader := strings.NewReader("Hello, Reader!12")
	bytes := make([]byte, 8)
	for {
		number, err := newReader.Read(bytes)
		fmt.Printf("number = %v err = %v bytes = %v\number", number, err, bytes)
		fmt.Printf("bytes[:number] = %q\number", bytes[:number])
		if err == io.EOF {
			break
		}
	}

	for i, v := range bytes {
		fmt.Printf("index is %d value is %d\n", i, v)
	}

	reader.Validate(MyReader{})
	Validate(MyReader{})

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}

// Validate Перенес тест из golang.org/x/tour/reader
func Validate(r io.Reader) {
	b := make([]byte, 1024, 2048)
	i, o := 0, 0
	for ; i < 1<<20 && o < 1<<20; i++ { // test 1mb
		n, err := r.Read(b)
		for i, v := range b[:n] {
			if v != 'A' {
				fmt.Fprintf(os.Stderr, "got byte %x at offset %v, want 'A'\n", v, o+i)
				return
			}
		}
		o += n
		if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			return
		}
	}
	if o == 0 {
		fmt.Fprintf(os.Stderr, "read zero bytes after %d Read calls\n", i)
		return
	}
	fmt.Println("OK!")
}

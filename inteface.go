package main

import (
	"bytes"
)

func main() {
	// var w Writer = &ConsoleWriter{}
	// w.Write([]byte("Hello world"))
	// myInt := IntCounter(0)
	// var inc Increment = &myInt
	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(inc.Increment())
	// }
}

type Writer interface {
	Write([]byte) (int, error)
}

// type ConsoleWriter struct{}

// func (cw *ConsoleWriter) Write(data []byte) (int, error) {
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }

// type Incrementer interface {
// 	Increment() int
// }

// type IntCounter int

// func (ic *IntCounter) Increment() int {
// 	*ic++
// 	return int(*ic)
// }

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

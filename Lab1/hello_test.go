package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func setupRedirection(t *testing.T) (stdoutChannel chan string, tearDown func(t *testing.T)) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	originalStdout := os.Stdout

	os.Stdout = w

	stdoutChannel = make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		stdoutChannel <- buf.String()
		_ = r.Close()
	}()
	tearDown = func(t *testing.T) { os.Stdout = originalStdout }
	return
}

func Test_main(t *testing.T) {
	stdoutChannel, tearDown := setupRedirection(t)
	defer tearDown(t)
	main()
	_ = os.Stdout.Close()

	if expected, actual := "Hello, world!\n", <-stdoutChannel; expected != actual {
		t.Errorf("Expected and actual differ:\n<%v>\n<%v>\n", expected, actual)
	}
}

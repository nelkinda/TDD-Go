package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func setupRedirection(t *testing.T, fileToRedirect **os.File) (redirectedChannel chan string, tearDown func(t *testing.T)) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	originalStdout := *fileToRedirect

	*fileToRedirect = w

	redirectedChannel = make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		redirectedChannel <- buf.String()
		_ = r.Close()
	}()
	tearDown = func(t *testing.T) { *fileToRedirect = originalStdout }
	return
}

func Test_main(t *testing.T) {
	stdoutChannel, tearDown := setupRedirection(t, &os.Stdout)
	defer tearDown(t)
	main()
	_ = os.Stdout.Close()

	if expected, actual := "Hello, world!\n", <-stdoutChannel; expected != actual {
		t.Errorf("Expected and actual differ:\n<%v>\n<%v>\n", expected, actual)
	}
}

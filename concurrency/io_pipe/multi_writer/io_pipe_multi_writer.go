package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	buff := new(bytes.Buffer)
	mw := io.MultiWriter(os.Stdout, os.Stderr, buff)
	fmt.Fprintln(mw , "hello")
	fmt.Println("form buffer:" , buff)
}

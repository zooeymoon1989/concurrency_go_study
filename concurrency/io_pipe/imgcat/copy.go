package main

import (
	"encoding/base64"
	"io"
	"log"
	"strings"
)

func Copy(w io.Writer, r io.Reader) error {

	header := strings.NewReader("\033]1337;File=inline=1:")
	footer := strings.NewReader("\a\n")

	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		wc := base64.NewEncoder(base64.StdEncoding, pw)
		_, err := io.Copy(wc, r)
		if err != nil {
			pr.CloseWithError(err)
			return
		}

		if err := wc.Close(); err != nil {
			pr.CloseWithError(err)
		}
	}()

	_, err := io.Copy(w, io.MultiReader(header, pr, footer))

	return err
}

func NewWriter(w io.Writer) io.WriteCloser {
	pr, pw := io.Pipe()

	wc := &writer{pw, make(chan struct{})}

	go func() {
		defer close(wc.done)
		err := Copy(w, pr)
		if err != nil {
			log.Fatal(err)
		}
	}()

	return wc
}

type writer struct {
	pw   *io.PipeWriter
	done chan struct{}
}

func (w *writer) Write(p []byte) (n int, err error) {
	return w.pw.Write(p)
}

func (w *writer) Close() error {
	if err := w.pw.Close(); err != nil {
		return err
	}
	<-w.done
	return nil
}

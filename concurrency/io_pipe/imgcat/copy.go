package main

import (
	"encoding/base64"
	"fmt"
	"io"
)

func Copy(w io.Writer, r io.Reader) error {
	_, err := fmt.Fprintf(w, "\033]1337;File=inline=1:")
	if err != nil {
		return nil
	}
	wc := base64.NewEncoder(base64.StdEncoding, w)

	_, err = io.Copy(wc, r)
	if err != nil {
		return err
	}

	if err = wc.Close(); err != nil {
		return err
	}

	_ , err = fmt.Fprintf(w,"\a\n")
	if err != nil {
		return err
	}
	
	return nil
}

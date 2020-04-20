package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing paths of images to cat")
		os.Exit(2)
	}

	for _, path := range os.Args[1:] {
		if err := cat(path); err != nil {
			fmt.Fprintf(os.Stderr, "could not cat %s: %v\n", path, err)
		}
	}

}

func cat(path string) error {

	// ESC ] 1337 ; File = [optional arguments] : base-64 encoded file contents ^G
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	err = Copy(os.Stdout , f)

	if err != nil {
		return err
	}


	return nil
}

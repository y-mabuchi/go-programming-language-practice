package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, getErr := http.Get(url)
		if getErr != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", getErr)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying url %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}

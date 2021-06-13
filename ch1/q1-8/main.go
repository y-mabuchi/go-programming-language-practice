package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, getErr := http.Get(url)
		if getErr != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", getErr)
			os.Exit(1)
		}
		b, readErr := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if readErr != nil {
			fmt.Fprintf(os.Stdout, "fetch: reading %s: %v\n", url, readErr)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

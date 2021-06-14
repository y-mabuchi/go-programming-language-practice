package main

import (
	"fmt"
	"io/ioutil"
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
		b, readErr := ioutil.ReadAll(resp.Body)
		fmt.Printf("fetch status: %s\n", resp.Status)
		resp.Body.Close()
		if readErr != nil {
			fmt.Fprintf(os.Stderr, "fetch reaading %s: %v\n", url, getErr)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

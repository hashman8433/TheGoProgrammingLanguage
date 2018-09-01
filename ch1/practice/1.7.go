package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	urls := []string{"http://www.baidu.com"}
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch : %v\n", err)
			os.Exit(1)
		}

		b := os.Stdout
		written, err := io.Copy(b, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("written:", written)
	}
}

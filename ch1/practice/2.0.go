package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	urls := []string{"https://finance.sina.com.cn/"}
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp1, err1 := http.Get(url)
	resp2, err2 := http.Get(url)

	if err1 != nil || err2 != nil {
		fmt.Sprintf("%s %s", err1, err2)
		return
	}
	resp1Byte, err1 := ioutil.ReadAll(resp1.Body)
	resp2Byte, err2 := ioutil.ReadAll(resp2.Body)
	if err1 != nil || err2 != nil {
		fmt.Sprintf("%s %s", err1, err2)
		return
	}

	respStr1 := fmt.Sprintf("%s", resp1Byte)
	respStr2 := fmt.Sprintf("%s", resp2Byte)

	compareResult := strings.Compare(respStr1, respStr2)

	sesc := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", sesc, compareResult, url)
}

package main

import (
	"net/http"
	"time"
	"fmt"
	"flag"
)

func main() {
	url := flag.String("url", "http://localhost", "help message for s option")
	flag.Parse()

	for i := 0; ; i++ {
		access(i, *url)
	}
}

func access(i int, url string) {
	const layout = "2006-01-02 15:04:05"
	now := time.Now()

	fmt.Printf("%d\t%v\t", i, now.Format(layout))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	status := res.StatusCode
	fin := time.Now()

	fmt.Printf("%d\t%v\n", status, fin.Format(layout))
}


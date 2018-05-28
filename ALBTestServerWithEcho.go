package main

import (
	"fmt"
	"time"
	"net/http"
	"math/rand"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		r := rand.Intn(999999999)

		fmt.Printf("%d\taccess\t%v\n", r, time.Now().Format("2006-01-02 15:04:05"))

		time.Sleep(20 * time.Second)

		fmt.Printf("%d\treturn\t%v\n", r, time.Now().Format("2006-01-02 15:04:05"))

		w.Write([]byte("hello world"))
	})
	
	http.ListenAndServe(":8080", nil)
}
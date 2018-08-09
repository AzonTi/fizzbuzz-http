package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func handlerFizzBuzz(w http.ResponseWriter, r *http.Request) {
	numStr := strings.TrimPrefix(r.URL.Path, "/fizzbuzz/")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "invalid number")
		return
	}
	if num%15 == 0 {
		fmt.Fprintln(w, "Fizz Buzz")
	} else if num%3 == 0 {
		fmt.Fprintln(w, "Fizz")
	} else if num%5 == 0 {
		fmt.Fprintln(w, "Buzz")
	} else {
		fmt.Fprintf(w, "%d\n", num)
	}
}

func main() {
	http.HandleFunc("/fizzbuzz/", handlerFizzBuzz)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", response)
	http.ListenAndServe(":80", nil)
}

func response(w http.ResponseWriter, req *http.Request) {
	doSqrt()
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func greeting(s string) string {
	return fmt.Sprintf("<b>%s</b>", s)
}

func doSqrt() int {
	x := 0.0001
	i := 0
	for i < 1000000 {
		x = math.Sqrt(x)
		i++
	}
	return i
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
)

func Fibonacci(n int) []int {

	if n < 2 {
		return []int{0}
	}
	f := make([]int, n, n+1)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n-1; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}

func handleFibonacci(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fib := ps.ByName("number")
	number, _ := strconv.Atoi(fib)
	returned := Fibonacci(number)
	json.NewEncoder(w).Encode(returned)
}

func main() {
	router := httprouter.New()
	router.GET("/api/fibonacci/:number", handleFibonacci)
	log.Fatal(http.ListenAndServe(":8080", router))
}
package main

import "net/http"

func main() {
	dist := http.FileServer(http.Dir("dist"))
	err := http.ListenAndServe("localhost:8080", dist)
	if err != nil {
		panic(err)
	}
}

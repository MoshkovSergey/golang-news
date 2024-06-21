package main

import "net/http"

type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	http.Handle("/", MyHandler{})
	err := http.ListenAndServe(":8080", nil); if err != nil {
		panic(err)
	}
}

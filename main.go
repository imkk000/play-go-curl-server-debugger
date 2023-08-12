package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.ListenAndServe("127.0.0.1:8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cReq, _ := httputil.DumpRequest(r, true)

		fmt.Println("request:", string(cReq))

		w.WriteHeader(http.StatusOK)
	}))
}

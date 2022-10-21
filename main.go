package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		d, err := io.ReadAll(r.Body)
		data := string(d)

		if err != nil {
			http.Error(rw, "An error occured", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("An error occured"))
			return
		}

		fmt.Fprintf(rw, "Hello %s\n", data)
	})

	http.ListenAndServe(":9090", nil)
}

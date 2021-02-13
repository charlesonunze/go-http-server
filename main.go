package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, res *http.Request) {
		names := res.URL.Query()["name"]
		var name string

		if len(names) == 1 && len(names[0]) > 0 {
			name = names[0]
			rw.Write([]byte("Hello " + name))
			return
		}

		rw.Write([]byte("Hello world"))
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

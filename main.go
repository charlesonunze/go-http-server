package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, res *http.Request) {
		names := res.URL.Query()["name"]
		var name string

		if len(names) == 1 && len(names[0]) > 0 {
			name = names[0]
			m := map[string]string{"name": name}
			enc := json.NewEncoder(rw)
			enc.Encode(m)
			return
		}

		rw.Write([]byte("Hello world"))
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

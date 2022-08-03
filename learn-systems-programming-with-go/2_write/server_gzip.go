package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Output json to STDOUT
// And response gzip-decoded json
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(map[string]string{
		"Hello": "World",
	})
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		return
	}

	z := gzip.NewWriter(w)
	defer z.Close()

	writer := io.MultiWriter(os.Stdout, z)

	io.WriteString(writer, string(data))
	z.Flush()

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

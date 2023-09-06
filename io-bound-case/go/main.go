package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Simulate I/O-bound operation (e.g., reading from a file or database)
        time.Sleep(100 * time.Millisecond)
        fmt.Fprintf(w, "Go Standard Server Response")
    })

    fmt.Println("Go Standard Server listening on :3000")
    http.ListenAndServe(":3000", nil)
}

package main


import (
    "fmt"
    "log"
    "os"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}


func main() {
    http.HandleFunc("/", helloHandler) // Update this line of code

    var addressListening = fmt.Sprintf(":%s", os.Getenv("PORT"))

    fmt.Printf("Starting server at port %s\n", addressListening)
    if err := http.ListenAndServe(addressListening, nil); err != nil {
        log.Fatal(err)
    }
}
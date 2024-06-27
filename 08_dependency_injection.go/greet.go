package main


import (
  "io"
  //"os"
  "fmt"
  "log"
  "net/http"
)


func Greet(writer io.Writer, txt string) {
  fmt.Fprintf(writer, "Hello %s", txt)
}


func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
  Greet(w, "GOKU!")
}


func main() {
  //Greet(os.Stdout, "DRG")
  log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

package main

import (
	"fmt"
	"github.com/rs/xid"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server stating, id is: %s", xid.New().String())
	defer func() {
		log.Println("Recovering server")
		recover()
	}()

	http.HandleFunc("/", handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Request is:\n%s", getFromatedRequest(r))
	if err != nil {
		panic(err)
	}
}

func getFromatedRequest(r *http.Request) string {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("Body error " + err.Error())
	}
	return fmt.Sprintf("Url: %s\nMethod is: %s\nProtocol: %s\nHeaders:\nAccept: %s\nUser-Agent: %s\nBody: %s\n",
		r.URL, r.Method, r.Proto, r.Header.Get("Accept"), r.Header.Get("User-Agent"), string(b))
}

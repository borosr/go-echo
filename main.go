package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"

	"github.com/rs/xid"
)

//go:embed response.tmpl
var responseTemplate string

type Response struct {
	URL       string
	Method    string
	Protocol  string
	Headers   string
	Accepts   string
	UserAgent string
	Body      string
	Form      string
	Form2      string
}

var id = xid.New().String()

func main() {
	log.Printf("Server stating, id is: %s", id)
	defer func() {
		log.Println("Recovering server")
		recover()
	}()

	http.HandleFunc("/", handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Request [%s] is:\n%s", id, getFormattedRequest(r))
	if err != nil {
		panic(err)
	}
}

func getFormattedRequest(r *http.Request) string {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("Body error " + err.Error())
	}

	var res bytes.Buffer
	if err := template.Must(template.New("response").Parse(responseTemplate)).Execute(&res, Response{
		URL:       r.URL.String(),
		Method:    r.Method,
		Protocol:  r.Proto,
		Headers:   joinHeaders(r.Header),
		Accepts:   r.Header.Get("Accept"),
		UserAgent: r.Header.Get("User-Agent"),
		Body:      string(b),
		Form: joinForm(r.Form),
		Form2: joinForm(r.PostForm),
	}); err != nil {
		panic(err)
	}

	return res.String()
}

func joinForm(form url.Values) string {
	if form == nil {
		return ""
	}
	var res []string
	for k := range form {
		res = append(res, "["+k+"="+form.Get(k)+"]")
	}
	return strings.Join(res, ",")
}

func joinHeaders(header http.Header) string {
	var res = make([]string, 0, len(header))
	for k := range header {
		res = append(res, "["+k+"="+header.Get(k)+"]")
	}
	return strings.Join(res, ",")
}

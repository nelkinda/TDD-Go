package main

import (
	"encoding/json"
	"flag"
	"github.com/nelkinda/http-go/header"
	"github.com/nelkinda/http-go/https"
	"github.com/nelkinda/http-go/mimetype"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	httpPortPtr := flag.String("http-port", ":http", "Port for http")
	flag.Parse()

	mux := createMux()

	https.MustServeHttpPort(*httpPortPtr, mux)
	https.WaitForIntOrTerm()
	os.Exit(0)
}

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)
	return mux
}

type LoginRequest struct {
	Username string
	Password string
}

func login(writer http.ResponseWriter, request *http.Request) {
	method := request.Method
	if method == http.MethodPost {
		reqBody, err := ioutil.ReadAll(request.Body)
		if err != nil {
			// TODO
		}
		var newLoginRequest LoginRequest
		json.Unmarshal(reqBody, &newLoginRequest)
		// TODO
	} else {
		writer.Header().Add(header.ContentType, mimetype.TextPlain)
		writer.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = writer.Write([]byte("405 Method Not Allowed\nThe method " + r.Method + " is not allowed for this resource."))
	}
}



package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	s "github.com/parikshitg/add-microservice/service"
	h "github.com/urantiatech/kit/transport/http"
)

func main() {

	var port int
	flag.IntVar(&port, "port", 8080, "Port number")
	flag.Parse()

	log.SetFlags(log.Lshortfile)

	if os.Getenv("PORT") != "" {
		p, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
		if err != nil {
			port = int(p)
		}
	}

	var svc = s.Add{}

	r := mux.NewRouter()

	r.Handle("/add", h.NewServer(s.MakeAddEndpoint(svc), s.DecodeAddRequest, s.EncodeResponse))

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)

}

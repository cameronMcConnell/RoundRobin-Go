package lib

import (
	"net/http"
	"sync"
)

type HttpServer struct {
	Address string
}

func newHttpServer(address string) HttpServer {
	return HttpServer{Address: address}
}

func bindServerRoute() {
	http.HandleFunc("/roundrobin-go", func(w http.ResponseWriter, r *http.Request) {
		go roundRobin(w, r) 
	})
}

func roundRobin(w http.ResponseWriter, r *http.Request) {

}

func (h *HttpServer) startServer() error {
	err := http.ListenAndServe(h.Address, nil)
	if err != nil {
		return err
	}

	return nil
}


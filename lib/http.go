package lib

import (
	"io"
	"net/http"
	"sync"
	"fmt"
)

type RoundRobinServer struct {
	Address string
	Servers Queue
}

func NewRoundRobinServer(address string) (*RoundRobinServer, error) {
	servers, err := NewQueue()
	if err != nil {
		return nil, err
	}
	
	return &RoundRobinServer{Address: address, Servers: *servers}, nil
}

func (rrs *RoundRobinServer) roundRobin(w http.ResponseWriter, r *http.Request) {
	var mutex sync.Mutex
	mutex.Lock()

	server, err := rrs.Servers.Deque()
	if err != nil {
		mutex.Unlock()
		http.Error(w, "No available servers", http.StatusServiceUnavailable)
		return
	}

	rrs.Servers.Enque(server)

	mutex.Unlock()

	forwardReq, err := http.NewRequest(r.Method, server, r.Body)
	if err != nil {
		http.Error(w, "Error creating forwarding request", http.StatusInternalServerError)
		return
	}
	
	forwardReq.Header = r.Header

	client := &http.Client{}
	resp, err := client.Do(forwardReq)
	if err != nil {
		http.Error(w, "Error forwarding request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

func (rrs *RoundRobinServer) StartServer() error {
	http.HandleFunc("/roundrobin-go", func(w http.ResponseWriter, r *http.Request) { 
		go rrs.roundRobin(w, r)
	})

	fmt.Printf("Listening on address: %s\n", rrs.Address)

	err := http.ListenAndServe(rrs.Address, nil)
	if err != nil {
		return err
	}

	return nil
}
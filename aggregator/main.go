package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/Jimbo8702/toll-calc/types"
)

func main() {
	listenAddr := flag.String("listenAdder", ":3000", "the listen address of the http server")
	flag.Parse()
	var (
		store 	= NewMemoryStore()
		svc 	= NewInvoiceAggregator(store)
	)
	makeHTTPTransport(*listenAddr, svc)
}

func makeHTTPTransport(listenAddr string, svc Aggregator) {
	fmt.Println("http transport running on port:", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.ListenAndServe(listenAddr, nil)
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println(distance)
	}
}
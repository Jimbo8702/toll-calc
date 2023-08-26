package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/Jimbo8702/toll-calc/types"
)

func main() {
	listenAddr := flag.String("listenAdder", ":4000", "the listen address of the http server")
	flag.Parse()
	var (
		svc Aggregator
		store = NewMemoryStore()
	)
	svc = NewInvoiceAggregator(store)
	svc = NewLogMiddleware(svc)
	
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
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		if err := svc.AggregateDistance(distance); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/Jimbo8702/toll-calc/aggregator/client"
	"github.com/sirupsen/logrus"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error 

func main() {
	listenAddr := flag.String("listenAddr", ":6000", "")
	flag.Parse()
	var (
		client 	   = client.NewHTTPClient("http://localhost:3000")
		invHandler = NewInvoiceHandler(client)
	)
	http.HandleFunc("/invoice", makeAPIFunc(invHandler.handleGetInvoice))
	logrus.Infof("gateway HTTP server running on port %s", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

type InvoiceHandler struct {
	client client.Client
}

func NewInvoiceHandler(client client.Client) *InvoiceHandler {
	return &InvoiceHandler{
		client: client,
	}
}

func (h *InvoiceHandler) handleGetInvoice(w http.ResponseWriter, r *http.Request) error {
	inv, err := h.client.GetInvoice(context.Background(), 1)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, inv)
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("Context-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func makeAPIFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}
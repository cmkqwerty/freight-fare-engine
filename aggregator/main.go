package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/cmkqwerty/freight-fare-engine/types"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"strconv"
)

func main() {
	httpListenAddr := flag.String("httpAddr", ":3000", "HTTP server listen address")
	grpcListenAddr := flag.String("grpcAddr", ":3001", "GRPC server listen address")
	flag.Parse()

	var (
		store = NewMemoryStore()
		svc   = NewInvoiceAggregator(store)
	)
	svc = NewLogMiddleware(svc)

	go makeGRPCTransport(*grpcListenAddr, svc)
	makeHTTPTransport(*httpListenAddr, svc)
}

func makeGRPCTransport(listenAddr string, svc Aggregator) error {
	fmt.Println("HTTP Transport running on", listenAddr)

	ln, err := net.Listen("TCP", listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	server := grpc.NewServer([]grpc.ServerOption{}...)

	types.RegisterAggregatorServer(server, NewAggregatorGRPCServer(svc))

	return server.Serve(ln)
}

func makeHTTPTransport(listenAddr string, svc Aggregator) {
	fmt.Println("HTTP Transport running on", listenAddr)

	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.HandleFunc("/invoice", handleGetInvoice(svc)) //invoice?id=OBU_ID
	http.ListenAndServe(listenAddr, nil)
}

func handleGetInvoice(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obuID, err := strconv.Atoi(r.URL.Query().Get("obu"))
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid obu id"})
			return
		}

		invoice, err := svc.CalculateInvoice(obuID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, invoice)
	}
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance

		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
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

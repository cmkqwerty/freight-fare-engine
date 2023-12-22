package main

import (
	"encoding/json"
	"fmt"
	"github.com/cmkqwerty/freight-fare-engine/types"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var (
		store          = makeStore()
		svc            = NewInvoiceAggregator(store)
		grpcListenAddr = os.Getenv("AGGREGATE_GRPC_ENDPOINT")
		httpListenAddr = os.Getenv("AGGREGATE_HTTP_ENDPOINT")
	)
	svc = NewMetricsMiddleware(svc)
	svc = NewLogMiddleware(svc)

	go func() {
		log.Fatal(makeGRPCTransport(grpcListenAddr, svc))
	}()
	log.Fatal(makeHTTPTransport(httpListenAddr, svc))
}

func makeGRPCTransport(listenAddr string, svc Aggregator) error {
	fmt.Println("GRPC Transport running on", listenAddr)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	server := grpc.NewServer([]grpc.ServerOption{}...)

	types.RegisterAggregatorServer(server, NewAggregatorGRPCServer(svc))

	return server.Serve(ln)
}

func makeHTTPTransport(listenAddr string, svc Aggregator) error {
	fmt.Println("HTTP Transport running on", listenAddr)

	var (
		mhAggregate      = newHTTPMetricHandler("aggregate")
		mhInvoice        = newHTTPMetricHandler("invoice")
		aggregateHandler = makeHTTPHandlerFunc(mhAggregate.instrument(handleAggregate(svc)))
		invoiceHandler   = makeHTTPHandlerFunc(mhInvoice.instrument(handleGetInvoice(svc)))
	)
	http.HandleFunc("/aggregate", aggregateHandler)
	http.HandleFunc("/invoice", invoiceHandler)
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(listenAddr, nil)
}

func makeStore() Storer {
	t := os.Getenv("AGGREGATE_STORE_TYPE")
	switch t {
	case "memory":
		return NewMemoryStore()
	default:
		log.Fatal("invalid store type")
		return nil
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}

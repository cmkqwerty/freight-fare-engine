package main

import (
	"encoding/json"
	"fmt"
	"github.com/cmkqwerty/freight-fare-engine/aggregator/client"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"time"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	listenAddr := os.Getenv("GATEWAY_SERVER_ENDPOINT")
	aggregatorServiceAddr := fmt.Sprintf("http://localhost%s", os.Getenv("AGGREGATE_HTTP_ENDPOINT"))

	var (
		newClient  = client.NewHTTPClient(aggregatorServiceAddr)
		invHandler = NewInvoiceHandler(newClient)
	)

	http.HandleFunc("/invoice", makeAPIFunc(invHandler.handleGetInvoice))

	logrus.Infof("HTTP gateway server started on %s", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
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
	invoice, err := h.client.GetInvoice(r.Context(), 192928971)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, invoice)
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(v)
}

func makeAPIFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(start time.Time) {
			logrus.WithFields(logrus.Fields{
				"took": time.Since(start),
				"uri":  r.RequestURI,
			}).Info("API req :: ")
		}(time.Now())

		if err := fn(w, r); err != nil {
			logrus.Errorf("API error: %s", err)
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

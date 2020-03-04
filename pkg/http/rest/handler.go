package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mronfim/invoicify/pkg/billing"
)

func Handler(invoiceHandler billing.InvoiceHandler) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/invoices", invoiceHandler.Get).Methods("GET")
	router.HandleFunc("/invoices", invoiceHandler.Create).Methods("POST")
	router.HandleFunc("/invoices/{id}", invoiceHandler.GetById).Methods("GET")

	http.Handle("/", router)

	return router
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin:", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

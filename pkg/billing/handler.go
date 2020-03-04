package billing

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type InvoiceHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type invoiceHandler struct {
	invoiceService InvoiceService
}

func NewInvoiceHandler(invoiceService InvoiceService) InvoiceHandler {
	return &invoiceHandler{
		invoiceService,
	}
}

func (h *invoiceHandler) Get(w http.ResponseWriter, r *http.Request) {
	invoices, _ := h.invoiceService.FindAllInvoices()

	response, _ := json.Marshal(invoices)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func (h *invoiceHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	invoice, _ := h.invoiceService.FindInvoiceById(id)

	response, _ := json.Marshal(invoice)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func (h *invoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var invoice Invoice
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&invoice)
	_ = h.invoiceService.CreateInvoice(&invoice)

	response, _ := json.Marshal(invoice)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

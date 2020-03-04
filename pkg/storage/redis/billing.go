package redis

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/mronfim/invoicify/pkg/billing"
)

const invoiceTable = "invoices"

type invoiceRepository struct {
	connection *redis.Client
	table      string
}

func NewRedisInvoiceRepository(connection *redis.Client) billing.InvoiceRepository {
	return &invoiceRepository{
		connection,
		"invoices",
	}
}

func (r *invoiceRepository) Create(invoice *billing.Invoice) error {
	encoded, err := json.Marshal(invoice)

	if err != nil {
		return err
	}

	r.connection.HSet(r.table, invoice.ID, encoded)
	return nil
}

func (r *invoiceRepository) FindById(id string) (*billing.Invoice, error) {
	b, err := r.connection.HGet(r.table, id).Bytes()

	if err != nil {
		return nil, err
	}

	invoice := new(billing.Invoice)
	err = json.Unmarshal(b, invoice)

	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (r *invoiceRepository) FindAll() (invoices []*billing.Invoice, err error) {
	all := r.connection.HGetAll(r.table).Val()

	for key, value := range all {
		invoice := new(billing.Invoice)
		err = json.Unmarshal([]byte(value), invoice)

		if err != nil {
			return nil, err
		}

		invoice.ID = key
		invoices = append(invoices, invoice)
	}

	return invoices, nil
}

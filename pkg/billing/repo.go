package billing

type InvoiceRepository interface {
	Create(invoice *Invoice) error
	FindById(id string) (*Invoice, error)
	FindAll() ([]*Invoice, error)
}

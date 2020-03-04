package billing

import (
	"time"

	"github.com/google/uuid"
)

type InvoiceService interface {
	CreateInvoice(invoice *Invoice) error
	FindInvoiceById(id string) (*Invoice, error)
	FindAllInvoices() ([]*Invoice, error)
}

type invoiceService struct {
	repo InvoiceRepository
}

func NewInvoiceService(repo InvoiceRepository) InvoiceService {
	return &invoiceService{
		repo,
	}
}

func (s *invoiceService) CreateInvoice(invoice *Invoice) error {
	invoice.ID = uuid.New().String()
	invoice.CreatedAt = time.Now()
	invoice.UpdatedAt = time.Now()
	invoice.DeletedAt = time.Time{}
	return s.repo.Create(invoice)
}

func (s *invoiceService) FindInvoiceById(id string) (*Invoice, error) {
	return s.repo.FindById(id)
}

func (s *invoiceService) FindAllInvoices() ([]*Invoice, error) {
	return s.repo.FindAll()
}

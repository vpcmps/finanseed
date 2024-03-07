package transaction

import "time"

type CreateTransactionRequest struct {
	Amount        float64         `json:"amount"`
	Type          TransactionType `json:"type"`
	Name          string          `json:"name"`
	ReferenceDate time.Time       `json:"referenceDate"`
}

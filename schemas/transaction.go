package schemas

import (
	"time"

	"gorm.io/gorm"
)

type TransactionType int

const (
	In TransactionType = iota
	Out
)

type Transaction struct {
	gorm.Model
	Amount        float64         `json:"amount"`
	Type          TransactionType `json:"type"`
	Name          string          `json:"name"`
	ReferenceDate time.Time       `json:"referenceDate"`
}

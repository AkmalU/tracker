package transactions

import (
	"encoding/json"
	"time"

	"github.com/AkmalUr/test/amounts"
)

// Transaction ...
type Transaction struct {
	ID         string          `json:"id"`
	CustomerID string          `json:"customer_id"`
	LoadAmount *amounts.Amount `json:"load_amount"`
	Time       time.Time       `json:"time"`
}

// TransactionResult ...
type TransactionResult struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}

// UnmarshalJSON converts transaction to it's JSON representation
func (t *Transaction) UnmarshalJSON(b []byte) error {
	var transaction map[string]string
	err := json.Unmarshal(b, &transaction)
	if err != nil {
		return err
	}
	layout := "2006-01-02T15:04:05Z"
	for key, value := range transaction {
		var err error
		switch key {
		case "id":
			t.ID = value
		case "customer_id":
			t.CustomerID = value
		case "time":
			t.Time, err = time.Parse(layout, value)
		case "load_amount":
			t.LoadAmount = amounts.Parse(value)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// SameAs compares if two transaction results are of the same value
func (tr *TransactionResult) SameAs(other *TransactionResult) bool {
	result := true
	result = result && tr.Accepted == other.Accepted
	result = result && tr.CustomerID == other.CustomerID
	result = result && tr.ID == other.ID
	return result
}

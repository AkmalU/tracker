package transactions_test

import (
	"testing"
	"time"

	"github.com/AkmalUr/test/amounts"
	"github.com/AkmalUr/test/transactions"
)

type testTransaction struct {
	transaction *transactions.Transaction
	result      *transactions.TransactionResult
}

var layout = "2006-01-02T15:04:05Z"
var testTransactionTable = []testTransaction{
	testTransaction{
		&transactions.Transaction{
			ID:         "1",
			CustomerID: "1",
			LoadAmount: amounts.Parse("$10"),
			Time:       time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		&transactions.TransactionResult{
			ID:         "1",
			CustomerID: "1",
			Accepted:   true,
		},
	},
	testTransaction{
		&transactions.Transaction{
			ID:         "2",
			CustomerID: "1",
			LoadAmount: amounts.Parse("$10"),
			Time:       time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		&transactions.TransactionResult{
			ID:         "2",
			CustomerID: "1",
			Accepted:   true,
		},
	},
	testTransaction{
		&transactions.Transaction{
			ID:         "3",
			CustomerID: "1",
			LoadAmount: amounts.Parse("$10"),
			Time:       time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		&transactions.TransactionResult{
			ID:         "3",
			CustomerID: "1",
			Accepted:   true,
		},
	},
	testTransaction{
		&transactions.Transaction{
			ID:         "4",
			CustomerID: "1",
			LoadAmount: amounts.Parse("$10"),
			Time:       time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		&transactions.TransactionResult{
			ID:         "4",
			CustomerID: "1",
			Accepted:   false,
		},
	},
	testTransaction{
		&transactions.Transaction{
			ID:         "5",
			CustomerID: "1",
			LoadAmount: amounts.Parse("$3000"),
			Time:       time.Date(2000, time.January, 2, 0, 0, 0, 0, time.UTC),
		},
		&transactions.TransactionResult{
			ID:         "5",
			CustomerID: "1",
			Accepted:   true,
		},
	},
	testTransaction{
		&transactions.Transaction{
			ID:         "6",
			CustomerID: "1",
			LoadAmount: amounts.Parse("$3000"),
			Time:       time.Date(2000, time.January, 2, 0, 0, 0, 0, time.UTC),
		},
		&transactions.TransactionResult{
			ID:         "6",
			CustomerID: "1",
			Accepted:   false,
		},
	},
	testTransaction{
		&transactions.Transaction{
			ID:         "7",
			CustomerID: "1",
			LoadAmount: amounts.Parse("$3000"),
			Time:       time.Date(2000, time.January, 3, 0, 0, 0, 0, time.UTC),
		},
		&transactions.TransactionResult{
			ID:         "7",
			CustomerID: "1",
			Accepted:   true,
		},
	},
	testTransaction{
		&transactions.Transaction{
			ID:         "8",
			CustomerID: "1",
			LoadAmount: amounts.Parse("$13000"),
			Time:       time.Date(2000, time.January, 4, 0, 0, 0, 0, time.UTC),
		},
		&transactions.TransactionResult{
			ID:         "8",
			CustomerID: "1",
			Accepted:   false,
		},
	},
	testTransaction{
		&transactions.Transaction{
			ID:         "9",
			CustomerID: "1",
			LoadAmount: amounts.Parse("$3000"),
			Time:       time.Date(2000, time.January, 8, 0, 0, 0, 0, time.UTC),
		},
		&transactions.TransactionResult{
			ID:         "9",
			CustomerID: "1",
			Accepted:   true,
		},
	},
}

func TestProcessTransaction(t *testing.T) {
	ua := &transactions.UserActivity{
		CustomerID:        "1",
		DailyAmount:       amounts.NewAmount(),
		WeeklyAmount:      amounts.NewAmount(),
		DailyTransactions: 0,
		Ledger:            map[string]*transactions.Transaction{},
	}

	for _, testTransaction := range testTransactionTable {
		tResult, _ := ua.AddTransaction(testTransaction.transaction)
		if !tResult.SameAs(testTransaction.result) {
			t.Errorf(
				"Failed processing transaction:\n\tExpected:%+v\n\tActual:%+v\n",
				*testTransaction.result,
				*tResult,
			)
		}
	}
}

func TestProcessDuplicateTransaction(t *testing.T) {
	ua := &transactions.UserActivity{
		CustomerID:        "1",
		DailyAmount:       amounts.NewAmount(),
		WeeklyAmount:      amounts.NewAmount(),
		DailyTransactions: 0,
		Ledger:            map[string]*transactions.Transaction{},
	}

	tt := testTransaction{
		&transactions.Transaction{
			ID:         "9",
			CustomerID: "1",
			LoadAmount: amounts.Parse("$3000"),
			Time:       time.Date(2000, time.January, 8, 0, 0, 0, 0, time.UTC),
		},
		&transactions.TransactionResult{
			ID:         "9",
			CustomerID: "1",
			Accepted:   true,
		},
	}

	ua.AddTransaction(tt.transaction)
	_, err := ua.AddTransaction(tt.transaction)

	if err == nil {
		t.Errorf("Failed to detect duplicate transaction")
	}
}

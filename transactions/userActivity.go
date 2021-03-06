package transactions

import (
	"fmt"
	"time"

	"github.com/AkmalUr/test/amounts"
)

var _MaxDailyAmount = amounts.Parse("$5000")
var _MaxWeeklyAmount = amounts.Parse("$20000")
var _MaxNumberOfDailyTransactions = 3

// UserActivity ...
type UserActivity struct {
	CustomerID        string
	Updated           time.Time
	DailyAmount       *amounts.Amount
	WeeklyAmount      *amounts.Amount
	DailyTransactions int
	Ledger            map[string]*Transaction
}

// AddTransaction adds a new transaction to user's activity
func (ua *UserActivity) AddTransaction(t *Transaction) (*TransactionResult, error) {
	result := &TransactionResult{
		ID:         t.ID,
		CustomerID: t.CustomerID,
		Accepted:   false,
	}

	if _, ok := ua.Ledger[t.ID]; ok {
		return nil, fmt.Errorf("Duplicate transaction: %s", t.ID)
	}
	ua.Ledger[t.ID] = t
	if !isSameDay(&ua.Updated, &t.Time) {
		ua.DailyAmount = amounts.Parse("$0")
		ua.DailyTransactions = 0
	}
	if !isSameWeek(&ua.Updated, &t.Time) {
		ua.WeeklyAmount = amounts.Parse("$0")
	}

	if ua.DailyTransactions == _MaxNumberOfDailyTransactions {
		return result, nil
	}
	if ua.DailyAmount.Sum(t.LoadAmount).Compare(_MaxDailyAmount) == 1 {
		return result, nil
	}
	if ua.WeeklyAmount.Sum(t.LoadAmount).Compare(_MaxWeeklyAmount) == 1 {
		return result, nil
	}
	ua.Updated = t.Time
	ua.DailyAmount = ua.DailyAmount.Add(t.LoadAmount)
	ua.WeeklyAmount = ua.WeeklyAmount.Add(t.LoadAmount)
	ua.DailyTransactions++

	result.Accepted = true
	return result, nil
}

func isSameDay(this *time.Time, other *time.Time) bool {
	result := true
	result = result && this.Year() == other.Year()
	result = result && this.Month() == other.Month()
	result = result && this.Day() == other.Day()
	return result
}

func isSameWeek(this *time.Time, other *time.Time) bool {
	result := true
	thisYear, thisWeek := this.ISOWeek()
	otherYear, otherWeek := other.ISOWeek()

	result = result && thisYear == otherYear
	result = result && thisWeek == otherWeek
	return result
}

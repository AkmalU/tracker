package transactions

import (
	"github.com/AkmalUr/test/amounts"
)

// UserActivityRepository provides an interface for interacting with
// user activities storage
type UserActivityRepository interface {
	Save(ua *UserActivity)
	Get(id string) *UserActivity
}

// InMemoryRepository stores all user activities in-memory
type InMemoryRepository struct {
	store map[string]*UserActivity
}

// NewInMemoryRepo returns a pointer to a new user activities repository
func NewInMemoryRepo() *InMemoryRepository {
	return &InMemoryRepository{map[string]*UserActivity{}}
}

// Save ...
func (r *InMemoryRepository) Save(ua *UserActivity) {
	r.store[ua.CustomerID] = ua
}

// Get ...
func (r *InMemoryRepository) Get(customerID string) *UserActivity {
	if _, ok := r.store[customerID]; ok {
		return r.store[customerID]
	}

	return &UserActivity{
		CustomerID:        customerID,
		DailyAmount:       amounts.NewAmount(),
		WeeklyAmount:      amounts.NewAmount(),
		DailyTransactions: 0,
		Ledger:            map[string]*Transaction{},
	}
}

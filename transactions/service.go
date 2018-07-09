package transactions

import "fmt"

// UserActivityService process all user's activity
type UserActivityService struct {
	repository UserActivityRepository
}

// NewService returns a pointer to a new instance of user activity service
func NewService(repo UserActivityRepository) *UserActivityService {
	return &UserActivityService{repo}
}

// ProcessTransaction ...
func (s *UserActivityService) ProcessTransaction(t *Transaction) *TransactionResult {
	userActivity := s.repository.Get(t.CustomerID)
	result, err := userActivity.AddTransaction(t)
	if err != nil {
		fmt.Printf("Failed to add transaction. Reason: %s", err)
	}
	if result.Accepted {
		s.repository.Save(userActivity)
	}
	return result
}

package amounts

import (
	"fmt"
	"math/big"
)

// Amount represents a monetary amount
type Amount struct {
	Raw      string
	Value    *big.Rat
	Currency string
}

// NewAmount returns a pointer to a new amount
func NewAmount() (a *Amount) {
	a = new(Amount)
	a.Value = new(big.Rat)
	return a
}

// Parse converts string representation to an amount
func Parse(amount string) *Amount {
	a := NewAmount()
	a.Raw = amount
	a.Value.SetString(a.parseValue())
	a.Currency = a.parseCurrency()

	return a
}

func (a *Amount) parseValue() string {
	v := a.Raw[1:]
	if len(v) == 0 {
		return "0"
	}

	return v
}

func (a *Amount) parseCurrency() string {
	if len(a.Raw) == 0 {
		return "USD"
	}
	if a.Raw[:1] == "$" {
		return "USD"
	}
	return "N/A"
}

func (a *Amount) toString() string {
	amount := a.Value.FloatString(2)
	if amount[len(amount)-2:] == "00" {
		return fmt.Sprintf("$%s", amount[:len(amount)-3])
	}
	return fmt.Sprintf("$%s", amount)
}

// Add the value of other amount to this
func (a *Amount) Add(other *Amount) *Amount {
	a.Value.Add(a.Value, other.Value)
	a.Raw = a.toString()
	return a
}

// Compare the values of two amounts
func (a *Amount) Compare(other *Amount) int {
	return a.Value.Cmp(other.Value)
}

func (a *Amount) String() string {
	return a.Raw
}

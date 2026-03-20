package domain

import "fmt"

// Wallet represents a user's wallet with sats (satoshis - like coins!)
// Think of this as your piggy bank that holds your money
type Wallet struct {
	Balance int `json:"balance"` // How many sats you have (whole numbers only!)
}

// NewWallet creates a new empty wallet
// It's like getting a brand new piggy bank!
func NewWallet() *Wallet {
	return &Wallet{
		Balance: 0, // Start with 0 sats
	}
}

// Fund adds 155 sats to the wallet
// This is like your mom giving you money!
// Rules: Can only fund if balance is 35 or less
func (w *Wallet) Fund() error {
	// Check if wallet already has too much money
	if w.Balance > 35 {
		return ErrWalletHasEnoughFunds
	}

	// Add exactly 155 sats (no more, no less!)
	w.Balance = 155

	return nil
}

// Debit takes sats away from the wallet
// This is like spending your money
// amount = how many sats to take away
func (w *Wallet) Debit(amount int) error {
	// Check if you have enough money
	if w.Balance < amount {
		return ErrInsufficientFunds
	}

	// Take away the money
	w.Balance -= amount

	return nil
}

// Credit adds sats to the wallet
// This is like receiving prize money!
// amount = how many sats to add
func (w *Wallet) Credit(amount int) {
	w.Balance += amount
}

// GetBalance returns how many sats you have
// This is like counting your coins!
func (w *Wallet) GetBalance() int {
	return w.Balance
}

// GetBalanceFormatted returns balance as a string with "sats" label
// Example: "155 sats"
// This is for displaying to users in a nice way!
func (w *Wallet) GetBalanceFormatted() string {
	return fmt.Sprintf("%d sats", w.Balance)
}

// Import fmt for formatting

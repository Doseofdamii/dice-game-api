package services

import (
	"dice-game-api/internal/domain"
	"fmt"
)

// GameService handles all the game logic
// This is the BRAIN of the game!
// It knows all the rules and enforces them
type GameService struct {
	wallet *domain.Wallet // The player's wallet
	game   *domain.Game   // The current game state
}

// NewGameService creates a new game service
// Call this once when starting the app
func NewGameService() *GameService {
	return &GameService{
		wallet: domain.NewWallet(), // Create empty wallet
		game:   &domain.Game{},     // Create inactive game
	}
}

// =====================================
// WALLET OPERATIONS
// =====================================

// FundWallet adds 155 sats to the wallet
// Rules: Can only fund if balance is 35 or less
func (s *GameService) FundWallet() error {
	// Check if wallet can be funded
	if s.wallet.Balance > domain.MinBalanceForFunding {
		return domain.ErrWalletHasEnoughFunds
	}
	
	// Set balance to exactly 155 sats
	s.wallet.Balance = domain.FundAmount
	
	return nil
}

// GetBalance returns the current wallet balance
func (s *GameService) GetBalance() int {
	return s.wallet.GetBalance()
}

// GetBalanceInfo returns detailed balance information
// Returns: balance (int), balance string, asset name
func (s *GameService) GetBalanceInfo() (int, string, string) {
	balance := s.wallet.GetBalance()
	balanceStr := fmt.Sprintf("%d", balance)
	asset := "sats"
	
	return balance, balanceStr, asset
}

// =====================================
// GAME OPERATIONS
// =====================================

// StartGame starts a new game session
// Costs: 20 sats
// Rules:
// - Must have at least 20 sats
// - Cannot start if game already active
func (s *GameService) StartGame() (*domain.Game, error) {
	// Rule 1: Check if game is already active
	if s.game.IsActive() {
		return nil, domain.ErrGameAlreadyActive
	}
	
	// Rule 2: Check if user has enough money (20 sats)
	if s.wallet.Balance < domain.GameStartCost {
		return nil, domain.ErrInsufficientFunds
	}
	
	// Deduct 20 sats from wallet
	err := s.wallet.Debit(domain.GameStartCost)
	if err != nil {
		return nil, err
	}
	
	// Create new game with random target number (2-12)
	s.game = domain.NewGame()
	
	return s.game, nil
}

// EndGame closes the current game session
// Rules: Can only end if game is active
func (s *GameService) EndGame() error {
	// Check if game is active
	if !s.game.IsActive() {
		return domain.ErrNoActiveGame
	}
	
	// End the game
	s.game.End()
	
	return nil
}

// IsGameActive checks if a game is currently in session
func (s *GameService) IsGameActive() bool {
	return s.game.IsActive()
}

// GetGameStatus returns the current game status
func (s *GameService) GetGameStatus() map[string]interface{} {
	return map[string]interface{}{
		"active":        s.game.IsActive(),
		"has_first_roll": !s.game.NeedsFirstRoll(),
	}
}

// =====================================
// DICE ROLLING OPERATIONS
// =====================================

// RollResult represents the result of a dice roll
type RollResult struct {
	Roll         int    `json:"roll"`           // The number you rolled (1-6)
	IsFirstRoll  bool   `json:"is_first_roll"`  // Is this the first roll of the pair?
	IsSecondRoll bool   `json:"is_second_roll"` // Is this the second roll of the pair?
	FirstRoll    int    `json:"first_roll,omitempty"`     // What was the first roll? (only shown on second roll)
	SecondRoll   int    `json:"second_roll,omitempty"`    // What was the second roll? (only shown on second roll)
	Sum          int    `json:"sum,omitempty"`            // Sum of both rolls (only on second roll)
	Target       int    `json:"target,omitempty"`         // The target number (only shown on second roll)
	Won          bool   `json:"won,omitempty"`            // Did you win? (only on second roll)
	Prize        int    `json:"prize,omitempty"`          // Prize amount if won (only if won)
	Balance      int    `json:"balance"`                  // Your new balance after this roll
	Message      string `json:"message"`                  // Helpful message
}

// RollDice rolls a single die
// This handles BOTH first and second rolls of a pair!
//
// How it works:
// 1. First roll: Costs 5 sats, saves the result
// 2. Second roll: Free! Checks if you won
func (s *GameService) RollDice() (*RollResult, error) {
	// Rule 1: Game must be active
	if !s.game.IsActive() {
		return nil, domain.ErrNoActiveGame
	}
	
	// Check if this is the first or second roll of the pair
	if s.game.NeedsFirstRoll() {
		// This is the FIRST ROLL of the pair
		return s.handleFirstRoll()
	} else {
		// This is the SECOND ROLL of the pair
		return s.handleSecondRoll()
	}
}

// handleFirstRoll handles the first die roll of a pair
// Costs: 5 sats
func (s *GameService) handleFirstRoll() (*RollResult, error) {
	// Check if user has enough money (5 sats)
	if s.wallet.Balance < domain.DiceRollCost {
		return nil, domain.ErrInsufficientFunds
	}
	
	// Charge 5 sats for the roll pair
	err := s.wallet.Debit(domain.DiceRollCost)
	if err != nil {
		return nil, err
	}
	
	// Roll the die! (1-6)
	roll := domain.RollDice()
	
	// Save the first roll
	s.game.SetFirstRoll(roll)
	
	// Return the result
	return &RollResult{
		Roll:        roll,
		IsFirstRoll: true,
		IsSecondRoll: false,
		Balance:     s.wallet.GetBalance(),
		Message:     fmt.Sprintf("First roll: %d. Roll again to complete the pair!", roll),
	}, nil
}

// handleSecondRoll handles the second die roll of a pair
// Cost: FREE! (already paid on first roll)
func (s *GameService) handleSecondRoll() (*RollResult, error) {
	// Roll the die! (1-6)
	roll := domain.RollDice()
	
	// Get the first roll we saved earlier
	firstRoll := s.game.GetFirstRoll()
	
	// Calculate the sum
	sum := domain.CalculateSum(firstRoll, roll)
	
	// Get the target number
	target := s.game.GetTargetNumber()
	
	// Check if won!
	won := domain.CheckWin(firstRoll, roll, target)
	
	result := &RollResult{
		Roll:         roll,
		IsFirstRoll:  false,
		IsSecondRoll: true,
		FirstRoll:    firstRoll,
		SecondRoll:   roll,
		Sum:          sum,
		Target:       target,
		Won:          won,
		Balance:      s.wallet.GetBalance(),
	}
	
	// If won, give prize!
	if won {
		s.wallet.Credit(domain.WinPrize)
		result.Prize = domain.WinPrize
		result.Balance = s.wallet.GetBalance()
		result.Message = fmt.Sprintf("🎉 YOU WIN! %d + %d = %d (target: %d). Won %d sats!", firstRoll, roll, sum, target, domain.WinPrize)
	} else {
		result.Message = fmt.Sprintf("You rolled %d + %d = %d (target: %d). Try again!", firstRoll, roll, sum, target)
	}
	
	// Reset for next roll pair
	s.game.ResetRollPair()
	
	return result, nil
}

// =====================================
// HELPER FUNCTIONS
// =====================================

// GetWallet returns the wallet (for testing)
func (s *GameService) GetWallet() *domain.Wallet {
	return s.wallet
}

// GetGame returns the game (for testing)
func (s *GameService) GetGame() *domain.Game {
	return s.game
}

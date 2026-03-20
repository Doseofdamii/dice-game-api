package domain

import (
	"errors"
	"math/rand"
	"time"
)

// Initialize random number generator
// This makes sure we get different random numbers each time!
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RollDice simulates rolling a single die (1-6)
// Just like rolling a real die!
// Returns a number between 1 and 6
func RollDice() int {
	// rand.Intn(6) gives 0-5, +1 makes it 1-6
	return rand.Intn(6) + 1
}

// CheckWin checks if the sum of two rolls equals the target
// firstRoll + secondRoll = target? You win!
func CheckWin(firstRoll, secondRoll, target int) bool {
	sum := firstRoll + secondRoll
	return sum == target
}

// CalculateSum adds two die rolls together
// Simple addition: 3 + 4 = 7
func CalculateSum(firstRoll, secondRoll int) int {
	return firstRoll + secondRoll
}

// =====================================
// ERROR DEFINITIONS
// =====================================
// These are all the things that can go wrong!
// We define them here so we can reuse them everywhere

// Wallet Errors
var (
	// ErrInsufficientFunds means you don't have enough sats
	// Like trying to buy something but your wallet is empty!
	ErrInsufficientFunds = errors.New("insufficient funds")
	
	// ErrWalletHasEnoughFunds means you already have enough money
	// Can't add more funds until you spend some!
	ErrWalletHasEnoughFunds = errors.New("wallet has enough funds, cannot fund until balance is 35 sats or less")
)

// Game Errors
var (
	// ErrGameAlreadyActive means a game is already running
	// You can't start two games at once!
	ErrGameAlreadyActive = errors.New("a game is already in session")
	
	// ErrNoActiveGame means no game is running
	// You need to start a game first!
	ErrNoActiveGame = errors.New("no game is currently active")
	
	// ErrGameNotActive is same as ErrNoActiveGame
	// (Just another way to say it)
	ErrGameNotActive = errors.New("game is not active")
)

// Constants for game costs and prizes
const (
	// GAME_START_COST is how many sats it costs to start a game
	GameStartCost = 20 // Starting a game costs 20 sats
	
	// DICE_ROLL_COST is how many sats it costs to roll a pair of dice
	DiceRollCost = 5 // Each pair of rolls costs 5 sats (charged on first roll)
	
	// WIN_PRIZE is how many sats you win if you match the target
	WinPrize = 10 // Winning gives you 10 sats!
	
	// FUND_AMOUNT is how many sats you get when funding wallet
	FundAmount = 155 // You always get exactly 155 sats
	
	// MIN_BALANCE_FOR_FUNDING is the maximum balance to allow funding
	MinBalanceForFunding = 35 // Can only fund if balance <= 35 sats
)

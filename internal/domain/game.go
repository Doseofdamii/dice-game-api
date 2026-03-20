package domain

import "math/rand"

// Game represents an active dice game session
// Think of this as the game board that tracks everything!
type Game struct {
	Active       bool `json:"active"`        // Is the game running? true/false
	TargetNumber int  `json:"target_number"` // The secret number you need to match (2-12)
	FirstRoll    int  `json:"first_roll"`    // What you rolled first (0 if haven't rolled yet)
	HasFirstRoll bool `json:"has_first_roll"` // Did you do the first roll? true/false
}

// NewGame creates a new game with a random target number
// It's like setting up a new round at the carnival!
// 
// The target number is between 2 and 12 because:
// - Minimum: 1 + 1 = 2 (rolling two 1s)
// - Maximum: 6 + 6 = 12 (rolling two 6s)
func NewGame() *Game {
	// Generate random number between 2 and 12
	targetNumber := rand.Intn(11) + 2 // rand.Intn(11) gives 0-10, +2 makes it 2-12
	
	return &Game{
		Active:       true,              // Game is now active!
		TargetNumber: targetNumber,      // This is the secret number!
		FirstRoll:    0,                 // Haven't rolled yet
		HasFirstRoll: false,             // No first roll yet
	}
}

// Start activates a new game
// Call this when user wants to play!
func (g *Game) Start() {
	// Generate new random target number
	g.TargetNumber = rand.Intn(11) + 2
	g.Active = true
	g.FirstRoll = 0
	g.HasFirstRoll = false
}

// End closes the game
// Call this when user wants to stop playing
func (g *Game) End() {
	g.Active = false
	g.TargetNumber = 0
	g.FirstRoll = 0
	g.HasFirstRoll = false
}

// IsActive checks if game is currently running
// Like asking: "Am I still playing?"
func (g *Game) IsActive() bool {
	return g.Active
}

// GetTargetNumber returns the secret number
// (We usually don't show this to the player - it's secret!)
func (g *Game) GetTargetNumber() int {
	return g.TargetNumber
}

// SetFirstRoll saves the first die roll
// Call this after rolling the first die
func (g *Game) SetFirstRoll(value int) {
	g.FirstRoll = value
	g.HasFirstRoll = true
}

// GetFirstRoll returns what you rolled first
func (g *Game) GetFirstRoll() int {
	return g.FirstRoll
}

// ResetRollPair clears the roll pair for next attempt
// Call this after completing a pair (win or lose)
func (g *Game) ResetRollPair() {
	g.FirstRoll = 0
	g.HasFirstRoll = false
}

// NeedsFirstRoll checks if we need to roll first die
// Returns true if we haven't done first roll yet
func (g *Game) NeedsFirstRoll() bool {
	return !g.HasFirstRoll
}

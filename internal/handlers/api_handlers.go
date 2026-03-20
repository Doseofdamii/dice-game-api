package handlers

import (
	"encoding/json"
	"net/http"
	
	"dice-game-api/internal/services"
)

// APIHandler handles all HTTP requests
// This is like the receptionist - takes requests and gives responses!
type APIHandler struct {
	gameService *services.GameService
}

// NewAPIHandler creates a new API handler
func NewAPIHandler(gameService *services.GameService) *APIHandler {
	return &APIHandler{
		gameService: gameService,
	}
}

// =====================================
// WALLET ENDPOINTS
// =====================================

// FundWallet handles POST /wallet/fund
// Adds 155 sats to wallet (if balance <= 35)
func (h *APIHandler) FundWallet(w http.ResponseWriter, r *http.Request) {
	// Only accept POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Try to fund the wallet
	err := h.gameService.FundWallet()
	if err != nil {
		// Send error message
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	
	// Success! Get the new balance
	balance, balanceStr, asset := h.gameService.GetBalanceInfo()
	
	// Send success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":        "Wallet funded successfully",
		"balance":        balance,
		"balance_string": balanceStr,
		"asset":          asset,
	})
}

// GetBalance handles GET /wallet/balance
// Returns current wallet balance
func (h *APIHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	// Only accept GET method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Get balance info
	balance, balanceStr, asset := h.gameService.GetBalanceInfo()
	
	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"balance":        balance,
		"balance_string": balanceStr,
		"asset":          asset,
	})
}

// =====================================
// GAME ENDPOINTS
// =====================================

// StartGame handles POST /game/start
// Starts a new game (costs 20 sats)
func (h *APIHandler) StartGame(w http.ResponseWriter, r *http.Request) {
	// Only accept POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Try to start game
	game, err := h.gameService.StartGame()
	if err != nil {
		// Send error message
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	
	// Success! Get new balance
	balance := h.gameService.GetBalance()
	
	// Send success response
	// NOTE: We DON'T show the target number - it's a secret!
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Game started! Roll dice to play.",
		"active":  game.Active,
		"balance": balance,
		"cost":    20,
	})
}

// EndGame handles POST /game/end
// Ends the current game session
func (h *APIHandler) EndGame(w http.ResponseWriter, r *http.Request) {
	// Only accept POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Try to end game
	err := h.gameService.EndGame()
	if err != nil {
		// Send error message
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	
	// Success! Get final balance
	balance := h.gameService.GetBalance()
	
	// Send success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Game ended",
		"balance": balance,
	})
}

// GetGameStatus handles GET /game/status
// Returns whether a game is currently active
func (h *APIHandler) GetGameStatus(w http.ResponseWriter, r *http.Request) {
	// Only accept GET method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Get game status
	status := h.gameService.GetGameStatus()
	
	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

// =====================================
// DICE ENDPOINTS
// =====================================

// RollDice handles POST /dice/roll
// Rolls a single die (part of a pair)
// First roll: costs 5 sats
// Second roll: FREE!
func (h *APIHandler) RollDice(w http.ResponseWriter, r *http.Request) {
	// Only accept POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Try to roll dice
	result, err := h.gameService.RollDice()
	if err != nil {
		// Send error message
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	
	// Success! Send the roll result
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// =====================================
// HELPER ENDPOINTS
// =====================================

// Home handles GET /
// Shows API information and available endpoints
func (h *APIHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "🎲 Dice Game API",
		"version": "1.0.0",
		"endpoints": map[string]interface{}{
			"wallet": map[string]string{
				"POST /wallet/fund":    "Fund wallet with 155 sats (only if balance <= 35)",
				"GET  /wallet/balance": "Get current wallet balance",
			},
			"game": map[string]string{
				"POST /game/start":  "Start new game (costs 20 sats)",
				"POST /game/end":    "End current game",
				"GET  /game/status": "Check if game is active",
			},
			"dice": map[string]string{
				"POST /dice/roll": "Roll dice (first roll: 5 sats, second: free)",
			},
		},
		"rules": map[string]interface{}{
			"game_cost":     "20 sats",
			"roll_cost":     "5 sats per pair (charged on first roll)",
			"win_prize":     "10 sats",
			"target_range":  "2-12",
			"fund_amount":   "155 sats",
			"fund_limit":    "Can only fund if balance <= 35 sats",
		},
	})
}

// Health handles GET /health
// Simple health check endpoint
func (h *APIHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
	})
}

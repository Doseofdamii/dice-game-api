package handlers

import "net/http"

// SetupRoutes configures all the API endpoints
// This is like creating a map of where everything is!
func SetupRoutes(handler *APIHandler) *http.ServeMux {
	mux := http.NewServeMux()
	
	// Home & Health
	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/health", handler.Health)
	
	// Wallet endpoints
	mux.HandleFunc("/wallet/fund", handler.FundWallet)       // POST - Fund wallet
	mux.HandleFunc("/wallet/balance", handler.GetBalance)    // GET - Check balance
	
	// Game endpoints
	mux.HandleFunc("/game/start", handler.StartGame)         // POST - Start game
	mux.HandleFunc("/game/end", handler.EndGame)             // POST - End game
	mux.HandleFunc("/game/status", handler.GetGameStatus)    // GET - Check status
	
	// Dice endpoints
	mux.HandleFunc("/dice/roll", handler.RollDice)           // POST - Roll dice
	
	return mux
}

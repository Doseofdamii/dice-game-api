package main

import (
	"fmt"
	"log"
	"net/http"
	
	"dice-game-api/internal/handlers"
	"dice-game-api/internal/services"
)

func main() {
	// Print welcome banner
	fmt.Println("================================")
	fmt.Println("🎲 Dice Game API Starting...")
	fmt.Println("================================")
	fmt.Println()
	
	// STEP 1: Create the game service
	// This is the brain that handles all the game logic!
	gameService := services.NewGameService()
	fmt.Println("✅ Game Service initialized")
	
	// STEP 2: Create the API handler
	// This handles all the HTTP requests!
	apiHandler := handlers.NewAPIHandler(gameService)
	fmt.Println("✅ API Handler created")
	
	// STEP 3: Setup routes
	// This maps URLs to handlers!
	router := handlers.SetupRoutes(apiHandler)
	fmt.Println("✅ Routes configured")
	
	fmt.Println()
	fmt.Println("================================")
	fmt.Println("🌐 Server running on http://localhost:8080")
	fmt.Println("================================")
	fmt.Println()
	fmt.Println("📝 Available Endpoints:")
	fmt.Println()
	fmt.Println("WALLET:")
	fmt.Println("  POST   /wallet/fund      - Fund wallet with 155 sats")
	fmt.Println("  GET    /wallet/balance   - Check wallet balance")
	fmt.Println()
	fmt.Println("GAME:")
	fmt.Println("  POST   /game/start       - Start new game (costs 20 sats)")
	fmt.Println("  POST   /game/end         - End current game")
	fmt.Println("  GET    /game/status      - Check if game is active")
	fmt.Println()
	fmt.Println("DICE:")
	fmt.Println("  POST   /dice/roll        - Roll dice (5 sats per pair)")
	fmt.Println()
	fmt.Println("OTHER:")
	fmt.Println("  GET    /                 - API information")
	fmt.Println("  GET    /health           - Health check")
	fmt.Println()
	fmt.Println("================================")
	fmt.Println()
	fmt.Println("💡 How to Play:")
	fmt.Println()
	fmt.Println("1. Fund your wallet:")
	fmt.Println("   POST /wallet/fund")
	fmt.Println()
	fmt.Println("2. Start a game:")
	fmt.Println("   POST /game/start")
	fmt.Println()
	fmt.Println("3. Roll dice (twice per pair):")
	fmt.Println("   POST /dice/roll  (1st roll - costs 5 sats)")
	fmt.Println("   POST /dice/roll  (2nd roll - FREE!)")
	fmt.Println()
	fmt.Println("4. Keep rolling or end game:")
	fmt.Println("   POST /dice/roll  (new pair - costs 5 sats)")
	fmt.Println("   POST /game/end   (finish playing)")
	fmt.Println()
	fmt.Println("================================")
	fmt.Println()
	fmt.Println("🎯 Game Rules:")
	fmt.Println("  - Start game: 20 sats")
	fmt.Println("  - Roll pair: 5 sats (charged on 1st roll)")
	fmt.Println("  - Win prize: 10 sats")
	fmt.Println("  - Target: Random number 2-12")
	fmt.Println("  - Goal: Roll sum = target number")
	fmt.Println()
	fmt.Println("================================")
	fmt.Println()
	fmt.Println("Press Ctrl+C to stop the server")
	fmt.Println()
	
	// STEP 4: Start the HTTP server!
	port := "8080"
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}

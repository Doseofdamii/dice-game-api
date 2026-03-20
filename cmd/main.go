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
	gameService := services.NewGameService()
	fmt.Println("✅ Game Service initialized")
	
	// STEP 2: Create the API handler
	apiHandler := handlers.NewAPIHandler(gameService)
	fmt.Println("✅ API Handler created")
	
	// STEP 3: Setup routes
	router := handlers.SetupRoutes(apiHandler)
	fmt.Println("✅ Routes configured")
	
	// STEP 4: Serve static files (the website!)
	// This serves the HTML file from the public folder
	fs := http.FileServer(http.Dir("./public"))
	router.Handle("/play", http.StripPrefix("/play", fs))
	fmt.Println("✅ Website configured at /play")
	
	fmt.Println()
	fmt.Println("================================")
	fmt.Println("🌐 Server running on http://localhost:8080")
	fmt.Println("================================")
	fmt.Println()
	fmt.Println("🎮 PLAY THE GAME:")
	fmt.Println("   Open browser: http://localhost:8080/play")
	fmt.Println()
	fmt.Println("📝 API Endpoints:")
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
	fmt.Println("  GET    /play             - Play the game in browser!")
	fmt.Println()
	fmt.Println("================================")
	fmt.Println()
	fmt.Println("Press Ctrl+C to stop the server")
	fmt.Println()
	
	// STEP 5: Start the HTTP server!
	port := "8080"
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}

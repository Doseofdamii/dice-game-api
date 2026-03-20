# 🎲 Dice Game REST API

A simple dice game where you roll dice to match a target number and win sats!

---

## 🎯 How The Game Works

**Like a carnival dice game:**
1. **Fund Wallet** → Get 155 sats (coins)
2. **Start Game** → Pay 20 sats, computer picks secret number (2-12)
3. **Roll Dice (Pair)** → Pay 5 sats, roll twice
4. **Win** → If your sum matches secret number, win 10 sats!
5. **Repeat** → Keep rolling or end game

---

## 📋 Game Rules

| Action | Cost | Result |
|--------|------|--------|
| Fund Wallet | Free | Get 155 sats (only if balance ≤ 35) |
| Start Game | 20 sats | Random target number generated (2-12) |
| Roll Dice (1st) | 5 sats | Roll a die (1-6), saved |
| Roll Dice (2nd) | Free | Roll again, check if won |
| Win | - | Get 10 sats! 🎉 |
| Lose | - | Lost 5 sats, try again |

---

## 🚀 How to Run

### 1. Create Project Folder
```bash
mkdir dice-game-api
cd dice-game-api
```

### 2. Download Files
Download all files and organize like this:
```
dice-game-api/
├── go.mod
├── cmd/
│   └── main.go
└── internal/
    ├── domain/
    │   ├── wallet.go
    │   ├── game.go
    │   └── dice.go
    ├── services/
    │   └── game_service.go
    └── handlers/
        ├── api_handlers.go
        └── routes.go
```

### 3. Run the Server
```bash
go run cmd/main.go
```

You should see:
```
🎲 Dice Game API Starting...
✅ Game Service initialized
✅ API Handler created
✅ Routes configured
🌐 Server running on http://localhost:8080
```

---

## 📡 API Endpoints

### **Wallet Endpoints**

#### Fund Wallet
```bash
POST /wallet/fund
```
**Response:**
```json
{
  "message": "Wallet funded successfully",
  "balance": 155,
  "balance_string": "155",
  "asset": "sats"
}
```

#### Get Balance
```bash
GET /wallet/balance
```
**Response:**
```json
{
  "balance": 155,
  "balance_string": "155",
  "asset": "sats"
}
```

---

### **Game Endpoints**

#### Start Game
```bash
POST /game/start
```
**Response:**
```json
{
  "message": "Game started! Roll dice to play.",
  "active": true,
  "balance": 135,
  "cost": 20
}
```

#### End Game
```bash
POST /game/end
```
**Response:**
```json
{
  "message": "Game ended",
  "balance": 140
}
```

#### Get Game Status
```bash
GET /game/status
```
**Response:**
```json
{
  "active": true,
  "has_first_roll": false
}
```

---

### **Dice Endpoints**

#### Roll Dice
```bash
POST /dice/roll
```

**First Roll Response:**
```json
{
  "roll": 3,
  "is_first_roll": true,
  "is_second_roll": false,
  "balance": 130,
  "message": "First roll: 3. Roll again to complete the pair!"
}
```

**Second Roll Response (Win):**
```json
{
  "roll": 4,
  "is_first_roll": false,
  "is_second_roll": true,
  "first_roll": 3,
  "second_roll": 4,
  "sum": 7,
  "target": 7,
  "won": true,
  "prize": 10,
  "balance": 140,
  "message": "🎉 YOU WIN! 3 + 4 = 7 (target: 7). Won 10 sats!"
}
```

**Second Roll Response (Lose):**
```json
{
  "roll": 2,
  "is_first_roll": false,
  "is_second_roll": true,
  "first_roll": 3,
  "second_roll": 2,
  "sum": 5,
  "target": 7,
  "won": false,
  "balance": 130,
  "message": "You rolled 3 + 2 = 5 (target: 7). Try again!"
}
```

---

## 🧪 Testing with Thunder Client

### Test 1: Fund Wallet
1. **POST** `http://localhost:8080/wallet/fund`
2. Click **Send**

**Expected:** `"balance": 155`

---

### Test 2: Check Balance
1. **GET** `http://localhost:8080/wallet/balance`
2. Click **Send**

**Expected:** `"balance": 155`

---

### Test 3: Start Game
1. **POST** `http://localhost:8080/game/start`
2. Click **Send**

**Expected:** 
- `"active": true`
- `"balance": 135` (155 - 20)

---

### Test 4: Roll Dice (First Roll)
1. **POST** `http://localhost:8080/dice/roll`
2. Click **Send**

**Expected:**
- `"is_first_roll": true`
- `"roll": 1-6` (random)
- `"balance": 130` (135 - 5)

---

### Test 5: Roll Dice (Second Roll)
1. **POST** `http://localhost:8080/dice/roll` (again!)
2. Click **Send**

**Expected:**
- `"is_second_roll": true`
- `"sum": ...` (first + second)
- `"won": true or false`
- If won: `"balance": 140` (130 + 10)
- If lost: `"balance": 130`

---

### Test 6: Keep Playing!
1. Keep calling **POST** `/dice/roll`
2. Each pair costs 5 sats
3. Win = +10 sats
4. Try to get as many sats as possible!

---

### Test 7: End Game
1. **POST** `http://localhost:8080/game/end`
2. Click **Send**

**Expected:** `"message": "Game ended"`

---

## 🎮 Complete Game Example

```bash
# 1. Fund wallet
POST /wallet/fund
→ Balance: 155 sats

# 2. Start game
POST /game/start
→ Balance: 135 sats (cost: 20)
→ Secret target: 7 (you don't see this!)

# 3. Roll #1 (first of pair)
POST /dice/roll
→ Balance: 130 sats (cost: 5)
→ Rolled: 3

# 4. Roll #2 (second of pair)
POST /dice/roll
→ Balance: 140 sats (won 10!)
→ Rolled: 4
→ Sum: 3 + 4 = 7
→ YOU WIN! 🎉

# 5. Roll #3 (new pair)
POST /dice/roll
→ Balance: 135 sats (cost: 5)
→ Rolled: 2

# 6. Roll #4 (second of pair)
POST /dice/roll
→ Balance: 135 sats
→ Rolled: 3
→ Sum: 2 + 3 = 5
→ Lost (target was 7)

# 7. Roll #5 (new pair)
POST /dice/roll
→ Balance: 130 sats (cost: 5)
→ Rolled: 5

# 8. Roll #6 (second of pair)
POST /dice/roll
→ Balance: 140 sats (won 10!)
→ Rolled: 2
→ Sum: 5 + 2 = 7
→ YOU WIN AGAIN! 🎉

# 9. End game
POST /game/end
→ Final balance: 140 sats
→ Profit: -15 sats (155 - 140)
```

---

## 📊 Complexity

**Rating: 10/20** 🎯

**What makes this project valuable:**
- ✅ Clean architecture (domain, services, handlers)
- ✅ State management (wallet + game state)
- ✅ Business logic (game rules enforcement)
- ✅ Random number generation
- ✅ In-memory storage
- ✅ RESTful API design
- ✅ Error handling
- ✅ Clear separation of concerns

**Perfect for:**
- Portfolio projects
- Learning game logic
- Understanding state management
- REST API practice

---

## 🎓 What You Learned

1. **State Management** - Tracking game state across requests
2. **Business Logic** - Implementing game rules
3. **Random Generation** - Dice rolls and target numbers
4. **API Design** - Clean, intuitive endpoints
5. **Error Handling** - Proper validation and errors

---

## 🏆 Challenge Yourself

Try adding these features:
- Transaction log (track all wallet changes)
- Multiple games (save game history)
- High score system
- Different difficulty levels (different target ranges)
- Betting multipliers

---

**Have fun playing! 🎲**

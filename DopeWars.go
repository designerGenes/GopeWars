package main

import (
	"fmt"
)

/**
Let's detail the basic rules of the game:
- The player starts with $2000 cash and $5500 debt.
- The player always inhabits one location at a time.  The player can move to a different location once per turn
- The game has 30 days.  Each day counts as 1 turn.  In each turn, the player can:
- - Buy and sell drugs as many times as they like
- - Pay back any amount of their debt or take out a loan if their debt is less than $1000
- - Move to a different location once.  This ends the turn
- - Quit the game
- Each drug has a random price and quantity available to purchase that is generated each day
- Each drug has an inherent value multiplier that is used to calculate the drug's price.  This multiplier is multiplied by the random price.
- For example, cocaine has a multiplier of 3.  So if the random price returns 100, the price of each unit of cocaine will be $300
- The player can sell and buy each drug at the same rate.  So if in one city the player can buy cocaine for $300, they can also sell it for $300
- The player can only buy as many drugs as they can afford
- The player can also spend money on paying back any amount of their debt.  If their debt is lower than $1000, they can take out a loan for up $10,000
- If the game ends (all the days/turns have run out), and the player has any debt remaining, they lose the game
- If the game ends and the player has no debt, the player's score is equal to the cash they have remaining

Later rules (not implemented yet):
- The player can purchase a gun for $500.  This allows the player to defend themselves against muggers and police
- Randomly, a player may be mugged.  If the player has a gun, they can defend themselves.
- - If they don't have a gun in this case, they lose half their cash.
- - If they do have at least one gun in this case, they sacrifice one gun
- A player can purchase any number of guns as long as they can afford them
*/

type Drug struct {
	Name       string
	Multiplier float64
}

type Player struct {
	Cash     int
	Drugs    DrugInventory
	Debt     int
	Location string
}

type Game struct {
	Player     Player
	Locations  []string
	DrugTypes  []Drug // slice of Drug structs
	CurrentDay int
	TotalDays  int
}

type DrugInventory struct {
	weed    int
	cocaine int
	heroin  int
	meth    int
	lsd     int
	ecstasy int
}

type DrugListing struct {
	Drug     Drug
	Price    int
	Quantity int
}

func NewGame() *Game {
	return &Game{
		Player: Player{
			Cash:  500,
			Drugs: DrugInventory{},
			Debt:  5500,
		},
		Locations: []string{"Bronx", "Brooklyn", "Manhattan", "Queens", "Staten Island"},
		DrugTypes: []Drug{
			Drug{"Weed", 1.0},
			Drug{"Cocaine", 3.0},
			Drug{"Heroin", 2.5},
			Drug{"Meth", 1.5},
			Drug{"LSD", 2.0},
			Drug{"Ecstasy", 2.0},
		},
		CurrentDay: 1,
		TotalDays:  30,
	}
}

func (g *Game) Start() {
	fmt.Println("Welcome to Dope Wars!")
	fmt.Printf("You have %d days to make as much money as possible.\n", g.TotalDays)

	for g.CurrentDay <= g.TotalDays {
		g.PlayTurn()
		g.CurrentDay++
	}

	g.EndGame()
}

func (g *Game) PlayTurn() {
	// init slice to hold drug prices which are dynamically generated for this turn
	drugMenu := []DrugListing{}
	for _, drug := range g.DrugTypes {
		// generate random price and quantity for each drug
		price := 100 * drug.Multiplier
		quantity := 10
		drugMenu = append(drugMenu, DrugListing{Drug: drug, Price: int(price), Quantity: quantity})
	}

	// update UI to list current location, cash, debt, drugs, etc.

	// clear terminal
	fmt.Print("\033[H\033[2J")

	fmt.Printf("\nDay: %d\n\t Days Remaining: %d", g.CurrentDay, g.TotalDays-g.CurrentDay)
	// print Player financial status
	fmt.Printf("\nCash: %d | Debt : %d", g.Player.Cash, g.Player.Debt)
	// print Player drug quantities
	fmt.Printf("\nWeed: %d | Cocaine: %d | Heroin: %d | Meth: %d | LSD: %d | Ecstasy: %d", g.Player.Drugs.weed, g.Player.Drugs.cocaine, g.Player.Drugs.heroin, g.Player.Drugs.meth, g.Player.Drugs.lsd, g.Player.Drugs.ecstasy)

	// print Player Location
	fmt.Printf("\n\nLocation: %s", g.Player.Location)
	// display drugs available for purchase
	fmt.Println("\n\nDrugs Available for Purchase:")
	for _, drug := range drugMenu {
		// generate random drug price and quantity, and multiply price by drug multiplier
		fmt.Printf("\n%s: $%d\tQ: %d|", drug.Drug.Name, drug.Price, drug.Quantity)

	}
	fmt.Printf("\n\nWhat would you like to do?\n")
	// accept user input for action

	_, playerInput := fmt.Scan("1. Buy Drugs\n2. Sell Drugs\n3. Pay Debt\n4. Take out Loan\n5. Move to a different location\n6. Quit the game\n")

	switch playerInput {
	case 1:
		fmt.Println("You chose to buy drugs")
	case 2:
		fmt.Println("You chose to sell drugs")
	case 3:
		fmt.Println("You chose to pay debt")
	case 4:
		fmt.Println("You chose to take out a loan")
	case 5:
		fmt.Println("You chose to move to a different location")
	case 6:
		fmt.Println("You chose to quit the game")
	

}

func (g *Game) EndGame() {
	fmt.Println("Thanos has arrived.  The game is over.")
}

func main() {
	game := NewGame()
	game.Start()
}

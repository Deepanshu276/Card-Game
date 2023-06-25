package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Card represents a playing card with a rank and suit.
type Card struct {
Rank string
Suit string
}

// Player represents a player in the game with a name and a hand of cards.
type Player struct {
Name  string
Hand  []Card
Skips int
}

// Game represents the card game with players, deck, and discard pile.
type Game struct {
Players     []*Player
Deck        []Card
DiscardPile []Card
Turn        int
Direction   int   //// 1 for clockwise, -1 for counterclockwise
PlayedCard  Card
}

// Initialize the game by creating players, shuffling the deck, and dealing cards.
func (g *Game) Initialize(numPlayers int) {
	reader := bufio.NewReader(os.Stdin)

if numPlayers < 2 || numPlayers > 4 {
	fmt.Println("Invalid number of players. Please select between 2 and 4 players.")
	fmt.Print("Enter the number of players: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, _ := strconv.Atoi(input)
	g.Initialize(n)
	return
}

for i := 0; i < numPlayers; i++ {
	fmt.Printf("Enter the name of Player %d: ", i+1)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	g.Players = append(g.Players, &Player{Name: name})
}

	// Create and shuffle the deck
	g.Deck = createDeck()
	g.shuffleDeck()

	// Deal 5 cards to each player

	for i := 0; i < 5; i++ {
		for _, player := range g.Players {
			card := g.drawCard()
			player.Hand = append(player.Hand, card)
	}
	}

	// Place the first card on the discard pile
	card := g.drawCard()
	g.DiscardPile = append(g.DiscardPile, card)

	g.Turn = 0

	g.Direction = 1
}

// Play a card from the current player's hand.
func (g *Game) PlayCard(playerIndex int, cardIndex int) (bool, error) {
	player := g.Players[playerIndex]
	if playerIndex != g.Turn {
	return false, fmt.Errorf("it's not %s's turn", player.Name)
	}
	
	if cardIndex < 0 || cardIndex >= len(player.Hand) {
	return false, fmt.Errorf("invalid card index")
	}
	
	card := player.Hand[cardIndex]
	
	topCard := g.DiscardPile[len(g.DiscardPile)-1]
	if card.Rank >= topCard.Rank && isSuitPriorityGreater(card.Suit, topCard.Suit) {
		// Special action cards
		switch card.Rank {
		case "A":
		nextPlayer := (g.Turn + g.Direction*2) % len(g.Players)
		if nextPlayer < 0 {
		nextPlayer += len(g.Players)
		}
		g.Players[nextPlayer].Skips++
		g.Turn = nextPlayer
	case "K":
	g.reverseDirection()
	case "Q":
	g.nextPlayerDrawCards(2)
	g.Players[g.getNextPlayerIndex()].Skips++
	case "J":
	g.nextPlayerDrawCards(4)
	g.Players[g.getNextPlayerIndex()].Skips++
	}
	
	g.PlayedCard = card 
	
	player.Hand = append(player.Hand[:cardIndex], player.Hand[cardIndex+1:]...)
	
	g.Turn = g.getNextPlayerIndex()
	
	return true, nil
	}
	
	return false, nil
	}
	

// Get the index of the next player in turn.
func (g *Game) getNextPlayerIndex() int {
	nextPlayer := (g.Turn + g.Direction) % len(g.Players)
	if nextPlayer < 0 {
		nextPlayer += len(g.Players)
		}
	return nextPlayer
		}

func isSuitPriorityGreater(suit1 string, suit2 string) bool {
	suitPriority := map[string]int{
	"Spades":   4,
	"Clubs":    1,
	"Diamonds": 2,
	"Hearts":   3,
}

return suitPriority[suit1] >= suitPriority[suit2]
}

func (g *Game) updateCurrentDiscard() {
	if len(g.Deck) > 0 {
	index := rand.Intn(len(g.Deck))
	card := g.Deck[index]
	g.Deck = append(g.Deck[:index], g.Deck[index+1:]...)
	g.DiscardPile = append(g.DiscardPile, card)
	}
	}

// Draw a card from the draw pile.
func (g *Game) DrawCard(playerIndex int) (bool, error) {
	player := g.Players[playerIndex]
	if playerIndex != g.Turn {
	return false, fmt.Errorf("it's not %s's turn", player.Name)
		}
		
	if player.Skips > 0 {
	player.Skips--
	return false, fmt.Errorf("%s's turn is skipped", player.Name)
		}
		
	card := g.drawCard()
	player.Hand = append(player.Hand, card)
		
	g.Turn = g.getNextPlayerIndex()
		
	g.updateCurrentDiscard()
		
	return true, nil
		}

// Reverse the direction of play.			

func (g *Game) reverseDirection() {
	g.Direction *= -1
}

	
// Draw a card from the draw pile.
func (g *Game) drawCard() Card {
	if len(g.Deck) == 0 {
	return Card{}
		}
	card := g.Deck[len(g.Deck)-1]
	g.Deck = g.Deck[:len(g.Deck)-1]
	return card
		}

// Make the next player draw a specified number of cards.
func (g *Game) nextPlayerDrawCards(numCards int) {
	nextPlayer := g.getNextPlayerIndex()
	for i := 0; i < numCards; i++ {
	card := g.drawCard()
	g.Players[nextPlayer].Hand = append(g.Players[nextPlayer].Hand, card)
	}
	}
	
func createDeck() []Card {
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}

	var deck []Card
	for _, suit := range suits {
	for _, rank := range ranks {
	card := Card{Rank: rank, Suit: suit}
	deck = append(deck, card)
}
}

return deck
}

func (g *Game) shuffleDeck() {
	rand.Seed(time.Now().UnixNano())
	n := len(g.Deck)
	for i := n - 1; i > 0; i-- {
	j := rand.Intn(i + 1)
	g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i]
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of players (2-4): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	numPlayers, _ := strconv.Atoi(input)

	game := Game{}
	game.Initialize(numPlayers)

for {
	currentPlayer := game.Players[game.Turn]
	fmt.Printf("Current Player: %s\n", currentPlayer.Name)
	
	fmt.Println("Current Discard Card:", game.DiscardPile[len(game.DiscardPile)-1])
	
	fmt.Printf("Player %s's Hand: ", currentPlayer.Name)
	for i, card := range currentPlayer.Hand {
		fmt.Printf("[%d: %s %s] ", i, card.Rank, card.Suit)
	}
	fmt.Println()
	
	fmt.Print("Enter the index of the card to play (or 'd' to draw): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	if input == "d" {
		success, err := game.DrawCard(game.Turn)
	if err != nil {
		fmt.Println("Error:", err)
	} else if success {
		fmt.Printf("Player %s drew a card\n", currentPlayer.Name)
	} else {
		fmt.Println("It's not your turn to draw a card")
	}
	} else {
		cardIndex, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input")
		continue
	}
	
	success, err := game.PlayCard(game.Turn, cardIndex)
	if err != nil {
	fmt.Println("Error:", err)
	} else if success {
	fmt.Printf("Player %s played a card\n", currentPlayer.Name)
	} else {
	fmt.Println("Invalid card play")
	}
	}
	
	fmt.Println("----------")
	
	if len(currentPlayer.Hand) == 0 {
	fmt.Printf("Player %s has no more cards!\n", currentPlayer.Name)
	break
	}
	}
	
	fmt.Println("Game Over")
	}
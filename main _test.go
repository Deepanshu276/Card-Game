package main

import (
	"testing"
)

func TestInitialize(t *testing.T) {
game := Game{}
game.Initialize(2)

if len(game.Players) != 2 {
t.Errorf("Expected 2 players, got %d", len(game.Players))
}

if len(game.Deck) != 42 { // 52 cards in a deck - 5 cards for each player - 1 card for discard pile
t.Errorf("Expected 42 cards in deck, got %d", len(game.Deck))
}

if len(game.DiscardPile) != 1 {
t.Errorf("Expected 1 card in discard pile, got %d", len(game.DiscardPile))
}
}

func TestPlayCard(t *testing.T) {
game := Game{}
game.Initialize(2)

// Assume the first player has a card with rank "2" and suit "Spades"
game.Players[0].Hand[0] = Card{Rank: "2", Suit: "Spades"}

// Assume the top card in the discard pile has rank "A" and suit "Clubs"
game.DiscardPile[0] = Card{Rank: "A", Suit: "Clubs"}

success, err := game.PlayCard(0, 0)
if err != nil {
t.Errorf("Expected no error, got %v", err)
}

if !success {
t.Errorf("Expected success, got %v", success)
}

if len(game.Players[0].Hand) != 1 { // Player should have 4 cards left after playing one
t.Errorf("Expected 5 cards in player's hand, got %d", len(game.Players[0].Hand))
}
}

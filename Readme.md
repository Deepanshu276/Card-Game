# Multiplayer Card Game

This is a multiplayer card game implemented in Go. The game supports 2 to 4 players and uses a standard deck of playing cards. The objective of the game is to be the first player to get rid of all the cards in their hand.

## Rules

- Each player starts with a hand of 5 cards.
- The game starts with a deck of 52 cards.
- Players take turns playing cards from their hand, following a set of rules that define what cards can be - played when.
- A player can only play a card if it matches either the suit or the rank of the top card on the discard    pile.
- If a player cannot play a card, they must draw a card from the draw pile. If the draw pile is empty, the game ends in a draw and no player is declared a winner.
- The game ends when one player runs out of cards.

## Actions Cards

- Ace(A): Skip the next player in turn

- Kings(K): Reverse the sequence of who plays next 

- Queens(Q): +2

- Jacks(J): +4

- NOTE: actions are not stackable i.e. if Q is played by player 1 then player two draws two cards and cannot play a Q from his hand on that turn even if available

# USAGE

1. Clone the Repositery

```
https://github.com/Deepanshu276/Card-Game

```
2. Install the dependency required to play the game 

```
go mod tidy 

```

3. Run the game using :

```
go run main.go

```

## Instruction to play the game 

1. Follow the on-screen instructions to play the game.
2. Enter the number of players (2-4) and provide the names of each player.
3. During each player's turn, enter the index of the card to play or enter 'd' to draw a card.
4. The game will display the current player, the discard pile's top card, and the player's hand.
5. Continue playing until a player runs out of cards or the draw pile is empty.



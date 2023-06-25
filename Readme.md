# Multiplayer Card Game

This is a multiplayer card game implemented in Go. The game supports 2 to 4 players and uses a standard deck of playing cards. The objective of the game is to be the first player to get rid of all the cards in their hand.

## Rules

- Each player starts with a hand of 5 cards.
- The game starts with a deck of 52 cards.
- Players take turns playing cards from their hand, following a set of rules that define what cards can be - played when.
- A player can only play a card if it matches either the suit or the rank of the top card on the discard    pile.
- If a player cannot play a card, they must draw a card from the draw pile. If the draw pile is empty, the game ends in a draw and no player is declared a winner.
- The game ends when one player runs out of cards.
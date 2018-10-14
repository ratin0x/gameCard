package gameCard

import (
	"math/rand"
	"time"
)

// CardType What type of card
type CardType int

const (
	// Nerf type
	Nerf CardType = 0
	// Buff type
	Buff CardType = 1
	// Score type
	Score CardType = 2
)

// Card type
type Card struct {
	TypeOfCard CardType
	Name       string
	Value      int
}

// Deck of cards
type Deck struct {
	Cards map[int]Card
	Name  string
	Size  int
}

// DeckConfig for Decks
type DeckConfig struct {
	TotalCards    int
	NumScoreCards int
	NumBuffCards  int
	NumNerfCards  int
}

// BasicScoreCard - The most basic type of score card
var BasicScoreCard = Card{Score, "Score", 1}

// BasicHeatCard - A basic type of heat card
var BasicHeatCard = Card{Nerf, "Heat", -5}

// BasicBuffCard - A basic type of buff card
var BasicBuffCard = Card{Buff, "Buff", 0}

var availableCards = []Card{
	BasicScoreCard,
	BasicHeatCard,
	BasicBuffCard,
}

// MakeBlankDeck - make a blank deck
func MakeBlankDeck(count int, name string) Deck {
	var cardsInDeck = make(map[int]Card, count)
	var deck Deck
	deck.Cards = cardsInDeck
	deck.Size = len(cardsInDeck)
	deck.Name = name
	return deck
}

// MakeRandomizedSimpleDeck - makes a deck from random selections of available card type
func MakeRandomizedSimpleDeck(count int, name string) Deck {
	// Make sure we're as close to random as possible by setting the seed
	var seed = time.Now().UnixNano()
	rand.Seed(seed)

	var cardsInDeck = make(map[int]Card, count)
	for i := 0; i < count; i++ {
		var randomCardSelector = rand.Intn(3)
		cardsInDeck[i] = availableCards[randomCardSelector]
		// fmt.Println("Selected ", randomCardSelector, "card ", availableCards[randomCardSelector])
	}
	var deck Deck
	deck.Cards = cardsInDeck
	deck.Size = len(cardsInDeck)
	deck.Name = name

	return deck
}

// MakeConfiguredDeck makes a deck conforming to the supplied DeckConfig
func MakeConfiguredDeck(config DeckConfig, name string) Deck {
	// Make sure we're as close to random as possible by setting the seed
	var seed = time.Now().UnixNano()
	rand.Seed(seed)

	var newDeck = MakeBlankDeck(config.TotalCards, name)
	var scoreCardsAdded, buffCardsAdded, nerfCardsAdded = 0, 0, 0

	for i := 0; i < config.TotalCards; i++ {
		var randomCardSelector = rand.Intn(3)
		switch randomCardSelector {
		case 0:
			if (scoreCardsAdded < config.NumScoreCards) || config.NumScoreCards == 0 {
				newDeck.Cards[i] = availableCards[randomCardSelector]
				scoreCardsAdded++
			}
		case 1:
			if (nerfCardsAdded < config.NumNerfCards) || config.NumNerfCards == 0 {
				newDeck.Cards[i] = availableCards[randomCardSelector]
				nerfCardsAdded++
			} else {
				newDeck.Cards[i] = availableCards[0]
			}
		case 2:
			if (buffCardsAdded < config.NumBuffCards) || config.NumBuffCards == 0 {
				newDeck.Cards[i] = availableCards[randomCardSelector]
				buffCardsAdded++
			} else {
				newDeck.Cards[i] = availableCards[0]
			}
		default:
			newDeck.Cards[i] = availableCards[randomCardSelector]

		}
	}
	return newDeck
}

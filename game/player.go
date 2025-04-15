package game

import (
	"errors"
	"fmt"
)

type Player struct {
	Name          string
	Hand 	      []*Card
	HasCalledUno  bool
	hasPlayedCard bool
}

func (p *Player) String() string {
	return fmt.Sprintf("Player %s (%d cards) | Uno(%t)", p.Name, len(p.Hand), p.HasCalledUno)
}

func NewPlayer(name string) *Player {
	player := &Player{
		Name: name,
		Hand: make([]*Card, 0),
		HasCalledUno: false,
		hasPlayedCard: false,
	}
	return player
}

func (p *Player) HandSize() int {
	return len(p.Hand)
}

func (p *Player) AddCard(card *Card) {
	if card != nil {
		p.Hand = append(p.Hand, card)

		if len(p.Hand) > 1 {
			p.HasCalledUno = false
		}
	}
}

func (p *Player) PlayCard(index int) (*Card, error) {
	if index < 0 || index >= len(p.Hand) {
		return nil, errors.New("Invalid card index")
	}

	card := p.Hand[index]

	p.Hand[index] = p.Hand[len(p.Hand)-1]
	p.Hand = p.Hand[:len(p.Hand)-1]
	
	if !p.hasPlayedCard {
		p.hasPlayedCard = true
	}

	return card, nil
}

func (p *Player) CallUno() {
	p.HasCalledUno = true
}

func (p *Player) HasWon() bool {
	return p.hasPlayedCard && len(p.Hand) == 0
}

func (p *Player) ShouldCallUno() bool {
	return p.hasPlayedCard && len(p.Hand) == 1
}

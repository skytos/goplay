package goplay

import "fmt"

// Player's can Play
type Player struct {
	Name string
	Score int
	Won bool
}

func (p *Player) Play() {
	if p.Score > 10 {
		p.Won = true
	}
	if p.Won {
		fmt.Printf("Player %s won with a score of %d\n", p.Name, p.Score)
	} else {
		fmt.Printf("Player %s is playing with a score of %d\n", p.Name, p.Score)
	}
}

// Play it's like work but more fun
func Play() {
	fmt.Println("Play!")
}

package goplay

import "fmt"

// Player's can Play
type Player struct {
	Name string
	Score int
}

func (p *Player) Play() {
	fmt.Printf("Player %s is playing with a score of %d\n", p.Name, p.Score)
}

// Play it's like work but more fun
func Play() {
	fmt.Println("Play!")
}

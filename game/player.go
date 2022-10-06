package game

import "strconv"

type Player struct {
	Id int
}

func SprintPlayer(p *Player) string {
	if p == nil {
		return "nil"
	}
	return strconv.Itoa(p.Id)
}

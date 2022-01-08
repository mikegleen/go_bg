package main

type Player struct {
	Id int
}

func SprintPlayer(p *Player) string {
	if p == nil {
		return "nil"
	}
	return string(p.Id)
}

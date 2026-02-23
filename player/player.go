package player

type AllPlayers struct {
	Players []*Player
}

type Player struct {
	Name            string
	CurrentPosition int
}

func New(names ...string) *AllPlayers {
	ap := &AllPlayers{}
	for _, n := range names {
		ap.Players = append(ap.Players, &Player{Name: n})
	}
	return ap
}

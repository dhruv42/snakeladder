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

// func (a *AllPlayers) Add(name string) *Player {
// 	// ap.Players = append(ap.Players, Player{Name: name})
// 	return &Player{
// 		Name:            name,
// 		CurrentPosition: 0,
// 	}
// }

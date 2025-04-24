package session

type GameSession struct {
	Score        int
	CurrentLevel int
}

func NewGameSession() *GameSession {
	return &GameSession{
		Score:        0,
		CurrentLevel: 0,
	}
}

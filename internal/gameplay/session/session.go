package session

type GameSession struct {
	Score int
}

func NewGameSession() *GameSession {
	return &GameSession{
		Score: 0,
	}
}

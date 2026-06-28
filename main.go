package main

type Teams struct {
	players []Player
}

type Player struct {
	Name                   string
	Goals, Misses, Assists int
	Rating                 float64
}

type TeamsInterface interface {
	NewPlayer(name string, goals, misses, assists int) Player
	goalsSort(players []Player) []Player
	ratingSort(players []Player) []Player
	gmSort(players []Player) []Player
}

func (player *Player) calculateRating() {
	var rating float64 = (float64(player.Goals) + float64(player.Assists)) / 2
	if player.Misses != 0 {
		rating /= float64(player.Misses)
	}
	player.Rating = rating
}

func NewPlayer(name string, goals, misses, assists int) Player {
	newPlayer := Player{Name: name, Goals: goals, Misses: misses, Assists: assists}
	newPlayer.calculateRating()
	return newPlayer
}

func main() {

}

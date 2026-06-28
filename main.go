package main

import (
	"slices"
)

type Player struct {
	Name                   string
	Goals, Misses, Assists int
	Rating                 float64
}

type PlayerInterface interface {
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

func goalsSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		switch {
		case a.Goals > b.Goals:
			return 1
		case a.Goals < b.Goals:
			return -1
		case a.Name > b.Name:
			return -1
		case a.Name < b.Name:
			return 1
		default:
			return 0
		}
	})
	return players
}

func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		switch {
		case a.Rating > b.Rating:
			return 1
		case a.Rating < b.Rating:
			return -1
		case a.Name < b.Name:
			return 1
		case a.Name > b.Name:
			return -1
		default:
			return 0
		}
	})
	return players
}

func gmSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		var gmA, gmB float64 = float64(a.Goals), float64(b.Goals)
		if a.Misses != 0 {
			gmA /= float64(a.Misses)
		}
		if b.Misses != 0 {
			gmB /= float64(b.Misses)
		}
		switch {
		case gmA < gmB:
			return 1
		case gmA > gmB:
			return -1
		case a.Name < b.Name:
			return 1
		case a.Name > b.Name:
			return -1
		default:
			return 0
		}
	})
	return players
}

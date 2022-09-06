package analysis

import (
	"fmt"

	"github.com/dpnetca/cflAnalysis/pkg/cfl"
)

func TotalPointsByGame(season cfl.Season) {
	for _, game := range season.Games {
		if !(game.Type.ID == 1 && game.Status.ID == 4) {
			continue
		}

		week := game.Week
		team1 := game.Team1.Nickname
		team2 := game.Team2.Nickname
		score := game.Team1.Score + game.Team2.Score
		fmt.Printf("Week %d: %s @ %s: %d\n", week, team1, team2, score)
	}
}

func AveragePointsPerGame(season cfl.Season) {
	gameCount := 0
	totalScore := 0
	for _, game := range season.Games {
		if !(game.Type.ID == 1 && game.Status.ID == 4) {
			continue
		}
		totalScore += game.Team1.Score
		totalScore += game.Team2.Score
		gameCount++
	}
	fmt.Printf("Average: %.1f\n", float32(totalScore)/float32(gameCount))
}

func TotalPointsByWeek(season cfl.Season) {
	previousWeek := 1
	week := 0
	score := 0
	for _, game := range season.Games {
		if !(game.Type.ID == 1 && game.Status.ID == 4) {
			continue
		}

		week = game.Week
		if week == previousWeek {
			score += game.Team1.Score
			score += game.Team2.Score
		} else {
			fmt.Printf("Week %d: %d\n", previousWeek, score)
			previousWeek = week
			score = game.Team1.Score + game.Team2.Score

		}

	}
	fmt.Printf("Week %d: %d\n", week, score)
}

func TotalOverUnderX(season cfl.Season, score float32) {
	over := 0
	under := 0
	match := 0
	for _, game := range season.Games {
		if !(game.Type.ID == 1 && game.Status.ID == 4) {
			continue
		}
		gameScore := game.Team1.Score + game.Team2.Score
		if float32(gameScore) > score {
			over++
		} else if float32(gameScore) < score {
			under++
		} else {
			match++
		}
	}
	fmt.Printf("Target: %.1f\n", score)
	fmt.Printf("Over: %d\n", over)
	fmt.Printf("Under: %d\n", under)
	fmt.Printf("Match: %d\n", match)
}

func WeeklyOverUnderX(season cfl.Season, score float32) {
	over := 0
	under := 0
	previousWeek := 1
	week := 0
	fmt.Printf("Target: %.1f\n", score)
	for _, game := range season.Games {
		if !(game.Type.ID == 1 && game.Status.ID == 4) {
			continue
		}

		week = game.Week
		if week != previousWeek {
			fmt.Printf("Week %d: Over: %d, Under: %d\n", previousWeek, over, under)
			previousWeek = week
			over = 0
			under = 0
		}

		gameScore := game.Team1.Score + game.Team2.Score
		if float32(gameScore) > score {
			over++
		} else {
			under++
		}
	}
}

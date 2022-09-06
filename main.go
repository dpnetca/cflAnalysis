package main

import (
	"fmt"
	"log"

	"github.com/dpnetca/cflAnalysis/pkg/analysis"
)

func main() {
	// godotenv.Load()
	// c := cfl.NewClient(os.Getenv("CFLAPIKEY"))

	// standings, err := c.GetStandings("2022")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// for i := 2000; i <= 2022; i++ {

	// 	games, err := c.GetAllGames(strconv.Itoa(i))
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	// analysis.TotalPointsByGame(games)
	// 	// fmt.Println("")
	// 	// analysis.TotalPointsByWeek(games)
	// 	// fmt.Println("")
	// 	fmt.Printf("%d ", i)
	// 	analysis.AveragePointsPerGame(games)
	// }

	ss, err := analysis.NewScheduleStrength("2022")
	if err != nil {
		log.Fatalln(err)
	}
	for _, team := range ss.Team {
		fmt.Printf("%v\n", team)
		playedWinningPct := float32(team.Played.OpposingWins) / float32(team.Played.GamesPlayed) * 100
		remainingWinningPct := float32(team.Remaining.OpposingWins) / float32(team.Remaining.GamesPlayed) * 100
		fmt.Printf("%s\n", team.TeamName)
		fmt.Printf("Played Winning Pct: %.2f\n", playedWinningPct)
		fmt.Printf("Remaining Winning Pct: %.2f\n", remainingWinningPct)
		fmt.Println()

	}
}

package analysis

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dpnetca/cflAnalysis/pkg/cfl"
	"github.com/joho/godotenv"
)

type total struct {
	GamesPlayed  int
	OpposingWins int
}
type lastYear struct {
	GamesPlayed  int
	OpposingWins int
}
type played struct {
	GamesPlayed  int
	OpposingWins int
}
type remaining struct {
	GamesPlayed  int
	OpposingWins int
}
type teamScheduleStrength struct {
	TeamId    int
	TeamName  string
	Played    played
	Remaining remaining
	Total     total
	LastYear  lastYear
}

type ScheduleStrength struct {
	Team []teamScheduleStrength
}

func NewScheduleStrength(year string) (ScheduleStrength, error) {
	var scheduleStrength ScheduleStrength
	godotenv.Load()
	c := cfl.NewClient(os.Getenv("CFLAPIKEY"))

	standings, err := c.GetStandings(year)
	if err != nil {
		log.Fatalln(err)
	}
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return scheduleStrength, err
	}

	previousYear := strconv.Itoa(yearInt - 1)
	previousStandings, err := c.GetStandings(previousYear)
	if err != nil {
		log.Fatalln(err)
	}
	schedule, err := c.GetAllGames(year)
	if err != nil {
		log.Fatalln(err)
	}

	teams, err := c.GetTeams()
	if err != nil {
		log.Fatalln(err)
	}

	// change this to initialize with loop through getTeams API
	// for i := 0; i <= 9; i++ {
	for _, team := range teams.Team {
		teamSchedule := teamScheduleStrength{TeamId: team.ID, TeamName: team.Name}
		scheduleStrength.Team = append(scheduleStrength.Team, teamSchedule)
	}

	err = scheduleStrength.total(schedule, standings)
	if err != nil {
		return scheduleStrength, err
	}
	err = scheduleStrength.lastYear(schedule, previousStandings)
	if err != nil {
		return scheduleStrength, err
	}

	err = scheduleStrength.remaining(schedule, standings)
	if err != nil {
		return scheduleStrength, err
	}
	err = scheduleStrength.played(schedule, standings)
	if err != nil {
		return scheduleStrength, err
	}

	return scheduleStrength, nil
}

func (s ScheduleStrength) total(schedule cfl.Season, standings cfl.Standings) error {
	opposingWins := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	gameCount := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, game := range schedule.Games {
		err := addTeamWins(game, standings, opposingWins, gameCount)
		if err != nil {
			return err
		}
	}
	for i, score := range opposingWins {
		if i == 0 {
			continue
		}
		s.Team[i-1].Total.OpposingWins = score
		s.Team[i-1].Total.GamesPlayed = gameCount[i]
	}
	return nil
}
func (s ScheduleStrength) lastYear(schedule cfl.Season, standings cfl.Standings) error {
	opposingWins := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	gameCount := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, game := range schedule.Games {
		err := addTeamWins(game, standings, opposingWins, gameCount)
		if err != nil {
			return err
		}
	}
	for i, score := range opposingWins {
		if i == 0 {
			continue
		}
		s.Team[i-1].LastYear.OpposingWins = score
		s.Team[i-1].LastYear.GamesPlayed = gameCount[i]
	}
	return nil
}

func (s ScheduleStrength) remaining(schedule cfl.Season, standings cfl.Standings) error {
	opposingWins := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	gameCount := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, game := range schedule.Games {
		if game.Status.ID == 4 {
			continue
		}
		err := addTeamWins(game, standings, opposingWins, gameCount)
		if err != nil {
			return err
		}
		// gameCount[game.Team1.ID]++
		// gameCount[game.Team2.ID]++
	}
	for i, score := range opposingWins {
		if i == 0 {
			continue
		}
		s.Team[i-1].Remaining.OpposingWins = score
		s.Team[i-1].Remaining.GamesPlayed = gameCount[i]
	}
	return nil
}

func (s ScheduleStrength) played(schedule cfl.Season, standings cfl.Standings) error {
	opposingWins := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	gameCount := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, game := range schedule.Games {
		if game.Status.ID != 4 {
			continue
		}
		err := addTeamWins(game, standings, opposingWins, gameCount)
		if err != nil {
			return err
		}
	}
	for i, score := range opposingWins {
		if i == 0 {
			continue
		}
		s.Team[i-1].Played.OpposingWins = score
		s.Team[i-1].Played.GamesPlayed = gameCount[i]
	}
	return nil
}

func addTeamWins(game cfl.Game, standings cfl.Standings, teamWins, teamGamesPlayed []int) error {
	played, wins, err := getPlayedAndWins(game.Team1.ID, &standings)
	if err != nil {
		return err
	}
	teamWins[game.Team2.ID] += wins
	teamGamesPlayed[game.Team2.ID] += played
	played, wins, err = getPlayedAndWins(game.Team2.ID, &standings)
	if err != nil {
		return err
	}
	teamWins[game.Team1.ID] += wins
	teamGamesPlayed[game.Team1.ID] += played
	return nil
}

func getPlayedAndWins(teamId int, standings *cfl.Standings) (int, int, error) {
	if teamId == 0 {
		return 0, 0, nil
	}

	for _, team := range standings.Data.Divisions.West.Standings {
		if teamId == team.ID {
			return team.GamesPlayed, team.Wins, nil
		}
	}
	for _, team := range standings.Data.Divisions.East.Standings {
		if teamId == team.ID {
			return team.GamesPlayed, team.Wins, nil
		}
	}
	return 0, 0, fmt.Errorf("team ID %d not found", teamId)
}

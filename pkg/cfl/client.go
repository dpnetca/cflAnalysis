package cfl

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	baseUrl string
	apiKey  string
}

func NewClient(apiKey string) Client {
	return Client{
		baseUrl: "http://api.cfl.ca/v1/",
		apiKey:  apiKey,
	}
}

func (c Client) execute(url string, d any) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	if res.StatusCode/100 != 2 {

		return fmt.Errorf("http error %s", res.Status)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(d)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) GetAllGames(year string) (Season, error) {
	var season Season
	url := c.baseUrl + "games/" + year + "?key=" + c.apiKey
	err := c.execute(url, &season)
	if err != nil {
		return season, err
	}

	return season, nil
}

func (c Client) GetStandings(year string) (Standings, error) {
	var standings Standings
	url := c.baseUrl + "standings/" + year + "?key=" + c.apiKey
	err := c.execute(url, &standings)
	if err != nil {
		return standings, err
	}
	return standings, nil

}
func (c Client) GetTeams() (Teams, error) {
	var teams Teams
	url := c.baseUrl + "teams/?key=" + c.apiKey
	err := c.execute(url, &teams)
	if err != nil {
		return teams, err
	}
	return teams, nil

}

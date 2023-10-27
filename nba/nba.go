package nba

import (
	"encoding/json"
	"fantasy-reminder-bot/models"
	"io"
	"net/http"
	"time"
)

func GetTodaysFirstGame() (error, *models.Game) {
	resp, err := http.Get("https://cdn.nba.com/static/json/liveData/scoreboard/todaysScoreboard_00.json")
	if err != nil {
		return err, nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}

	var data models.NBAData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err, nil
	}

	if len(data.Scoreboard.Games) == 0 {
		return nil, nil
	}

	return nil, &data.Scoreboard.Games[0]
}

func TimeUntilGame(game *models.Game) (error, time.Duration) {
	gameTime, err := time.Parse(time.RFC3339, game.GameTimeUTC)
	if err != nil {
		return err, 0
	}

	return nil, gameTime.Sub(time.Now().UTC())
}

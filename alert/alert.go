package alert

import (
	"fantasy-reminder-bot/models"
	"fantasy-reminder-bot/nba"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var Cache map[string]models.AlertEvent

func Init() {
	Cache = make(map[string]models.AlertEvent)
}

func SendAlertIfNeeded(s *discordgo.Session, firstGame *models.Game) {
	nextAlert, ok := Cache[firstGame.GameID]

	if !ok {
		nextAlert = models.AlertEvent{
			Game:            *firstGame,
			AlertedAt5Hrs:   false,
			AlertedAt30Mins: false,
		}
	}

	if nextAlert.AlertedAt5Hrs && nextAlert.AlertedAt30Mins {
		return
	}

	err, timeUntilGame := nba.TimeUntilGame(firstGame)
	if err != nil {
		panic(err)
	}

	if timeUntilGame.Hours() < 5 && timeUntilGame.Minutes() > 30 {
		if !nextAlert.AlertedAt5Hrs {
			nextAlert.AlertedAt5Hrs = true
			durationAsString := fmt.Sprintf("@everyone The first game today (between %s and %s) is in: 5 hours. Don't forget to set your rosters, you worthless losers!", firstGame.HomeTeam.TeamName, firstGame.AwayTeam.TeamName)
			s.ChannelMessageSend("1167132859966967863", durationAsString)
		}
	}

	if timeUntilGame.Minutes() < 30 && timeUntilGame.Minutes() > 0 {
		if !nextAlert.AlertedAt30Mins {
			nextAlert.AlertedAt30Mins = true
			durationAsString := fmt.Sprintf("@everyone The first game today (between %s and %s) is in: 30 minutes. Don't forget to set your rosters, you worthless losers!", firstGame.HomeTeam.TeamName, firstGame.AwayTeam.TeamName)
			s.ChannelMessageSend("1167132859966967863", durationAsString)
		}
	}

	Cache[firstGame.GameID] = nextAlert
}

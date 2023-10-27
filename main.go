package main

import (
	"fantasy-reminder-bot/alert"
	"fantasy-reminder-bot/models"
	"fantasy-reminder-bot/nba"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token     string
	NextAlert models.AlertEvent
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	alert.Init()
}

func main() {
	fmt.Println(Token)

	if Token == "" {
		Token = "MTE2NzEzMzAyNjQ0NzI2NTkxMw.Gi0HOp.QPCDJhB1T_d19v6-MZqt1h-_6Ua7pqEJBVq0MM"
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	scheduleTask(dg)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func scheduleTask(s *discordgo.Session) {
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				err, firstGame := nba.GetTodaysFirstGame()
				if err != nil {
					panic(err)
				}

				if firstGame == nil {
					log.Print("No games today")
					continue
				}

				alert.SendAlertIfNeeded(s, firstGame)
			}
		}
	}()
}

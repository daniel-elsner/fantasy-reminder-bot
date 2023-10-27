package models

type AlertEvent struct {
	Game            Game
	AlertedAt5Hrs   bool
	AlertedAt30Mins bool
}

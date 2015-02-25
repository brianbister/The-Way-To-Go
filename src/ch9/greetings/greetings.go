package greetings

import "time"

func IsAM() bool {
	return time.Now().Hour() < 12
}
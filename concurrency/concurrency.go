package concurrency

import (
	"time"
)

func Count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	// as a sender, close the channel if you don't need it anymore
	// never close a channel as a receiver
	close(c)
}

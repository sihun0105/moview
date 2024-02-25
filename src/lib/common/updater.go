package common

import (
	"fmt"
	"time"
)


func StartUpdater(updateFunc func()) {
	ticker := time.NewTicker(24 * time.Hour)
	now := time.Now()
	firstTick := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.Local)
	if firstTick.Before(now) {
		firstTick = firstTick.Add(24 * time.Hour)
	}
	timeUntilFirstTick := firstTick.Sub(now)

	fmt.Println("Next update at:", firstTick)

	time.Sleep(timeUntilFirstTick)

	for {
		<-ticker.C
		fmt.Println("Updating movies...")
		updateFunc()
	}
}

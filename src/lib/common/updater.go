package common

import (
	"fmt"

	"github.com/robfig/cron/v3"
)


func StartUpdater(updateFunc func()) {
    c := cron.New(cron.WithSeconds())
    
    _, err := c.AddFunc("0 0 0 * * *", func() {
        fmt.Println("Updating movies...")
        updateFunc()
    })
    if err != nil {
        fmt.Println("Error scheduling update:", err)
        return
    }
    c.Start()
    fmt.Println("Cron scheduler started. Next update at midnight.")
}

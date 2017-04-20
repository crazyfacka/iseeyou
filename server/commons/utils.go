package commons

import (
	"fmt"
	"time"
)

// Debug prints debug information
func Debug(message interface{}, args ...interface{}) {
	if cfg.Debug {
		t := time.Now()
		if len(args) > 0 {
			fmt.Printf("["+t.Format("Mon Jan _2 15:04:05 2006")+"] "+message.(string)+"\n", args...)
		} else {
			if val, ok := message.(string); ok {
				fmt.Println("[" + t.Format("Mon Jan _2 15:04:05 2006") + "] " + val)
			} else {
				fmt.Println("[" + t.Format("Mon Jan _2 15:04:05 2006") + "]")
				fmt.Println(message)
				fmt.Println("== DEBUG-END ===")
			}
		}
	}
}

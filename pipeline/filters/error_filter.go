package filters

import (
	"fmt"

	"github.com/jordanmarta/software-architecture-labs/pipeline/model"
)

func ErrorFilter(input <-chan model.LogEntry) <-chan model.LogEntry {
	out := make(chan model.LogEntry)

	go func() {
		defer close(out)

		for log := range input {
			if log.Level == "ERROR" {
				fmt.Println("[ErrorFilter] accepted:", log.Message)
				out <- log
			} else {
				fmt.Println("[ErrorFilter] discarded:", log.Level, "-", log.Message)
			}
		}
	}()

	return out
}

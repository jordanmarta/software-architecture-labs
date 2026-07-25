package filters

import (
	"fmt"
	"strings"
	"time"

	"github.com/jordanmarta/software-architecture-labs/pipeline/model"
)

func Parser(input <-chan string) <-chan model.LogEntry {
	out := make(chan model.LogEntry)

	go func() {
		defer close(out)

		for rawLog := range input {
			parts := strings.SplitN(rawLog, "|", 3)

			if len(parts) != 3 {
				continue
			}

			timestamp, err := time.Parse(time.RFC3339, parts[0])
			if err != nil {
				continue
			}

			fmt.Printf("[Parser] parsed: %s | %s | %s\n",
				timestamp.Format(time.RFC3339),
				parts[1],
				parts[2],
			)

			out <- model.LogEntry{
				Timestamp: timestamp,
				Level:     parts[1],
				Message:   parts[2],
			}
		}
	}()

	return out
}

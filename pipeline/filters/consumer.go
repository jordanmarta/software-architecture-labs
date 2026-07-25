package filters

import (
	"fmt"

	"github.com/jordanmarta/software-architecture-labs/pipeline/model"
)

// Consumer representa o estágio final do pipeline.
// Em um sistema real, poderia persistir o log ou enviá-lo para outro destino.
func Consumer(input <-chan model.LogEntry) {
	for log := range input {
		fmt.Printf(
			"[%s] %s: %s\n",
			log.Timestamp.Format("2006-01-02 15:04:05"),
			log.Level,
			log.Message)
	}
}

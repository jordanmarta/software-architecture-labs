package main

import "github.com/jordanmarta/software-architecture-labs/pipeline/filters"

func main() {
	logs := []string{
		"2026-07-25T17:30:00Z|INFO|application started",
		"2026-07-25T17:30:04Z|INFO|database connection established",
		"2026-07-25T17:30:11Z|INFO|payment received",
		"2026-07-25T17:30:13Z|WARN|payment processing is taking longer than expected",
		"2026-07-25T17:30:18Z|INFO|payment approved",
		"2026-07-25T17:31:02Z|INFO|payment received",
		"2026-07-25T17:31:05Z|ERROR|database unavailable",
		"2026-07-25T17:31:08Z|WARN|retrying database connection",
		"2026-07-25T17:31:12Z|INFO|database connection restored",
		"2026-07-25T17:32:20Z|INFO|payment received",
		"2026-07-25T17:32:25Z|ERROR|payment processing timeout",
		"2026-07-25T17:33:00Z|INFO|health check completed",
	}

	produced := filters.Producer(logs)
	parsed := filters.Parser(produced)
	errorsOnly := filters.ErrorFilter(parsed)

	filters.Consumer(errorsOnly)
}

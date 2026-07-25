package filters

import "fmt"

func Producer(logs []string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for _, log := range logs {
			fmt.Println("[Producer] sending:", log)
			out <- log
		}
	}()

	return out
}

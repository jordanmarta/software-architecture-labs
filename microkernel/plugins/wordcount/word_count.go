package wordcount

import (
	"fmt"
	"strings"

	"github.com/jordanmarta/software-architecture-labs/microkernel/core"
)

type WordCountPlugin struct{}

func (p WordCountPlugin) Name() string {
	return "word-count"
}

func (p WordCountPlugin) OnEvent(
	event core.Event,
	document *core.Document,
) core.Result {

	if event != core.EventTextChange {
		return core.Result{}
	}

	count := len(strings.Fields(document.Content))

	return core.Result{
		Message: fmt.Sprintf("%d words", count),
	}
}

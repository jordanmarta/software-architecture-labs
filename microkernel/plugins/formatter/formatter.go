package formatter

import (
	"strings"

	"github.com/jordanmarta/software-architecture-labs/microkernel/core"
)

type FormatterPlugin struct{}

func (p FormatterPlugin) Name() string {
	return "formatter"
}

func (p FormatterPlugin) OnEvent(
	event core.Event,
	document *core.Document,
) core.Result {

	if event != core.EventSave {
		return core.Result{}
	}

	document.Content = strings.TrimSpace(document.Content)

	return core.Result{
		Message: "document formatted before save",
	}
}

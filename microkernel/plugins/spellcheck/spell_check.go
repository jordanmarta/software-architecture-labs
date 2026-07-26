package spellcheck

import (
	"fmt"
	"strings"

	"github.com/jordanmarta/software-architecture-labs/microkernel/core"
)

type SpellCheckPlugin struct{}

func (p SpellCheckPlugin) Name() string {
	return "spell-check"
}

func (p SpellCheckPlugin) OnEvent(
	event core.Event,
	document *core.Document,
) core.Result {

	if event != core.EventTextChange {
		return core.Result{}
	}

	knownErrors := []string{
		"arquetetura",
		"microkenel",
		"plguin",
	}

	var errors []string

	for _, word := range knownErrors {
		if strings.Contains(strings.ToLower(document.Content), word) {
			errors = append(errors, word)
		}
	}

	if len(errors) == 0 {
		return core.Result{
			Message: "no spelling issues found",
		}
	}

	return core.Result{
		Message: fmt.Sprintf("possible spelling issues: %s", strings.Join(errors, ", ")),
	}
}

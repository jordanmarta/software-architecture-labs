package markdown

import (
	"fmt"
	"strings"

	"github.com/jordanmarta/software-architecture-labs/microkernel/core"
)

type MarkdownPreviewPlugin struct{}

func (p MarkdownPreviewPlugin) Name() string {
	return "markdown-preview"
}

func (p MarkdownPreviewPlugin) OnEvent(
	event core.Event,
	document *core.Document,
) core.Result {

	if event != core.EventTextChange {
		return core.Result{}
	}

	if !strings.HasSuffix(document.Name, ".md") {
		return core.Result{}
	}

	html := document.Content

	// Conversão extremamente simplificada apenas para demonstrar o plugin.
	if strings.HasPrefix(html, "# ") {
		html = strings.Replace(html, "# ", "<h1>", 1)
		html = fmt.Sprintf("%s</h1>", html)
	}

	return core.Result{
		Message: "preview generated",
		Output:  html,
	}
}

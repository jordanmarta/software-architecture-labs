package main

import (
	"github.com/jordanmarta/software-architecture-labs/microkernel/core"
	"github.com/jordanmarta/software-architecture-labs/microkernel/plugins/formatter"
	"github.com/jordanmarta/software-architecture-labs/microkernel/plugins/markdown"
	"github.com/jordanmarta/software-architecture-labs/microkernel/plugins/spellcheck"
	"github.com/jordanmarta/software-architecture-labs/microkernel/plugins/wordcount"
)

func main() {
	editor := core.NewEditor()

	editor.RegisterPlugin(wordcount.WordCountPlugin{})
	editor.RegisterPlugin(spellcheck.SpellCheckPlugin{})
	editor.RegisterPlugin(markdown.MarkdownPreviewPlugin{})
	editor.RegisterPlugin(formatter.FormatterPlugin{})

	editor.Open(
		"README.md",
		"# Microkernel Architecture",
	)

	editor.Show()

	editor.Edit(
		"# Microkenel Architecture\n\nThis plguin demonstrates arquetetura.\n\n",
	)

	editor.Show()

	editor.Save()

	editor.Show()
}

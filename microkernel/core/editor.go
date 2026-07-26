package core

import "fmt"

type Editor struct {
	document Document
	plugins  map[string]Plugin
}

func NewEditor() *Editor {
	return &Editor{
		plugins: make(map[string]Plugin),
	}
}

func (e *Editor) RegisterPlugin(plugin Plugin) {
	e.plugins[plugin.Name()] = plugin
}

func (e *Editor) Open(name, content string) {
	e.document = Document{
		Name:    name,
		Content: content,
	}

	fmt.Printf("[Core] Opening document: %s\n", name)

	e.notifyPlugins(EventOpen)
}

func (e *Editor) Edit(content string) {
	e.document.Content = content

	fmt.Println("[Core] Document changed")

	e.notifyPlugins(EventTextChange)
}

func (e *Editor) Save() {
	e.notifyPlugins(EventSave)

	fmt.Printf("[Core] Saving document: %s\n", e.document.Name)
}

func (e *Editor) Show() {
	fmt.Println("\n--- Document ---")
	fmt.Println(e.document.Content)
	fmt.Println("----------------")
}

func (e *Editor) notifyPlugins(event Event) {
	for _, plugin := range e.plugins {
		result := plugin.OnEvent(event, &e.document)

		if result.Message != "" {
			fmt.Printf("[%s] %s\n", plugin.Name(), result.Message)
		}

		if result.Output != "" {
			fmt.Printf("[%s]\n%s\n", plugin.Name(), result.Output)
		}
	}
}

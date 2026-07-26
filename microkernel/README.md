# Microkernel / Plugin Architecture

Este lab demonstra o estilo **Microkernel**, também chamado de **Plugin Architecture**.

A ideia é manter um **Core pequeno e estável**, enquanto funcionalidades adicionais são implementadas como plugins.

## Estrutura

```text
microkernel/
├── main.go
├── core/
│   ├── document.go
│   ├── editor.go
│   ├── plugin.go
│   └── result.go
└── plugins/
    ├── formatter/
    ├── markdown/
    ├── spellcheck/
    └── wordcount/
```

## Exemplo

O `Editor` representa o Core e possui funcionalidades básicas:

```text
Open
Edit
Save
Show
```

Os plugins adicionam novas capacidades:

```text
WordCount       → conta palavras
SpellCheck      → procura erros
MarkdownPreview → gera preview
Formatter       → formata antes de salvar
```

O Core conhece apenas o contrato:

```go
type Plugin interface {
	Name() string
	OnEvent(event Event, document *Document) Result
}
```

## Pontos de extensão

O Core dispara eventos:

```go
const (
	EventOpen       Event = "open"
	EventTextChange Event = "text_change"
	EventSave       Event = "save"
)
```

Fluxo:

```text
TextChange
├── WordCount
├── SpellCheck
└── MarkdownPreview

Save
└── Formatter
```

Cada plugin decide a quais eventos reage.

## Registry

Os plugins são registrados no Core:

```go
editor.RegisterPlugin(wordcount.WordCountPlugin{})
editor.RegisterPlugin(spellcheck.SpellCheckPlugin{})
editor.RegisterPlugin(markdown.MarkdownPreviewPlugin{})
editor.RegisterPlugin(formatter.FormatterPlugin{})
```

O `Editor` não conhece as implementações concretas.

## Por que é Microkernel?

Porque existe:

```text
Core estável
     ↓
Contrato de extensão
     ↓
Plugins independentes
```

O sistema funciona sem os plugins, mas pode ser estendido sem alterar o Core.

Isso diferencia Microkernel de Strategy:

```text
Strategy    → diferentes formas de executar a mesma operação
Microkernel → novas capacidades adicionadas ao sistema
```

Os eventos lembram Observer, mas aqui são apenas o mecanismo usado para implementar os pontos de extensão.

## Vantagens

- extensibilidade;
- isolamento das funcionalidades;
- Core mais estável;
- possibilidade de diferentes conjuntos de plugins.

## Trade-offs

- contrato entre Core e plugins precisa ser estável;
- plugins podem aumentar a complexidade;
- dependências entre plugins devem ser evitadas;
- debugging pode ficar mais difícil em sistemas grandes.

## Quando usar

Quando existe um núcleo relativamente estável e funcionalidades que precisam variar ou ser adicionadas separadamente.

Exemplos:

```text
IDEs
editores
browsers
ferramentas de build
engines de regras
```

## Executando

```bash
go run ./microkernel
```

## Principal aprendizado

O Core não precisa conhecer todas as funcionalidades existentes ou futuras.

Ele precisa apenas definir **como extensões podem se conectar a ele**.
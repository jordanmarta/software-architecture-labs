# Software Architecture Labs

Repositório de estudos práticos sobre **Arquitetura de Software**, criado para transformar conceitos teóricos em pequenos exemplos de código.

O estudo atual acompanha o livro:

> **Fundamentos da Arquitetura de Software: Uma Abordagem de Engenharia**  
> **Autores:** Mark Richards e Neal Ford

## Objetivo

A ideia deste repositório não é construir aplicações completas.

Cada lab busca demonstrar de forma simples:

- Como um estilo arquitetural é estruturado.
- Quais são seus principais componentes.
- Como esses componentes se comunicam.
- Quais responsabilidades existem.
- Vantagens e trade-offs.
- Quando aquele estilo faz sentido.

Os exemplos são mantidos propositalmente pequenos para que a arquitetura seja o foco principal.

## Tecnologias

Os labs são desenvolvidos principalmente em **Go**.

Infraestrutura, bancos de dados e frameworks só serão utilizados quando realmente ajudarem a demonstrar algum conceito arquitetural.

## Labs

```text
software-architecture-labs/
├── layered/
├── pipeline/
├── microkernel/
├── service-based/
└── event-driven/
```

Cada arquitetura utiliza o domínio que melhor ajuda a demonstrar suas características, evitando adaptar artificialmente o mesmo problema para todos os estilos.

## Estilos estudados

- [x] Layered Architecture
- [ ] Pipeline Architecture
- [ ] Microkernel Architecture
- [ ] Service-Based Architecture
- [ ] Event-Driven Architecture

## Estrutura de estudo

Cada lab procura seguir a mesma linha:

```text
Conceito
   ↓
Estrutura
   ↓
Código
   ↓
Fluxo
   ↓
Vantagens e Trade-offs
```

O objetivo final é construir uma coleção de exemplos pequenos que possam servir tanto como material de revisão quanto como referência prática sobre diferentes estilos de arquitetura de software.

## Escopo futuro

Este repositório não ficará limitado ao livro atual.

Novos estilos, padrões e conceitos de arquitetura poderão ser adicionados conforme o avanço dos estudos em outros livros e materiais.
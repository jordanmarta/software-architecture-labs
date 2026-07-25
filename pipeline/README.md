# Pipeline Architecture

Lab simples para demonstrar o estilo de **Arquitetura Pipeline** (ou **Pipes-and-Filters**), baseado nos conceitos apresentados em *Fundamentos da Arquitetura de Software*.

## O que caracteriza esse estilo

A arquitetura pipeline é formada por:

- **Filtros**: etapas que executam uma responsabilidade específica.
- **Canais**: meios pelos quais os dados passam de um filtro para outro.

Cada filtro recebe uma entrada, faz um trabalho específico e encaminha o resultado para o próximo estágio.

## Estrutura do lab

Neste lab, o pipeline processa linhas de log.

```text
Producer → Parser → ErrorFilter → Consumer
```

## Fluxo da aplicação

A entrada do sistema é uma lista de logs em formato texto:

```text
2026-07-25T17:30:00Z|INFO|application started
2026-07-25T17:31:05Z|ERROR|database unavailable
2026-07-25T17:32:25Z|ERROR|payment processing timeout
```

Fluxo:

1. **Producer** envia os logs crus para o pipeline.
2. **Parser** transforma cada linha em uma estrutura `LogEntry`.
3. **ErrorFilter** deixa passar apenas logs com nível `ERROR`.
4. **Consumer** consome o resultado final e exibe os logs filtrados.

## Principais componentes

### Producer

Arquivo:

`filters/producer.go`

Responsável por iniciar o fluxo, enviando as strings de log para o primeiro canal.

### Parser

Arquivo:

`filters/parser.go`

Responsável por transformar a string bruta em um objeto estruturado:

- timestamp
- level
- message

### ErrorFilter

Arquivo:

`filters/error_filter.go`

Responsável por verificar se o log deve continuar no pipeline.

Neste exemplo, apenas logs com `Level == "ERROR"` seguem adiante.

### Consumer

Arquivo:

`filters/consumer.go`

Responsável por consumir o resultado final do pipeline.

Neste lab, ele apenas exibe os logs. Em um cenário real, poderia gravar em banco, arquivo ou enviar para outro sistema.

## Correlação com o livro

No livro, os papéis principais aparecem como:

- **Producer**
- **Tester / Verificador**
- **Transformer / Transformador**
- **Consumer**

No nosso lab, a correlação é:

```text
Livro                     Nosso lab
Producer                  Producer
Tester / Verificador      ErrorFilter
Transformer               Parser
Consumer                  Consumer
```

## Observação importante sobre a ordem

No exemplo conceitual do livro, os papéis podem aparecer assim:

```text
Producer → Tester → Transformer → Consumer
```

No nosso lab, usamos:

```text
Producer → Parser → ErrorFilter → Consumer
```

Isso **não é um problema**.

Os papéis do pipeline não têm uma ordem fixa obrigatória.  
A ordem depende do processamento necessário.

No nosso caso, faz mais sentido:

1. **transformar primeiro** a string em `LogEntry`;
2. **verificar depois** se `Level == "ERROR"`.

## Canais e filtros no nosso lab

A ideia central do pipeline aparece claramente aqui:

```text
[]string
   ↓
Producer
   ↓ chan string
Parser
   ↓ chan LogEntry
ErrorFilter
   ↓ chan LogEntry
Consumer
```

Cada estágio conhece apenas:

- o que recebe;
- o que devolve.

Isso reduz acoplamento e facilita a troca ou evolução dos filtros.

## Vantagens

- Estrutura simples e fácil de entender.
- Cada filtro tem uma responsabilidade clara.
- Filtros podem ser alterados de forma isolada.
- Boa modularidade.
- Fácil de adicionar, remover ou reordenar etapas.

## Trade-offs

- Baixa elasticidade e baixa escalabilidade quando comparada a arquiteturas distribuídas.
- Tolerância a falhas limitada em implementações monolíticas.
- Pode não ser adequada para fluxos muito complexos ou com muitas dependências entre etapas.
- O sistema inteiro ainda pode precisar ser implantado junto, dependendo da implementação.

## Quando usar

Faz sentido quando o problema pode ser dividido em etapas bem definidas de processamento, como:

- processamento de logs;
- transformação de arquivos;
- ETL;
- validação e enriquecimento de dados;
- processamento sequencial de informações.

## Quando evitar

Pode não ser a melhor escolha quando:

- as etapas têm dependências muito fortes entre si;
- o fluxo não é naturalmente sequencial;
- há necessidade forte de escalabilidade independente;
- o sistema precisa de alta tolerância a falhas.

## Como executar

```bash
go run ./pipeline
```

Saída esperada:

```text
[2026-07-25 17:31:05] ERROR: database unavailable
[2026-07-25 17:32:25] ERROR: payment processing timeout
```
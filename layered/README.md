# Layered Architecture

Lab simples para demonstrar o estilo de **Arquitetura em Camadas**, baseado nos conceitos apresentados em *Fundamentos de Arquitetura de Software*.

## Estrutura

```text
Presentation
    ↓
Business
    ↓
Persistence
```

Neste exemplo usamos um fluxo simples de criação de pedidos.

## Fluxo

1. `Presentation` recebe os dados do pedido.
2. `Business` valida o pedido e calcula o valor total.
3. `Persistence` simula a gravação do pedido.

```text
OrderHandler
    ↓
OrderService
    ↓
OrderRepository
```

## Responsabilidades

### Presentation Layer

Responsável pela entrada e saída da aplicação.

Arquivo:

`presentation/order_handler.go`

Cria o pedido, chama a camada de negócio e apresenta o resultado.

### Business Layer

Responsável pelas regras de negócio.

Arquivo:

`business/order_service.go`

Regras implementadas:

- O pedido deve possuir pelo menos um item.
- O valor total do pedido é calculado nessa camada.

### Persistence Layer

Responsável pelo acesso aos dados.

Arquivo:

`persistence/order_repository.go`

Não existe banco real neste lab. A persistência é apenas simulada para manter o foco na estrutura arquitetural.

## Layer Isolation

Neste exemplo utilizamos **camadas fechadas**.

Cada camada conhece apenas a camada imediatamente abaixo:

```text
Presentation → Business → Persistence
```

A camada de apresentação não acessa diretamente a persistência.

## Architecture Sinkhole

Um problema comum da arquitetura em camadas é o **Architecture Sinkhole Anti-Pattern**.

Ele acontece quando uma requisição atravessa várias camadas sem que elas executem trabalho relevante.

Exemplo:

```go
func (s OrderService) CreateOrder(order model.Order) error {
	return s.repository.Save(order)
}
```

Nesse caso, a camada de negócio existe, mas apenas repassa a chamada.

Como regra prática apresentada no livro:

- Aproximadamente **80% dos fluxos** devem utilizar realmente as responsabilidades das camadas.
- Até aproximadamente **20% dos fluxos** podem ser sinkhole sem necessariamente representar um problema.

Se grande parte das requisições apenas atravessa as camadas, isso pode indicar que a arquitetura está adicionando estrutura sem gerar valor e talvez outro estilo arquitetural seja mais adequado.

## Vantagens

- Estrutura simples e fácil de entender.
- Separação clara de responsabilidades.
- Baixo custo inicial.
- Familiar para a maioria dos desenvolvedores.

## Trade-offs

- Pode gerar acoplamento entre camadas.
- Mudanças podem atravessar várias partes da aplicação.
- Pode gerar código apenas para repassar chamadas.
- Existe risco de Architecture Sinkhole.
- Escalabilidade e elasticidade são limitadas quando comparadas a arquiteturas distribuídas.

## Quando usar

Faz sentido quando o sistema possui responsabilidades que naturalmente podem ser divididas em apresentação, regras de negócio e persistência.

É uma boa opção quando simplicidade, organização e baixo custo são prioridades.

## Quando evitar

Pode não ser uma boa escolha quando:

- A maior parte das camadas apenas repassa chamadas.
- Existe necessidade forte de escalabilidade independente.
- Componentes precisam ser implantados ou evoluir separadamente.
- O domínio não possui responsabilidades suficientes para justificar várias camadas.
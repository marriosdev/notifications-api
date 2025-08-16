```shell
# Estrutura 
│
├── cmd/
│   └── api/
│       └── main.go            # Ponto de entrada, boot da aplicação
│
├── internal/
│   ├── api/
│   │   ├── handler.go         # HTTP Handlers (chamam services)
│   │   └── router.go          # Configuração de rotas
│   │
│   ├── config/
│   │   └── config.go          # Load de config (sem init problemático)
│   │
│   ├── domain/
│   │   └── usuario/
│   │       ├── model.go       # Structs com campos exportados, sem getters bobos
│   │       ├── service.go     # Lógica de negócio, sem aninhamento exagerado
│   │       └── repository.go  # Acesso a dados
│   │
│   ├── middleware/            # CORS, JWT, logging, etc.
│   │   └── auth.go
│   │
│   └── util/                  # Funções reutilizáveis (pequenas e claras)
│       └── linter_integration.go
│
├── pkg/                       # Bibliotecas reutilizáveis (se surgir)
│
├── .golangci.yml              # Configuração de linters consistente
├── go.mod
├── go.sum
└── README.md
```

cmd/api/main.go — lançamento explícito da aplicação. Evita lógica escondida no init.

config — carrega/envia configs via função, capturando erros; nada de init() silencioso.

domain/usuario — cada entidade em seu pacote: sem colisão de nomes, sem códigos perdidos.

model.go — structs explícitos exportados. Nada de GetName() se bastaria user.Name.

service.go — lógica de negócio isolada, usando early return pra reduzir nesting.

repository.go — abstração sobre armazenamento, interface deve ser definida no service (consumidor), não no repo.

Handlers — apenas coordenação; sem regras de negócio nem lógica de DB.

Linters configurados — garantem código limpo e prevenção automática de sombras, erros comuns etc.
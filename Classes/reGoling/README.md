# reGoling - Reaprendendo Golang

## Sobre

### Plano Geral

Plano de para estudo da Golang para quem já conhece mas faz algum tempo que não utiliza a linguagem.
Proposta para 10 dias, sendo 30-60min/dia de estudo.

### Filosofia

- Foco em código que funciona e resolve problemas, não em teoria pura
- Ênfase em padrões que você verá em ferramentas de infraestrutura (kubectl, helm, operators)
- Pragmatismo sobre perfeccionismo - você vai usar GO como ferramenta, não como linguagem principal

### Mapa de Conceitos Essenciais

1. Estrutura de projeto e módulos - Como organizar código e gerenciar dependências
1. Pacotes e exportação - A base de qualquer programa em GO
1. Tipos básicos, structs e interfaces - O coração da modelagem de dados
1. Goroutines e channels - O que torna GO especial para infraestrutura
1. Manipulação de erros - Diferente de outras linguagens, crucial para sistemas robustos
1. Context - Essencial para serviços que rodam em Kubernetes
1. JSON/YAML e I/O - Você vai lidar com configurações o tempo todo
1. Testes básicos - Para validar suas automações

### Estrutura Diária

- **DIA 1**: Setup e Visão Geral
- **DIA 2**: Pacotes e Exportação
- **DIA 3**: Tipos e Structs
- **DIA 4**: Interfaces
- **DIA 6**: Goroutines e Channels (Parte 2)
- **DIA 7**: Tratamento de Erros
- **DIA 8**: Context e Timeouts
- **DIA 9**: JSON/YAML e I/O
- **DIA 10**: Testes e Integração com K8s

### Conceitos Não Cobertos

Para futuros estudos sugere-se:

- Generics (GO 1.18+) - útil mas não crítico
- Reflect - avançado, pouco usado
- GO GC e performance tuning
- CGO e chamadas C
- Plugins e build tags
- Profiling e trace

### Recursos para seu contexto específico

Bibliotecas que você vai usar:

- k8s.io/client-go - Cliente oficial Kubernetes
- sigs.k8s.io/controller-runtime - Para operadores
- k8s.io/apimachinery - Tipos K8s
- github.com/spf13/cobra - CLI apps
- github.com/sirupsen/logrus - Logging estruturado

## DIA 1: Setup e Visão Geral

### Proposta

**Conceitos**:

- Módulos (go.mod, go.sum)
- GOPATH vs módulos (não use mais GOPATH)
- Estrutura de diretórios padrão
- Ferramentas: go run, go build, go mod tidy

**Analogia**: Módulos são como packages no Maven/NPM - definem dependências e versões.

**Aplicação prática**:

```bash
mkdir meu-projeto && cd meu-projeto
go mod init github.com/seu-user/meu-projeto
```

**Erro comum**: Tentar colocar código fora do módulo ou sem go.mod.

**Exercício**:

Crie um programa que imprima "Hello, K8s!" usando módulos. Compile e execute.

**Para aprofundar**: [Documentação oficial de módulos](https://go.dev/doc/modules/managing-dependencies)

### Instalação

Seguir roteiro em [GO: Download and install](https://go.dev/doc/install).

### Entendendo Módulos

Em Go, um módulo é simplesmente uma coleção de pacotes Go com um arquivo go.mod na raiz. Pense nele como:

- O `pom.xml` do **Maven** ou o `package.json` do **NPM** para Go.
- O `go.mod` declara o nome do seu módulo e lista as dependências externas (bibliotecas) que seu código precisa.
- O `go.sum` é um arquivo de bloqueio que contém hashes criptográficos das dependências, garantindo que todos que baixarem seu projeto usem exatamente as mesmas versões.

**O que mudou?** Antigamente (pré-1.11) usávamos o `GOPATH`, que forçava uma estrutura de pastas rígida. **Com módulos, você pode trabalhar em qualquer diretório do seu computador.** Esqueça o `GOPATH`!

### Criando Seu Primeiro Módulo

Vamos colocar a mão na massa. Abra seu terminal e siga:

```sh
# 1. Crie uma pasta para o projeto (pode ser em qualquer lugar)
mkdir -p dia1/t1-hellok8s && cd dia1/t1-hellok8s/

# 2. Inicie o módulo. O nome é importante: use um caminho que reflita onde ele ficará (ex: github.com/seu-usuario/hellok8s)
go mod init github.com/tarsoqueiroz/regod1t1-hellok8s

# 3. Veja o arquivo criado
cat go.mod
```

O go.mod deve aparecer assim:

```text
module github.com/tarsoqueiroz/regod1t1-hellok8s

go 1.26
```

> **Dica importante**: Escolha um nome de módulo que você acredite que será único. Se for um projeto pessoal, pode usar algo como `meuprojeto` ou `exemplo/hellok8s`. O importante é que não vai conflitar com módulos públicos.

### Escrevendo o Código

Crie um arquivo `main.go` com o conteúdo:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, K8s!")
}
```

### Compilando e Executando

Agora, veja a mágica dos módulos:

```sh
# Execute diretamente (sem compilar)
go run main.go

# Compile para um binário executável
go build -o hellok8s main.go

# Execute o binário compilado (no Linux/Mac)
./hellok8s
# Ou no Windows: hellok8s.exe
```

### Entendendo a Estrutura de Pastas

A estrutura de um projeto Go comum que você verá em ferramentas de K8s é:

```text
meu-projeto/
├── go.mod          # Declaração do módulo e dependências
├── go.sum          # Checksums das dependências (segurança)
├── main.go         # Ponto de entrada
├── pkg/            # Código que pode ser importado por outros
│   └── k8s/
│       └── client.go
└── cmd/            # Executáveis (se houver mais de um)
    └── myapp/
        └── main.go
```

### Adicionando uma Dependência

Agora vamos fazer algo mais útil: importar uma biblioteca externa. Por exemplo, a biblioteca para manipular YAML, muito usada com Kubernetes.

```sh
# Adicione a dependência
go get gopkg.in/yaml.v3
```

Observe:

- O `go.mod` foi atualizado com uma nova linha `require`
- O `go.sum` foi criado/atualizado com o hash da biblioteca

Execute novamente `go run main.go` - tudo continua funcionando.

### Erros e Confusões Comuns

| Erro | Por que acontece? | Como resolver? |
| :--- | :---------------- | :------------- |
| `go: cannot find main module` | Você está executando `go` em um diretório sem `go.mod` | Crie o módulo com `go mod init` |
| `go: module example.com/... found, but does not contain package` | Você tentou importar um submódulo que não existe | Verifique o caminho do pacote na documentação |
| `invalid version: unknown revision` | A versão especificada não existe | Use `@latest` ou verifique as tags no repositório |
| Esquecer de adicionar `go.mod` no repositório | Outros desenvolvedores não saberão as dependências | Sempre comite `go.mod` e `go.sum` |

### Exercício para Fixar

**Objetivo**: Criar um programa que use uma dependência externa para ler e exibir uma configuração em YAML.

- Crie um novo módulo chamado `exercicio-dia1`.
- Adicione a dependência `gopkg.in/yaml.v3`.
- Crie uma estrutura (`struct`) `Config` com campos `Timeout` (`int`) e Retries (`int`).
- Crie uma variável com dados **YAML** simulando uma configuração:

```yaml
timeout: 30
retries: 5
```

- Use a biblioteca **YAML** para "parsear" (desserializar) o YAML para a struct.
- Imprima os valores formatados.

> **Dica**: A função para desserializar é `yaml.Unmarshal([]byte(yamlString), &config)`.

**Solução**:

```sh
# criar pasta para o projeto
mkdir -p dia1/exerc-d1 && cd dia1/exerc-d1

# inicializar o modulo
go mod init github.com/tarsoqueiroz/regod1e1-exercicio

# adicionar dependência
go get gopkg.in/yaml.v3

# criar arquivo do app
touch main.go
```

- `main.go`

```go
package main

import (
    "fmt"
    "gopkg.in/yaml.v3"
)

type Config struct {
    Timeout int `yaml:"timeout"`
    Retries int `yaml:"retries"`
}

func main() {
    yamlData := `
timeout: 30
retries: 5
`
    var config Config
    err := yaml.Unmarshal([]byte(yamlData), &config)
    if err != nil {
        fmt.Println("Erro ao parsear YAML:", err)
        return
    }
    fmt.Printf("Configuração carregada: Timeout=%d, Retries=%d\n", config.Timeout, config.Retries)
}
```

```sh
# executando o código
go run main.go
```

### O que Estudar para Aprofundar

- **Documentação oficial**: [Managing dependencies](https://go.dev/doc/modules/managing-dependencies) - especialmente a seção sobre `replace` para usar versões locais de módulos (muito útil quando você está desenvolvendo e testando localmente).
- **Comandos úteis**: 
  - `go mod tidy`: limpa dependências não usadas
  - `go mod vendor`: cria uma cópia local das dependências
- **Proxy e privacidade**: A seção sobre `GOPROXY` e `GOPRIVATE` no link acima - essencial para empresas que usam repositórios privados.

### Checklist de Conclusão do Dia 1

- □ Criei um módulo com `go mod init`
- □ Escrevi um programa simples e executei com `go run`
- □ Compilei um binário com `go build`
- □ Adicionei uma dependência com `go get`
- □ Entendi que `go.mod` lista as dependências e `go.sum` garante integridade
- □ Completei o exercício do YAML

## DIA 2: Pacotes e Exportação

### Proposta

**Conceitos**:

- Pacotes = pastas
- Exportação = primeira letra maiúscula
- main é especial - ponto de entrada
- Importação de pacotes

**Analogia**: Pacotes são como namespaces em C# ou packages em Java.

**Relação**: Organização de código afeta diretamente como você usará bibliotecas de terceiros.

Aplicação prática:

```txt
meu-projeto/
├── main.go
├── k8s/
│   └── helpers.go  // pacote k8s
└── utils/
    └── strings.go   // pacote utils
```

**Erro comum**: Esquecer que apenas funções com letra maiúscula são exportadas.

**Exercício**:

Crie pacote k8s com função GetPodName() (exportada) e parseYaml() (não exportada). Use no main.

**Para aprofundar**: Organização de pacotes em GO

### 1. A Estrutura de Pacotes: A Base da Organização (10 min)

Em Go, a organização é simples e direta: cada diretório é um pacote. O nome do pacote é definido pela primeira linha do arquivo, com package nome.















## DIA 3: Tipos e Structs

### Proposta

Conceitos:

Tipos básicos: string, int, bool, float64

Structs = objetos (mas sem herança)

Métodos em structs

Ponteiros (* e &) - quando usar

Analogia: Struct é como uma classe em Java, mas só dados. Métodos são funções anexadas.

Relação: Structs são usados para modelar recursos Kubernetes (Pod, Service, etc).

Aplicação prática:

go
type Pod struct {
    Name      string
    Namespace string
    Ready     bool
}

func (p Pod) GetFullName() string {
    return p.Namespace + "/" + p.Name
}
Erro comum: Esquecer que GO passa tudo por valor (cópia) - use ponteiros para modificar.

Exercício:
Crie struct Deployment com campos Name, Replicas, Image. Adicione método Scale() que altera replicas.

Para aprofundar: Métodos e ponteiros

### 

## DIA 4: Interfaces

### Proposta

Conceitos:

Interface = contrato de métodos

Satisfação implícita (não precisa declarar "implements")

Interface vazia (interface{}) = any

Analogia: Interface é como uma API que você precisa implementar, mas sem declaração explícita.

Relação: Interfaces permitem desacoplar código - essencial para testabilidade.

Aplicação prática:

go
type K8sClient interface {
    GetPods(namespace string) ([]Pod, error)
    DeletePod(name string) error
}
Erro comum: Pensar que precisa declarar que implementa uma interface.

Exercício:
Crie interface Logger com método Log(message string). Implemente com FileLogger e ConsoleLogger.

Para aprofundar: Interfaces no GO

### 

## DIA 5: Goroutines e Channels (Parte 1)

### Proposta

Conceitos:

Goroutine = thread leve

go func() - inicia uma goroutine

WaitGroups para sincronização

Analogia: Goroutine é como uma thread em Java, mas muito mais leve (custa ~2KB).

Relação: O coração do modelo de concorrência em GO, usado em todos os operadores K8s.

Aplicação prática:

go
var wg sync.WaitGroup
for _, pod := range pods {
    wg.Add(1)
    go func(p Pod) {
        defer wg.Done()
        processPod(p)
    }(pod)
}
wg.Wait()
Erro comum: Esquecer de passar parâmetros para a goroutine, causando race conditions.

Exercício:
Crie programa que processa 1000 pods em paralelo usando goroutines (limite a 10 simultâneas).

### 

## DIA 6: Goroutines e Channels (Parte 2)

### Proposta

Conceitos:

Channels = comunicação entre goroutines

Buffered vs unbuffered

Select para multiplexar

Range sobre channel

Analogia: Channel é como uma fila (queue) thread-safe.

Relação: Channels permitem comunicação segura sem locks.

Aplicação prática:

go
ch := make(chan string, 5) // buffered
go func() {
    ch <- "mensagem"
}()
msg := <-ch
Erro comum: Deadlock - enviar para channel sem receptor ou vice-versa.

Exercício:
Faça pipeline: produtor → processador → consumidor usando channels.

Para aprofundar: Padrões de concorrência

### 

## DIA 7: Tratamento de Erros

### Proposta

Conceitos:

error é uma interface

Retorno múltiplo (valor, err)

Errors.Is e errors.As para wrapping

Panic/recover - NÃO USE em código normal

Analogia: Diferente de exceptions (try/catch), erros são valores de retorno.

Relação: Crucial para serviços long-running em K8s - erros não podem quebrar tudo.

Aplicação prática:

go
result, err := doSomething()
if err != nil {
    return fmt.Errorf("processando pod %s: %w", podName, err)
}
Erro comum: Ignorar erros com _ - SEMPRE trate ou propague.

Exercício:
Crie função que tenta conectar a um servidor (mock), retry com backoff exponencial.

Para aprofundar: Erros em GO

### 

## DIA 8: Context e Timeouts

### Proposta

Conceitos:

Context = contexto para timeouts, deadlines e cancelamentos

context.Background() vs context.TODO()

context.WithTimeout, WithCancel, WithDeadline

Analogia: Context é como um "passaporte" que carrega prazos e sinais de cancelamento.

Relação: Essencial para serviços em K8s que precisam responder a interrupções.

Aplicação prática:

go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

select {
case <-ctx.Done():
    return ctx.Err()
case result := <-longOperation():
    return result
}
Erro comum: Não propagar context para funções que fazem I/O.

Exercício:
Faça função que consulta API externa com timeout de 5s. Simule atraso e cancelamento.

Para aprofundar: Context na prática

### 

## DIA 9: JSON/YAML e I/O

### Proposta

Conceitos:

encoding/json e gopkg.in/yaml.v3

Tags: json:"field,omitempty"

io.Reader e io.Writer

ioutil (deprecated) vs os e io

Analogia: Marshal/Unmarshal = serialização/deserialização.

Relação: Você vai ler ConfigMaps e CRDs o tempo todo.

Aplicação prática:

go
type Config struct {
    Name    string `json:"name" yaml:"name"`
    Timeout int    `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}

data, _ := json.Marshal(config)
var newConfig Config
json.Unmarshal(data, &newConfig)
Erro comum: Esquecer tags para campos com nomes diferentes.

Exercício:
Leia arquivo YAML de configuração, modifique, e salve como JSON.

### 

## DIA 10: Testes e Integração com K8s

### Proposta

Conceitos:

Testing package

Testes unitários: func TestXxx(*testing.T)

Table-driven tests

Mocks para clientes K8s

Testes com controladores (envtest)

Analogia: Testes em GO são simples e built-in, sem frameworks complexos.

Relação: Testes garantem que suas automações não quebrem com atualizações do K8s.

Aplicação prática:

go
func TestGetPods(t *testing.T) {
    mock := &MockClient{}
    mock.On("GetPods").Return([]Pod{{Name: "test"}}, nil)
    
    result := ProcessPods(mock)
    if result != 1 {
        t.Errorf("Esperado 1, obtido %d", result)
    }
}
Erro comum: Testar implementação ao invés de comportamento.

Exercício:
Escreva testes para uma função que valida ConfigMap antes de aplicar.

Para aprofundar: Testes em GO

### 

## That's all

...Folks!!!


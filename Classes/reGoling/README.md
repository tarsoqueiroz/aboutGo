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

# 2. Inicie o módulo. O nome é importante: use um caminho que reflita onde ele ficará (ex: github.com/tarsoqueiroz/hellok8s)
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

### A Estrutura de Pacotes: A Base da Organização

Em Go, a organização é simples e direta: **cada diretório é um pacote**. O nome do pacote é definido pela primeira linha do arquivo, com `package nome`.

**Analogia Prática**: Imagine um escritório de arquitetura:

- **Pacotes** = Gavetas do arquivo (cada uma com uma função)
- **Funções exportadas (Maiúsculas)** = Documentos que podem ser compartilhados com outros escritórios
- **Funções não exportadas (Minúsculas)** = Anotações internas que só seu escritório vê

**A Regra de Ouro**:

- ✅ `func GetPodName()` → Exportada (visível para outros pacotes)
- ❌ `func parseYaml()` → Não exportada (visível apenas dentro do mesmo pacote)
- ✅ `type PodInfo struct` → Exportada
- ❌ `type podCache struct` → Não exportada

### Estrutura de Diretórios na Prática

Vamos criar a estrutura que você verá em 90% dos projetos Go profissionais:

```sh
# Crie um novo projeto para o dia 2
mkdir projetodia2 && cd projetodia2
go mod init github.com/tarsoqueiroz/projetodia2

# Crie a estrutura de pastas
mkdir k8s
mkdir utils

# Crie os arquivos
touch main.go
touch k8s/helpers.go
touch utils/strings.go
```

Sua estrutura agora é:

```text
projetodia2/
├── go.mod
├── main.go
├── k8s/
│   └── helpers.go
└── utils/
    └── strings.go
```

### Implementando os Pacotes

Arquivo `k8s/helpers.go`:

```go
package k8s

import "fmt"

// GetPodName é uma função EXPORTADA (maiúscula)
// Pode ser usada por outros pacotes
func GetPodName(namespace, name string) string {
    return fmt.Sprintf("%s/%s", namespace, name)
}

// parseYaml é uma função NÃO EXPORTADA (minúscula)
// Só pode ser usada dentro do pacote k8s
func parseYaml(yamlData string) string {
    // Simulação: num cenário real, usaria a biblioteca yaml
    return "Parsed: " + yamlData
}

// PodInfo é uma struct EXPORTADA
type PodInfo struct {
    Name      string
    Namespace string
    Ready     bool
}

// podCache é uma struct NÃO EXPORTADA (uso interno)
type podCache struct {
    pods []PodInfo
}
```

Arquivo `utils/strings.go`:

```go
package utils

import "strings"

// UpperExportada - função exportada
func UpperExportada(texto string) string {
    return strings.ToUpper(texto)
}

// lowerInterna - função não exportada
func lowerInterna(texto string) string {
    return strings.ToLower(texto)
}
```

Arquivo `main.go`:

```go
package main

import (
    "fmt"
    
    // Importando nossos pacotes locais
    "github.com/tarsoqueiroz/projetodia2/k8s"
    "github.com/tarsoqueiroz/projetodia2/utils"
)

func main() {
    // Usando função exportada do pacote k8s
    podName := k8s.GetPodName("default", "nginx-pod")
    fmt.Println("Nome do Pod:", podName)
    
    // Usando função exportada do pacote utils
    upper := utils.UpperExportada("hello kubernetes")
    fmt.Println("Upper:", upper)
    
    // ❌ Isso NÃO funciona - parseYaml não é exportada
    // k8s.parseYaml("teste")  // Erro de compilação!
    
    // Usando struct exportada
    pod := k8s.PodInfo{
        Name:      "nginx-pod",
        Namespace: "default",
        Ready:     true,
    }
    fmt.Printf("Pod: %+v\n", pod)
    
    // ✅ CORRETO: usar o alias para simplificar
    // k8s é o nome do pacote, não precisa de alias a menos que haja conflito
}
```

### Alias de Importação: Quando Usar

Às vezes você precisa dar um apelido para o pacote:

```go
import (
    "fmt"
    
    // Alias "k8sclient" para o pacote
    k8sclient "github.com/tarsoqueiroz/projetodia2/k8s"
    
    // Alias "_" para ignorar (útil para init())
    _ "github.com/lib/pq"  // Inicializa o driver PostgreSQL
)

func main() {
    // Agora usa o alias
    nome := k8sclient.GetPodName("default", "pod")
}
```

### O Pacote `main` é Especial

- `package main` define um programa executável
- Deve ter a função `func main()`
- Não pode ser importado por outros pacotes
- É o ponto de entrada do seu binário

### Erros e Confusões Comuns

| Erro | Sintoma | Solução |
| :--- | :------ | :------ |
| Exportação esquecida | `undefined: k8s.parseYaml` | Mude para `ParseYaml` (maiúscula) |
| Pacote com nome diferente da pasta | `import "projeto/k8s"` mas `package kubernetes` | Nome do pacote DEVE ser `k8s` |
| Importação com caminho errado | `cannot find package` | Use o caminho completo do módulo + pasta |
| Ciclo de importação | `import cycle not allowed` | Pacote A importa B, e B importa A - reorganize |
| Função main em pacote não-main | `main redeclared` | Só pode ter `func main()` no pacote `main` |

### Exercício para Fixar

**Objetivo**: Criar um sistema simples de gerenciamento de "Pods" com pacotes bem organizados.

**Instruções**:

- Estrutura:

```text
gerenciador-pods/
├── go.mod
├── main.go
├── pods/
│   ├── manager.go     // Gerencia operações com pods
│   └── types.go       // Define structs de pod
└── utils/
    └── logger.go      // Funções de log
```

- Requisitos:

Em `pods/types.go`: Crie struct `Pod` exportada com campos: `Name`, `Namespace`, `Status` (string)

Em pods/manager.go:

- Função exportada `ListPods()` que retorna um slice de `Pod` (simule com 2-3 pods)
- Função não exportada `validatePod(p Pod) bool` (verifica se Name não está vazio)

Em `utils/logger.go`:

- Função exportada `Info(msg string)` que imprime `[INFO] msg`
- Função não exportada `logWithLevel(level, msg string)`

Em `main.go`:

- Importe os pacotes
- Liste os pods usando pods.ListPods()
- Log cada pod usando utils.Info()

**Execute e veja funcionar**.

### Solução

```sh
# criando estrutura de diretórios
mkdir -p gerenciador-pods/{pods,utils}
cd gerenciador-pods/

# gerando 
touch main.go pods/manager.go pods/types.go utils/logger.go
```

- `pods/types.go`:

```go
package pods

// Pod - exportada
type Pod struct {
    Name      string
    Namespace string
    Status    string
}
```

- `pods/manager.go`:

```go
package pods

// ListPods - exportada
func ListPods() []Pod {
    // Simulando dados
    return []Pod{
        {Name: "nginx", Namespace: "default", Status: "Running"},
        {Name: "redis", Namespace: "cache", Status: "Pending"},
        {Name: "api", Namespace: "production", Status: "Running"},
    }
}

// validatePod - NÃO exportada
func validatePod(p Pod) bool {
    return p.Name != ""
}
```

- `utils/logger.go`:

```go
package utils

import "fmt"

// Info - exportada
func Info(msg string) {
    logWithLevel("INFO", msg)
}

// logWithLevel - NÃO exportada
func logWithLevel(level, msg string) {
    fmt.Printf("[%s] %s\n", level, msg)
}
```

- `main.go`:

```go
package main

import (
    "fmt"
    
    "github.com/tarsoqueiroz/gerenciador-pods/pods"
    "github.com/tarsoqueiroz/gerenciador-pods/utils"
)

func main() {
    utils.Info("Iniciando gerenciador de pods")
    
    podsList := pods.ListPods()
    utils.Info(fmt.Sprintf("Encontrados %d pods", len(podsList)))
    
    for _, pod := range podsList {
        utils.Info(fmt.Sprintf("Pod: %s/%s - %s", 
            pod.Namespace, pod.Name, pod.Status))
    }
    
    // ❌ Isto não compila (validatePod não é exportada)
    // pods.validatePod(pods.Pod{Name: "teste"})
}
```

```sh
# inicializando o projeto
go mod init  github.com/tarsoqueiroz/gerenciador-pods

# executando o programa
go run main.go 
```

### O que Estudar para Aprofundar

- **Pacotes internos (`internal/`)**: [Go Internal Packages](https://go.dev/doc/go1.4#internalpackages) - como criar pacotes que só podem ser usados dentro do seu módulo.
- **Inicialização (`init()`)**: [Package initialization](https://go.dev/doc/effective_go#init) - funções `init()` rodam antes do `main()`.
- **Pacotes com múltiplos arquivos**: Todos os arquivos no mesmo diretório devem ter o mesmo `package nome`.
- **Blank identifier em imports**: `_ "pacote"` para importar apenas para executar init().

### Checklist de Conclusão do Dia 2

- □ Entendi que **exportado = Maiúsculo, não exportado = Minúsculo**
- □ Criei pacotes em subdiretórios
- □ Importei e usei meus próprios pacotes
- □ Entendi a diferença entre `package main` e outros pacotes
- □ Completei o exercício do gerenciador de pods
- □ Testei o que acontece quando tento usar uma função não exportada

### Dica para seu Contexto (Kubernetes)

Em projetos reais de Kubernetes com Go, você verá esta estrutura frequentemente:

```go
// client-go usa este padrão
import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"  // Alias comum
    corev1 "k8s.io/api/core/v1"                    // Alias para versão
)
```

O padrão é: `<pacote>v<versão>` para clientes de API, permitindo várias versões simultâneas.

## DIA 3: Tipos e Structs

### Proposta

**Conceitos**:

- Tipos básicos: string, int, bool, float64
- Structs = objetos (mas sem herança)
- Métodos em structs
- Ponteiros (* e &) - quando usar

**Analogia**: Struct é como uma classe em Java, mas só dados. Métodos são funções anexadas.

**Relação**: Structs são usados para modelar recursos Kubernetes (Pod, Service, etc).

**Aplicação prática**:

```go
type Pod struct {
    Name      string
    Namespace string
    Ready     bool
}

func (p Pod) GetFullName() string {
    return p.Namespace + "/" + p.Name
}
```

**Erro comum**: Esquecer que GO passa tudo por valor (cópia) - use ponteiros para modificar.

**Exercício**:

Crie struct Deployment com campos Name, Replicas, Image. Adicione método Scale() que altera replicas.

**Para aprofundar**: Métodos e ponteiros

### O Que São Structs? A Base dos Dados (10 min)

Analogia: Structs são como "formulários" ou "fichas cadastrais" - um molde que define quais campos um dado deve ter. Pense em um formulário de cadastro de Pod:



















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


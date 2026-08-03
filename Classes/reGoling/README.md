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

### O Que São Structs? A Base dos Dados

**Analogia**: Structs são como "formulários" ou "fichas cadastrais" - um molde que define quais campos um dado deve ter. Pense em um formulário de cadastro de Pod:

```text
Formulário do Pod:
- Nome: _____________ (string)
- Namespace: ________ (string) 
- Réplicas: _________ (int)
- Ready: ____________ (bool)
```

Em Go, isso se torna:

```go
type Pod struct {
    Name      string
    Namespace string
    Replicas  int
    Ready     bool
}
```

**Comparação com outras linguagens**:

- **Java/C#**: Struct é como uma classe POJO, mas sem herança
- **Python**: Similar a uma dataclass, mas com métodos separados
- **JavaScript**: Como um objeto com tipagem estática

> **Importante**: Structs são tipos de valor (passados por cópia), não referência como classes em Java.

### Criando e Usando Structs

```go
package main

import "fmt"

// Definindo uma struct
type Pod struct {
	Name      string
	Namespace string
	Replicas  int
	Ready     bool
}

// Struct aninhada (comum em Kubernetes)
type Deployment struct {
	Name   string
	Pod    Pod // Composição (não herança)
	Labels map[string]string
}

func main() {
	// Forma 1: Declarativa (recomendada)
	pod1 := Pod{
		Name:      "nginx",
		Namespace: "default", // Vírgula obrigatória mesmo na última linha
		Replicas:  1,
		Ready:     true,
	}

	// Forma 2: Posicional (NÃO RECOMENDADO - frágil)
	pod2 := Pod{"redis", "cache", 3, false}

	// Forma 3: Vazia (campos com valores zero)
	pod3 := Pod{}

	// Forma 4: Parcial (campos não especificados ficam com valor zero)
	pod4 := Pod{Name: "api", Namespace: "prod"}

	fmt.Println(pod1)
	fmt.Printf("Pod 1: %+v\n", pod1) // %+v mostra nomes dos campos
	fmt.Printf("Pod 2: %+v\n", pod2) // %+v mostra nomes dos campos
	fmt.Printf("Pod 3: %+v\n", pod3) // %+v mostra nomes dos campos
	fmt.Printf("Pod 4: %+v\n", pod4) // %+v mostra nomes dos campos
}
```

### Métodos: Comportamento Anexado à Struct

**Analogia**: Se a struct é um "formulário", os métodos são as "ações" que você pode fazer com esse formulário.

```go
package main

import "fmt"

type Pod struct {
	Name      string
	Namespace string
	Replicas  int
	Ready     bool
}

// Método com RECEPTOR POR VALOR (cópia) - NÃO MODIFICA o original
func (p Pod) GetFullName() string {
	return p.Namespace + "/" + p.Name
}

// Método com RECEPTOR POR PONTEIRO (*) - PODE MODIFICAR o original
func (p *Pod) SetReplicas(n int) {
	p.Replicas = n // Modifica o campo
}

// Método com RECEPTOR POR PONTEIRO - útil para validação
func (p *Pod) IsValid() bool {
	return p.Name != "" && p.Namespace != ""
}

// Método que retorna uma cópia modificada
func (p Pod) WithReadyStatus(ready bool) Pod {
	p.Ready = ready
	return p // Retorna uma nova struct
}

func main() {
	pod := Pod{Name: "nginx", Namespace: "default", Replicas: 1, Ready: true}

	fmt.Printf("Pod: %+v\n", pod) // %+v mostra nomes dos campos

	// Método com receptor por valor (cópia)
	fullName := pod.GetFullName()
	fmt.Println("Nome completo:", fullName)

	// Método com receptor por ponteiro (modifica original)
	pod.SetReplicas(3)
	fmt.Printf("Pod após SetReplicas: %+v\n", pod)

	// Método que retorna nova cópia
	newPod := pod.WithReadyStatus(false)
	fmt.Printf("Novo Pod: %+v\n", newPod)
	fmt.Printf("Pod original ainda é: %+v\n", pod) // Não mudou
}
```

### Regra de Ouro: Quando Usar Ponteiro vs Valor?

| Situação | Use  | Motivo |
| :------- | :--- | :----- |
| Precisa MODIFICAR a struct               | `*Pod`          | Ponteiro permite alterar o original |
| A struct é GRANDE (ex: > 64 bytes)       | `*Pod`          | Evita cópia cara |
| Precisa garantir que a struct é IMUTÁVEL | `Pod`           | Cópia segura |
| Método que apenas LÊ dados               | `Pod` ou `*Pod` | Ambos funcionam, mas valor é mais seguro |
| A struct contém um MUTEX ou similar      | `*Pod`          | Mutex deve ser passado por ponteiro |

> **Dica Prática**: Em dúvida, use ponteiro (`*Pod`). É mais seguro e performático.

### Campos com Tags: Metadados para Serialização

Tags são anotações que dizem como a struct deve ser serializada/deserializada. **Essencial para Kubernetes**:

```go
package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type ConfigMap struct {
	Name      string            `json:"name" yaml:"name"`
	Namespace string            `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Data      map[string]string `json:"data" yaml:"data"`
	Internal  string            `json:"-" yaml:"-"` // Ignorado em ambos
}

func main() {
	cm := ConfigMap{
		Name:      "app-config",
		Namespace: "default",
		Data:      map[string]string{"key": "value", "timeout": "30s"},
		Internal:  "segredo",
	}

	// Serializando para JSON
	jsonData, _ := json.MarshalIndent(cm, "", "  ")
	fmt.Println("=== JSON ===")
	fmt.Println(string(jsonData))

	// Serializando para YAML
	yamlData, _ := yaml.Marshal(cm)
	fmt.Println("\n=== YAML ===")
	fmt.Println(string(yamlData))
}
```

### Diferenças Importantes entre JSON e YAML

| Característica | JSON | YAML |
| :------------- | :--- | :--- |
| Uso em K8s     | API Server (internamente)  | Manifestos (arquivos .yaml) |
| Legibilidade   | Boa, mas com muitas chaves | Excelente, mais limpo |
| Comentários    | Não suporta                | Suporta (# comentário) |
| Estrutura      | Chaves e colchetes         | Indentação (como Python) |
| Tags           | `json:"nome"`              | `yaml:"nome"` |

### Caso Real: Lendo e Escrevendo ConfigMap do Kubernetes

```go
package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type ConfigMap struct {
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace"`
	Data      map[string]string `yaml:"data"`
}

func main() {
	// Simulando um ConfigMap que você leu de um arquivo
	yamlString := `
name: nginx-config
namespace: production
data:
  nginx.conf: |
    server {
        listen 80;
        server_name localhost;
    }
  timeout: 60s
`

	var cm ConfigMap
	err := yaml.Unmarshal([]byte(yamlString), &cm)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n=== ConfigMap lido ===")
	fmt.Printf("ConfigMap lido: %+v\n", cm)

	// Modificando e salvando de volta
	cm.Data["new-key"] = "new-value"

	novoYAML, _ := yaml.Marshal(cm)
	fmt.Println("\n=== ConfigMap modificado ===")
	fmt.Println("\nConfigMap modificado:")
	fmt.Println(string(novoYAML))
}
```

### Dica para seu Contexto:

No dia a dia com Kubernetes, você vai usar **YAML para arquivos de manifesto** e **JSON para comunicação com a API** (via client-go). Por isso, é comum ver structs com **ambas as tags**:

```go
type Pod struct {
    Name      string `json:"name" yaml:"name"`
    Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
    // Em K8s real, você veria muitas outras tags como `protobuf:"..."`
}
```

### Erros e Confusões Comuns

| Erro | Sintoma | Solução |
| :--- | :------ | :------ |
| Esquecer vírgula na última linha              | `missing ',' before newline`               | TODAS as linhas de campos têm vírgula |
| Usar valor sem ponteiro e esperar modificação | `p.Replicas = 3` não altera o original     | Use receptor com ponteiro: `(p *Pod)` |
| Esquecer `&` ao criar ponteiro                | `&Pod{Name: "test"}` para criar referência | Use `&` ou `new(Pod)` |
| Campo com nome maiúsculo não exportado?       | Campo começa com maiúscula = exportado     | Minúsculo para interno |
| Comparar structs com campos não comparáveis   | `map` ou `slice` causam erro               | Use `DeepEqual` do pacote `reflect` |

### Exercício para Fixar

**Objetivo**: Modelar um recurso Kubernetes Deployment com métodos para gerenciá-lo.

**Instruções**:

- Crie um novo módulo `k8s-model`:

```bash
mkdir k8s-model && cd k8s-model
go mod init github.com/seu-usuario/k8s-model
```

- Crie o arquivo `main.go` com:
  - **Parte 1 - Definir Structs**:
    - `Pod`: campos Name (string), Image (string), Ports ([]int)
    - `Deployment`: campos Name (string), Replicas (int), PodTemplate (Pod)
    - `DeploymentStatus`: campos Ready (bool), AvailableReplicas (int)
    - `Annotations`: campos (map[string]string)
  - **Parte 2 - Métodos**:
    - `(d *Deployment) Scale(n int)`: atualiza Replicas
    - `(d Deployment) GetFullName() string`: retorna "deployment/name"
    - `(d *Deployment) Validate() bool`: verifica se Name != "" e Replicas > 0
    - `(d *Deployment) Status() DeploymentStatus`: retorna status (Ready = Replicas > 0)
  - **Parte 3 - Teste**:
    - Crie um deployment com 3 réplicas
    - Escale para 5 réplicas
    - Valide e mostre o status
    - Serializar o deployment em YAML e JSON
    - Imprimir ambos os formatos
- Execute e veja funcionar.

### Solução

```go
package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

// Pod representa um pod Kubernetes
type Pod struct {
	Name  string `json:"name" yaml:"name"`
	Image string `json:"image" yaml:"image"`
	Ports []int  `json:"ports,omitempty" yaml:"ports,omitempty"`
}

// Deployment representa um deployment Kubernetes
type Deployment struct {
	Name        string            `json:"name" yaml:"name"`
	Replicas    int               `json:"replicas" yaml:"replicas"`
	PodTemplate Pod               `json:"podTemplate" yaml:"podTemplate"`
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
}

// DeploymentStatus representa o status de um deployment
type DeploymentStatus struct {
	Ready             bool `json:"ready" yaml:"ready"`
	AvailableReplicas int  `json:"availableReplicas" yaml:"availableReplicas"`
}

// Scale - modifica o número de réplicas (PONTEIRO)
func (d *Deployment) Scale(n int) {
	if n > 0 {
		d.Replicas = n
	}
}

// GetFullName - retorna nome completo (VALOR - apenas leitura)
func (d Deployment) GetFullName() string {
	return "deployment/" + d.Name
}

// Validate - verifica se o deployment é válido (PONTEIRO)
func (d *Deployment) Validate() bool {
	return d.Name != "" && d.Replicas > 0
}

// Status - retorna o status atual (VALOR - cria cópia)
func (d Deployment) Status() DeploymentStatus {
	return DeploymentStatus{
		Ready:             d.Replicas > 0,
		AvailableReplicas: d.Replicas,
	}
}

// String - implementa a interface fmt.Stringer (método especial)
func (d Deployment) String() string {
	return fmt.Sprintf("Deployment: %s (replicas: %d)", d.Name, d.Replicas)
}

func main() {
	// Criando um deployment
	deploy := Deployment{
		Name:     "nginx-deployment",
		Replicas: 3,
		PodTemplate: Pod{
			Name:  "nginx-pod",
			Image: "nginx:1.21",
			Ports: []int{80},
		},
		Annotations: map[string]string{
			"description": "Test deployment for Day 3",
			"owner":       "arquiteto",
			"environment": "dev",
		},
	}

	fmt.Println("=== INFORMAÇÕES DO DEPLOYMENT ===")
	fmt.Println("Criado:", deploy.String())
	fmt.Println("Nome completo:", deploy.GetFullName())
	fmt.Println("Válido?", deploy.Validate())

	// Status inicial
	status := deploy.Status()
	fmt.Printf("Status inicial: %+v\n", status)

	// Escalando
	deploy.Scale(5)
	fmt.Println("Após scale:", deploy.String())

	// Novo status
	newStatus := deploy.Status()
	fmt.Printf("Novo status: %+v\n", newStatus)

	// Teste com deployment inválido
	invalid := Deployment{Name: ""}
	fmt.Println("Deployment inválido é válido?", invalid.Validate())

	// ============================================
	// SERIALIZAÇÃO EM JSON E YAML (conforme solicitado)
	// ============================================
	fmt.Println("\n=== SERIALIZAÇÃO ===")

	// Serializando para JSON
	jsonData, err := json.MarshalIndent(deploy, "", "  ")
	if err != nil {
		fmt.Println("Erro ao serializar para JSON:", err)
	} else {
		fmt.Println("JSON:")
		fmt.Println(string(jsonData))
	}

	// Serializando para YAML
	yamlData, err := yaml.Marshal(deploy)
	if err != nil {
		fmt.Println("Erro ao serializar para YAML:", err)
	} else {
		fmt.Println("\nYAML:")
		fmt.Println(string(yamlData))
	}

	// ============================================
	// DEMONSTRAÇÃO DE DESERIALIZAÇÃO (bônus)
	// ============================================
	fmt.Println("\n=== DESERIALIZAÇÃO A PARTIR DO YAML ===")
	yamlString := `
name: redis-deployment
replicas: 2
podTemplate:
  name: redis-pod
  image: redis:alpine
  ports:
  - 6379
annotations:
  description: Redis cache
  team: database
`
	var newDeploy Deployment
	err = yaml.Unmarshal([]byte(yamlString), &newDeploy)
	if err != nil {
		fmt.Println("Erro ao desserializar YAML:", err)
	} else {
		fmt.Println("Deployment criado a partir do YAML:")
		fmt.Printf("  Nome: %s\n", newDeploy.Name)
		fmt.Printf("  Réplicas: %d\n", newDeploy.Replicas)
		fmt.Printf("  Pod: %s (%s)\n", newDeploy.PodTemplate.Name, newDeploy.PodTemplate.Image)
		fmt.Printf("  Portas: %v\n", newDeploy.PodTemplate.Ports)
		fmt.Printf("  Anotações: %v\n", newDeploy.Annotations)
	}
}
```

### O que Estudar para Aprofundar

- **Métodos com ponteiro vs valor em profundidade**: [Go Tour - Methods](https://go.dev/tour/methods/1) - seção completa
- **Embedding de structs (composição)**: [Composition with Structs](https://go.dev/doc/effective_go#embedding) - como "herdar" campos
- **Tags customizadas**: [Struct Tags](https://go.dev/ref/spec#Tag) - para criar suas próprias anotações
- **Comparação de structs**: `reflect.DeepEqual` e `cmp.Diff` para testes
- **Métodos com receptores `nil`**: Quando o ponteiro pode ser `nil` e como tratar

### Contexto Kubernetes: Como Você Vai Usar Isso

No mundo real com client-go, você verá structs como estas:

```go
import (
    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Pod do Kubernetes (versão simplificada da real)
type Pod struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata"`
    Spec    PodSpec   `json:"spec"`
    Status  PodStatus `json:"status"`
}
```

**Padrão**:

- `ObjectMeta` contém Name, Namespace, Labels, etc.
- `Spec` é o "desejo" (o que você quer)
- `Status` é a "realidade" (o que está acontecendo)
- Métodos usam ponteiros para modificar o Spec

### Checklist de Conclusão do Dia 3

- □ Entendi que struct = molde para dados (como formulário)
- □ Sei criar structs de 4 formas diferentes
- □ Compreendo a diferença entre receptor por valor (`p Pod`) e ponteiro (`p *Pod`)
- □ Sei quando usar ponteiro vs valor para métodos
- □ Usei tags `json` e `yaml` para serialização
- □ Completei o exercício do Deployment
- □ Compreendo que em K8s, `Spec` = desejo, `Status` = realidade

## DIA 4: Interfaces

### Proposta

**Conceitos**:

- Interface = contrato de métodos
- Satisfação implícita (não precisa declarar "implements")
- Interface vazia (interface{}) = any

**Analogia**: Interface é como uma API que você precisa implementar, mas sem declaração explícita.

**Relação**: Interfaces permitem desacoplar código - essencial para testabilidade.

**Aplicação prática**:

```go
type K8sClient interface {
    GetPods(namespace string) ([]Pod, error)
    DeletePod(name string) error
}
```

**Erro comum**: Pensar que precisa declarar que implementa uma interface.

**Exercício**:

- Crie interface Logger com método Log(`message string`). Implemente com FileLogger e ConsoleLogger.

**Para aprofundar**: Interfaces no GO

### O Que São Interfaces? O Contrato

**Analogia**: Pense em uma interface como um **contrato de serviço** ou um **plugue universal**:

- Uma tomada (interface) define que qualquer dispositivo com plugue padrão pode ser conectado
- Não importa se é um carregador, liquidificador ou computador - todos seguem o mesmo contrato

Em Go, uma interface define **o que** um tipo deve fazer, não como ele faz.

```go
// Interface = Contrato
type Tomada interface {
    Ligar() string
    Desligar() string
}

// Diferentes tipos implementam o mesmo contrato
type Liquidificador struct{}
func (l Liquidificador) Ligar() string   { return "Liquidificador ligado!" }
func (l Liquidificador) Desligar() string { return "Liquidificador desligado!" }

type Computador struct{}
func (c Computador) Ligar() string   { return "Computador bootando..." }
func (c Computador) Desligar() string { return "Computador desligando..." }

// Ambos podem ser usados na tomada!
```

### Satisfação Implícita: O Diferencial do Go

**A mágica do Go**: Um tipo **implementa automaticamente** uma interface se tiver todos os métodos necessários. **Não precisa declarar** `implements` como em Java/C#.

```go
package main

import "fmt"

// Interface
type Saudacao interface {
    Ola() string
}

// Tipo 1 - implementa automaticamente
type Portugues struct{}
func (p Portugues) Ola() string { return "Olá!" }

// Tipo 2 - também implementa
type Ingles struct{}
func (i Ingles) Ola() string { return "Hello!" }

// Função que aceita QUALQUER tipo que implemente Saudacao
func Cumprimentar(s Saudacao) {
    fmt.Println(s.Ola())
}

func main() {
    // Nenhuma declaração "implements" necessária!
    p := Portugues{}
    i := Ingles{}
    
    Cumprimentar(p) // "Olá!"
    Cumprimentar(i) // "Hello!"
}
```

**Por que isso é poderoso**?

- Você pode adicionar interfaces **depois** que os tipos já foram criados
- Bibliotecas externas podem implementar suas interfaces sem saber delas
- Facilita testes: crie mocks facilmente

### A Interface Vazia: `interface{}` = `any`

```go
package main

// Interface vazia - aceita QUALQUER tipo
var qualquer interface{}
qualquer = 42
qualquer = "string"
qualquer = struct{Nome string}{"João"}

// Em Go 1.18+, pode usar 'any' (alias para interface{})
var qualquer2 any
qualquer2 = true

// Útil para funções genéricas (mas prefira tipos específicos quando possível)
func Imprimir(v any) {
    fmt.Printf("Valor: %v, Tipo: %T\n", v, v)
}

func main() {
    Imprimir(42)           // Valor: 42, Tipo: int
    Imprimir("teste")      // Valor: teste, Tipo: string
    Imprimir([]int{1,2,3}) // Valor: [1 2 3], Tipo: []int
}
```

> **⚠️ Cuidado**: Use any com moderação! Perde a segurança de tipos.

### Type Assertion e Type Switch: Trabalhando com Interfaces

```go
func Processar(v any) {
    // Type assertion - verifica se é um tipo específico
    if str, ok := v.(string); ok {
        fmt.Println("É uma string:", str)
        return
    }
    
    // Type switch - verifica múltiplos tipos
    switch t := v.(type) {
    case int:
        fmt.Println("Inteiro:", t*2)
    case string:
        fmt.Println("String em maiúsculas:", strings.ToUpper(t))
    case bool:
        fmt.Println("Booleano:", !t)
    default:
        fmt.Println("Tipo desconhecido:", t)
    }
}
```

### Aplicação Prática: Interface para Cliente Kubernetes (10 min)

```go
package main

import "fmt"

// Pod - struct simples para demonstração
type Pod struct {
	Name      string
	Namespace string
	Ready     bool
}

// Interface do cliente Kubernetes
type K8sClient interface {
	GetPods(namespace string) ([]Pod, error)
	DeletePod(name string, namespace string) error
	GetPodLogs(name string, namespace string) (string, error)
}

// Implementação REAL (simulada)
type RealK8sClient struct {
	serverURL string
}

func (c RealK8sClient) GetPods(namespace string) ([]Pod, error) {
	// Simulação: em produção, chamaria a API do K8s
	fmt.Printf("🔴 Conectando a %s para buscar pods\n", c.serverURL)
	return []Pod{
		{Name: "nginx", Namespace: namespace, Ready: true},
		{Name: "redis", Namespace: namespace, Ready: false},
	}, nil
}

func (c RealK8sClient) DeletePod(name string, namespace string) error {
	fmt.Printf("🔴 Deletando pod %s/%s\n", namespace, name)
	return nil
}

func (c RealK8sClient) GetPodLogs(name string, namespace string) (string, error) {
	fmt.Printf("🔴 Buscando logs de %s/%s\n", namespace, name)
	return "Logs simulados...", nil
}

// Implementação MOCK (para testes)
type MockK8sClient struct {
	ShouldFail bool
}

func (m MockK8sClient) GetPods(namespace string) ([]Pod, error) {
	if m.ShouldFail {
		return nil, fmt.Errorf("erro simulado")
	}
	return []Pod{
		{Name: "mock-pod-1", Namespace: namespace, Ready: true},
		{Name: "mock-pod-2", Namespace: namespace, Ready: true},
	}, nil
}

func (m MockK8sClient) DeletePod(name string, namespace string) error {
	if m.ShouldFail {
		return fmt.Errorf("erro ao deletar mock")
	}
	fmt.Printf("✅ Mock: Pod %s/%s deletado\n", namespace, name)
	return nil
}

func (m MockK8sClient) GetPodLogs(name string, namespace string) (string, error) {
	if m.ShouldFail {
		return "", fmt.Errorf("erro ao buscar logs mock")
	}
	return "✅ Logs mockados...", nil
}

// Função que usa a interface (desacoplada da implementação)
func ProcessarPods(cliente K8sClient, namespace string) {
	pods, err := cliente.GetPods(namespace)
	if err != nil {
		fmt.Printf("❌ Erro ao buscar pods: %v\n", err)
		return
	}

	fmt.Printf("📦 Encontrados %d pods no namespace %s\n", len(pods), namespace)
	for _, pod := range pods {
		status := "✅"
		if !pod.Ready {
			status = "❌"
		}
		fmt.Printf("  %s Pod: %s (Ready: %v)\n", status, pod.Name, pod.Ready)
	}
}

func main() {
	// Usando o cliente REAL
	realClient := RealK8sClient{serverURL: "https://k8s-api.cluster.local"}
	fmt.Println("=== Teste com Cliente REAL ===")
	ProcessarPods(realClient, "default")

	// Usando o cliente MOCK (para testes)
	mockClient := MockK8sClient{ShouldFail: false}
	fmt.Println("\n=== Teste com Cliente MOCK ===")
	ProcessarPods(mockClient, "test")

	// Simulando falha no mock
	mockClientFail := MockK8sClient{ShouldFail: true}
	fmt.Println("\n=== Teste com MOCK que FALHA ===")
	ProcessarPods(mockClientFail, "test")
}
```

### Erros e Confusões Comuns

| Erro | Sintoma | Solução |
| :--- | :------ | :------ |
| Achar que precisa declarar `implements` | Procura sintaxe como Java      | Não precisa! Apenas implemente os métodos |
| Esquecer um método da interface         | `cannot use type as K8sClient` | Implemente TODOS os métodos da interface |
| Método com assinatura diferente         | `cannot use type as K8sClient` | Verifique parâmetros e retornos exatos |
| Usar `any` demais                       | Perda de segurança de tipos    | Use tipos específicos sempre que possível |
| Type assertion sem verificação          | `panic: interface conversion`  | Sempre use `v, ok := x.(T)` |

### Exercício para Fixar

**Objetivo**: Criar um sistema de logging com múltiplas implementações usando interfaces.

**Instruções**:

- Crie a interface `Logger`:

```go
type Logger interface {
    Log(message string)
    LogWithLevel(level string, message string)
}
```

- Implemente `ConsoleLogger`:
  - `Log(message)`: imprime no console com prefixo `[INFO]`
  - `LogWithLevel(level, message)`: imprime `[LEVEL] message`
- Implemente `FileLogger`:
  - Armazena mensagens em um slice interno (simulando arquivo)
  - `Log(message)`: adiciona ao slice com prefixo `[INFO]`
  - `LogWithLevel(level, message)`: adiciona ao slice com prefixo `[LEVEL]`
  - Adicione método `GetMessages() []string` (NÃO está na interface)
- Implemente MultiLogger (desafio):
  - Aceita múltiplos loggers
  - Implementa `Logger` chamando todos os loggers internos
- Função de teste:
  - Crie função `ProcessarOperacao(logger Logger, operacao string)`
  - Registre início, progresso e fim da operação
  - Use diferentes loggers

### Solução

```go
package main

import "fmt"

// ==================== INTERFACE ====================
type Logger interface {
	Log(message string)
	LogWithLevel(level string, message string)
}

// ==================== CONSOLE LOGGER ====================
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	c.LogWithLevel("INFO", message)
}

func (c ConsoleLogger) LogWithLevel(level string, message string) {
	fmt.Printf("[%s] %s\n", level, message)
}

// ==================== FILE LOGGER ====================
type FileLogger struct {
	messages []string
}

func (f *FileLogger) Log(message string) {
	f.LogWithLevel("INFO", message)
}

func (f *FileLogger) LogWithLevel(level string, message string) {
	f.messages = append(f.messages, fmt.Sprintf("[%s] %s", level, message))
}

// Método extra (NÃO está na interface)
func (f *FileLogger) GetMessages() []string {
	return f.messages
}

// ==================== MULTI LOGGER (DESAFIO) ====================
type MultiLogger struct {
	loggers []Logger
}

func (m MultiLogger) Log(message string) {
	for _, logger := range m.loggers {
		logger.Log(message)
	}
}

func (m MultiLogger) LogWithLevel(level string, message string) {
	for _, logger := range m.loggers {
		logger.LogWithLevel(level, message)
	}
}

// ==================== FUNÇÃO QUE USA A INTERFACE ====================
func ProcessarOperacao(logger Logger, operacao string) {
	logger.LogWithLevel("START", fmt.Sprintf("Iniciando operação: %s", operacao))
	logger.Log(fmt.Sprintf("Processando etapa 1 de %s", operacao))
	logger.Log(fmt.Sprintf("Processando etapa 2 de %s", operacao))
	logger.LogWithLevel("END", fmt.Sprintf("Finalizando operação: %s", operacao))
}

// ==================== MAIN ====================
func main() {
	fmt.Println("=== CONSOLE LOGGER ===")
	console := ConsoleLogger{}
	ProcessarOperacao(console, "deploy-nginx")

	fmt.Println("\n=== FILE LOGGER ===")
	file := &FileLogger{}
	ProcessarOperacao(file, "scale-redis")

	// Acessando método extra (não disponível via interface)
	fmt.Println("\nMensagens no arquivo:")
	for i, msg := range file.GetMessages() {
		fmt.Printf("  %d: %s\n", i+1, msg)
	}

	fmt.Println("\n=== MULTI LOGGER ===")
	multi := MultiLogger{
		loggers: []Logger{
			ConsoleLogger{},
			&FileLogger{},
		},
	}
	ProcessarOperacao(multi, "rollback-api")
}
```

### Saída Esperada

```text
=== CONSOLE LOGGER ===
[START] Iniciando operação: deploy-nginx
[INFO] Processando etapa 1 de deploy-nginx
[INFO] Processando etapa 2 de deploy-nginx
[END] Finalizando operação: deploy-nginx

=== FILE LOGGER ===

Mensagens no arquivo:
  1: [START] Iniciando operação: scale-redis
  2: [INFO] Processando etapa 1 de scale-redis
  3: [INFO] Processando etapa 2 de scale-redis
  4: [END] Finalizando operação: scale-redis

=== MULTI LOGGER ===
[START] Iniciando operação: rollback-api
[INFO] Processando etapa 1 de rollback-api
[INFO] Processando etapa 2 de rollback-api
[END] Finalizando operação: rollback-api
```

### O que Estudar para Aprofundar

- **Interfaces e Testes**: [Mocking in Go](https://go.dev/blog/using-go-modules) - como usar interfaces para criar testes unitários
- **Interface `io.Reader` e `io.Writer`**: As interfaces mais importantes da stdlib - usadas em todo lugar
- **`error` é uma interface**: Entenda como funciona o tratamento de erros
- **Composição de interfaces**: `type ReadWriter interface { Reader; Writer }`
- **Contexto com interfaces**: Como `context.Context` é uma interface essencial

### Contexto Kubernetes: Como Você Vai Usar Isso

Em projetos com client-go, você verá este padrão:

```go
// Interface do cliente Kubernetes (simplificada)
type Interface interface {
    CoreV1() CoreV1Interface
    AppsV1() AppsV1Interface
}

// Permite mocks em testes
type mockClient struct {}
func (m mockClient) CoreV1() CoreV1Interface { return &mockCoreV1{} }
```

> **Benefício**: Você pode testar seus operadores/controladores sem precisar de um cluster K8s real!

### Checklist de Conclusão do Dia 4

- □ Entendi que interface = contrato de métodos
- □ Compreendo que implementação é **implícita** (não precisa declarar)
- □ Sei a diferença entre `interface{}` (any) e interfaces específicas
- □ Usei type assertion e type switch com segurança
- □ Criei múltiplas implementações da mesma interface
- □ Completei o exercício do Logger
- □ Entendo como interfaces facilitam testes com mocks

## DIA 5: Goroutines e Channels (Parte 1)

### Proposta

**Conceitos**:

- Goroutine = thread leve
- `go func()` - inicia uma goroutine
- WaitGroups para sincronização

**Analogia**: Goroutine é como uma thread em Java, mas muito mais leve (custa ~2KB).

**Relação**: O coração do modelo de concorrência em GO, usado em todos os operadores K8s.

**Aplicação prática**:

```go
var wg sync.WaitGroup
for _, pod := range pods {
    wg.Add(1)
    go func(p Pod) {
        defer wg.Done()
        processPod(p)
    }(pod)
}
wg.Wait()
```

**Erro comum**: Esquecer de passar parâmetros para a goroutine, causando race conditions.

**Exercício**:

- Crie programa que processa 1000 pods em paralelo usando goroutines (limite a 10 simultâneas).

### O Que São Goroutines? Threads Leves

**Analogia**: Imagine um restaurante com muitos chefs:

- **Thread tradicional (Java)**: Cada chef tem sua própria cozinha completa (pesado, ~1MB de memória)
- **Goroutine (Go)**: Chefs compartilham a mesma cozinha, mas cada um tem sua própria estação de trabalho (leve, ~2KB)

**Características**:

- Criadas com a palavra-chave `go`
- Custam ~2KB de pilha (vs ~1MB para threads Java)
- Podem criar milhares sem problemas
- Gerenciadas pelo runtime do Go (não pelo SO)

```go
package main

import (
    "fmt"
    "time"
)

func dizerOla() {
    fmt.Println("Olá de uma goroutine!")
}

func main() {
    // Inicia uma goroutine
    go dizerOla()
    
    // Goroutine com função anônima
    go func() {
        fmt.Println("Olá de outra goroutine!")
    }()
    
    // Dá tempo para as goroutines executarem
    time.Sleep(100 * time.Millisecond)
    fmt.Println("Programa principal terminou")
}
```

### WaitGroups: Sincronização Básica

**Problema**: O programa principal termina antes das goroutines finalizarem.

**Solução**: sync.WaitGroup para esperar todas terminarem.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func processarItem(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Decrementa o contador quando terminar
    
    fmt.Printf("Iniciando item %d\n", id)
    time.Sleep(time.Second) // Simula trabalho
    fmt.Printf("Finalizando item %d\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 5; i++ {
        wg.Add(1) // Incrementa o contador
        go processarItem(i, &wg)
    }
    
    wg.Wait() // Espera todas as goroutines terminarem
    fmt.Println("Todos os itens processados!")
}
```

**Fluxo do WaitGroup**:

- `wg.Add(1)` → contador = 1
- `go processarItem()` → inicia goroutine
- `defer wg.Done()` → quando terminar, contador--
- `wg.Wait()` → bloqueia até contador = 0

### Padrão: Processamento Paralelo com Limite

**O problema real**: Processar 1000 pods em paralelo poderia sobrecarregar o sistema.

**Solução**: Worker pool com limite de concorrência.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Pod - struct simples
type Pod struct {
	Name      string
	Namespace string
}

// processarPod - simula processamento de um pod
func processarPod(p Pod, workerID int) {
	fmt.Printf("[Worker %d] Processando pod %s/%s\n",
		workerID, p.Namespace, p.Name)
	time.Sleep(100 * time.Millisecond) // Simula trabalho
}

func main() {
	// Criando 1000 pods para processar
	pods := make([]Pod, 1000)
	for i := 0; i < 1000; i++ {
		pods[i] = Pod{
			Name:      fmt.Sprintf("pod-%d", i),
			Namespace: "default",
		}
	}

	var wg sync.WaitGroup
	maxWorkers := 10
	semaphore := make(chan struct{}, maxWorkers) // Controle de concorrência

	start := time.Now()

	for i, pod := range pods {
		wg.Add(1)
		semaphore <- struct{}{} // Ocupa um slot do pool

		go func(p Pod, workerID int) {
			defer wg.Done()
			defer func() { <-semaphore }() // Libera o slot

			processarPod(p, workerID)
		}(pod, i%maxWorkers)
	}

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("\nProcessados %d pods em %v\n", len(pods), elapsed)
	fmt.Printf("Média: %v por pod\n", elapsed/time.Duration(len(pods)))
}
```

### Padrão Worker Pool Mais Robusto

Versão mais profissional usando fila de trabalho:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Pod struct {
	Name      string
	Namespace string
}

// WorkerPool - estrutura para gerenciar workers
type WorkerPool struct {
	numWorkers int
	jobs       chan Pod
	wg         sync.WaitGroup
}

// NewWorkerPool - cria um pool com N workers
func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		jobs:       make(chan Pod, 100), // Buffer para jobs
	}
}

// Start - inicia os workers
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// worker - processa jobs do canal
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for pod := range wp.jobs {
		fmt.Printf("[Worker %d] Processando %s/%s\n",
			id, pod.Namespace, pod.Name)
		time.Sleep(50 * time.Millisecond)
	}
}

// Submit - adiciona um job à fila
func (wp *WorkerPool) Submit(pod Pod) {
	wp.jobs <- pod
}

// Wait - espera todos os jobs terminarem
func (wp *WorkerPool) Wait() {
	close(wp.jobs) // Fecha o canal para workers pararem
	wp.wg.Wait()
}

func main() {
	// Criar 1000 pods
	pods := make([]Pod, 1000)
	for i := 0; i < 1000; i++ {
		pods[i] = Pod{
			Name:      fmt.Sprintf("pod-%d", i),
			Namespace: "default",
		}
	}

	// Criar pool com 10 workers
	pool := NewWorkerPool(10)

	start := time.Now()

	// Iniciar workers
	pool.Start()

	// Enviar jobs
	for _, pod := range pods {
		pool.Submit(pod)
	}

	// Esperar finalizar
	pool.Wait()

	elapsed := time.Since(start)
	fmt.Printf("\n✅ Processados %d pods em %v\n", len(pods), elapsed)
	fmt.Printf("📊 Média: %v por pod\n", elapsed/time.Duration(len(pods)))
}
```

### Erros e Confusões Comuns

| Erro | Sintoma | Solução |
| :--- | :------ | :------ |
| Esquecer `wg.Add(1)` antes da goroutine | `panic: sync: negative WaitGroup counter` | Sempre chame Add antes de iniciar a goroutine |
| Passar variável de loop por referência | Todas as goroutines usam o mesmo valor | Passe como parâmetro: `go func(p Pod) { ... }(pod)` |
| Esquecer `defer wg.Done()` | Deadlock em `wg.Wait()` | Use defer para garantir |
| Sem limite de concorrência | Consome muitos recursos | Use semáforo ou worker pool |
| Não fechar channel | `fatal error: all goroutines are asleep` | Feche channels quando não houver mais dados |

### O Problema do Loop (Muito Comum!)

```go
// ❌ ERRADO: todas as goroutines usam a MESMA variável 'i'
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println(i) // Pode imprimir 5 para todas!
    }()
}

// ✅ CORRETO: cada goroutine recebe uma CÓPIA de 'i'
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(valor int) {
        defer wg.Done()
        fmt.Println(valor)
    }(i)
}
```

### Exercício para Fixar

**Objetivo**: Criar um processador de pods com diferentes cenários.

**Instruções**:

- Crie um módulo `processador-pods` com a estrutura básica
- Implemente um processador de pods que:
	- Processa 1000 pods (simule com struct com Name, Namespace, CPURequest, MemoryRequest)
	- Cada pod leva entre 50-150ms para processar (use `time.Sleep` com random)
	- Use no máximo 10 workers simultâneos
	- Mantenha estatísticas: total processado, tempo total, tempo médio por pod
- Adicione 3 cenários de processamento:
  - **Cenário 1**: Processamento simples (apenas log)
  - **Cenário 2**: Processamento com "falha" aleatória (10% de chance de erro)
  - **Cenário 3**: Processamento com retry em caso de falha
- Mostre estatísticas:
  - Quantos pods processados com sucesso
  - Quantos com falha
  - Tempo total de execução

**Esboço da Solução**:

```go
type Pod struct {
  Name          string
  Namespace     string
  CPURequest    int
  MemoryRequest int
}

type Stats struct {
  Total     int
  Success   int
  Failed    int
  StartTime time.Time
  EndTime   time.Time
}

// Use worker pool com canais
```

### Solução

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Pod - struct representando um pod
type Pod struct {
	Name          string
	Namespace     string
	CPURequest    int
	MemoryRequest int
}

// Stats - estatísticas de processamento
type Stats struct {
	Total     int
	Success   int
	Failed    int
	StartTime time.Time
	EndTime   time.Time
}

// PodProcessor - processador de pods
type PodProcessor struct {
	numWorkers int
	pods       []Pod
	stats      Stats
	mu         sync.Mutex
}

// NewPodProcessor - cria novo processador
func NewPodProcessor(numWorkers int, pods []Pod) *PodProcessor {
	return &PodProcessor{
		numWorkers: numWorkers,
		pods:       pods,
		stats: Stats{
			StartTime: time.Now(),
		},
	}
}

// processPod - processa um único pod com possibilidade de falha
func (pp *PodProcessor) processPod(pod Pod, workerID int) error {
	fmt.Printf("[Worker %d] Processando pod %s/%s (CPU: %d, Mem: %d)\n",
		workerID, pod.Namespace, pod.Name, pod.CPURequest, pod.MemoryRequest)

	// Simula tempo de processamento variável
	processTime := time.Duration(50+rand.Intn(100)) * time.Millisecond
	time.Sleep(processTime)

	// Simula falha aleatória (10% de chance)
	if rand.Float32() < 0.1 {
		return fmt.Errorf("falha ao processar pod %s/%s", pod.Namespace, pod.Name)
	}

	return nil
}

// processPodWithRetry - processa com retry
func (pp *PodProcessor) processPodWithRetry(pod Pod, workerID int) error {
	maxRetries := 3
	for attempt := 0; attempt < maxRetries; attempt++ {
		err := pp.processPod(pod, workerID)
		if err == nil {
			return nil
		}

		if attempt < maxRetries-1 {
			fmt.Printf("[Worker %d] ⚠️  Falha no pod %s (tentativa %d), retentando...\n",
				workerID, pod.Name, attempt+1)
			time.Sleep(100 * time.Millisecond)
		} else {
			return fmt.Errorf("falha após %d tentativas: %v", maxRetries, err)
		}
	}
	return nil
}

// Run - executa o processamento
func (pp *PodProcessor) Run() {
	pp.stats.Total = len(pp.pods)

	var wg sync.WaitGroup
	jobs := make(chan Pod, len(pp.pods))

	// Iniciar workers
	for i := 0; i < pp.numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for pod := range jobs {
				// Processa com retry
				err := pp.processPodWithRetry(pod, workerID)

				pp.mu.Lock()
				if err != nil {
					pp.stats.Failed++
					fmt.Printf("[Worker %d] ❌ Falha no pod %s: %v\n",
						workerID, pod.Name, err)
				} else {
					pp.stats.Success++
				}
				pp.mu.Unlock()
			}
		}(i)
	}

	// Enviar jobs
	for _, pod := range pp.pods {
		jobs <- pod
	}
	close(jobs)

	wg.Wait()
	pp.stats.EndTime = time.Now()
}

// PrintStats - imprime estatísticas
func (pp *PodProcessor) PrintStats() {
	duration := pp.stats.EndTime.Sub(pp.stats.StartTime)
	avgTime := duration / time.Duration(pp.stats.Total)

	fmt.Println("\n=== ESTATÍSTICAS DE PROCESSAMENTO ===")
	fmt.Printf("📊 Total de pods: %d\n", pp.stats.Total)
	fmt.Printf("✅ Sucessos: %d (%.1f%%)\n",
		pp.stats.Success, float64(pp.stats.Success)/float64(pp.stats.Total)*100)
	fmt.Printf("❌ Falhas: %d (%.1f%%)\n",
		pp.stats.Failed, float64(pp.stats.Failed)/float64(pp.stats.Total)*100)
	fmt.Printf("⏱️  Tempo total: %v\n", duration)
	fmt.Printf("📈 Tempo médio por pod: %v\n", avgTime)
	fmt.Printf("🚀 Workers utilizados: %d\n", pp.numWorkers)
}

func main() {
	// Seed do random
	rand.Seed(time.Now().UnixNano())

	// Criar 1000 pods
	pods := make([]Pod, 1000)
	for i := 0; i < 1000; i++ {
		pods[i] = Pod{
			Name:          fmt.Sprintf("pod-%d", i),
			Namespace:     "default",
			CPURequest:    100 + rand.Intn(900),
			MemoryRequest: 256 + rand.Intn(1024),
		}
	}

	fmt.Printf("📦 Criados %d pods para processar\n", len(pods))
	fmt.Println("=== INICIANDO PROCESSAMENTO ===")

	// Processar com 10 workers
	processor := NewPodProcessor(10, pods)
	processor.Run()
	processor.PrintStats()
}
```

### O que Estudar para Aprofundar

- **Race conditions**: Execute com `go run -race` para detectar condições de corrida
- **Mutex vs Channels**: Quando usar cada um
- **Context com goroutines**: Como cancelar operações em andamento
- **Worker pool patterns**: Diferentes implementações de pools
- **Fan-out/Fan-in**: Padrões avançados de concorrência

### Contexto Kubernetes: Como Você Vai Usar Isso

Em operadores e controladores K8s, você verá goroutines para:

- **Watch de recursos**: Monitorar mudanças em pods, deployments, etc.
- **Processamento em lote**: Processar muitos recursos simultaneamente
- **Health checks**: Verificar status de múltiplos serviços
- **Reconciliação**: Processar fila de eventos em paralelo

```go
// Exemplo real de controller-runtime
func (r *PodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    // Cada reconciliação roda em uma goroutine separada
    // O controller gerencia automaticamente o worker pool
}
```

### Checklist de Conclusão do Dia 5

- □ Entendo que goroutine = thread leve (~2KB)
- □ Sei usar `go func()` para iniciar goroutines
- □ Uso `sync.WaitGroup` para sincronização
- □ Entendo o problema de passar variáveis de loop
- □ Implementei worker pool com limite de concorrência
- □ Completei o exercício do processador de pods
- □ Entendo como isso se aplica a operadores K8s

## DIA 6: Goroutines e Channels (Parte 2)

### Proposta

**Conceitos**:

- `Channels` = comunicação entre goroutines
- Buffered vs unbuffered
- `Select` para multiplexar
- Range sobre `channel`

**Analogia**: Channel é como uma fila (queue) thread-safe.

**Relação**: Channels permitem comunicação segura sem locks.

**Aplicação prática**:

```go
ch := make(chan string, 5) // buffered
go func() {
    ch <- "mensagem"
}()
msg := <-ch
```

**Erro comum**: Deadlock - enviar para channel sem receptor ou vice-versa.

**Exercício**:

- Faça pipeline: produtor → processador → consumidor usando channels.

**Para aprofundar**: Padrões de concorrência

### 






















...

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


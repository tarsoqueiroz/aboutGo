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

### O Que São Channels? Filas Thread-Safe

**Analogia**: Channels são como **esteiras rolantes** em uma fábrica:

- **Unbuffered**: Esteira onde o operador só coloca a peça quando o próximo operador está pronto para pegar
- **Buffered**: Esteira com capacidade para N peças, permitindo que o operador continue trabalhando mesmo se o próximo estiver ocupado

**Características**:

- Comunicação segura entre goroutines (sem locks manuais)
- Bloqueiam automaticamente quando necessário
- Podem ser bidirecionais ou direcionais (só enviar, só receber)

```go
package main

import "fmt"

func main() {
    // Criando um channel
    ch := make(chan string)  // Unbuffered (sem buffer)
    
    // Enviar e receber em goroutines separadas
    go func() {
        ch <- "mensagem"  // Envia
    }()
    
    msg := <-ch  // Recebe
    fmt.Println(msg)
}
```

### Unbuffered vs Buffered Channels

**Unbuffered (sem buffer)**:

- Só permite enviar quando há um receptor pronto
- Só permite receber quando há um emissor pronto
- **Sincronização garantida** - emissor e receptor se encontram

```go
package main

import (
	"fmt"
	"time"
)

// Unbuffered - comunicação síncrona
func unbufferedExample() {
	ch := make(chan int) // Sem buffer

	fmt.Println("Chamando func inline pra enviar via channel")
	go func() {
		ch <- 42 // BLOQUEIA até alguém receber
		fmt.Println("Enviado!")
	}()

	fmt.Println("Dando um tempo pra ver o sync")
	time.Sleep(2 * time.Second) // Simula atraso

	valor := <-ch // BLOQUEIA até alguém enviar
	fmt.Println("Recebido:", valor)
}

func main() {
	fmt.Println("Call unbuffered channel function")
	unbufferedExample()
	fmt.Println("Voltei do call da unbuffered channel function")
}
```

**Buffered (com buffer)**:

- Permite enviar até que o buffer esteja cheio
- Permite receber até que o buffer esteja vazio
- **Comunicação assíncrona** (até o limite do buffer)

```go
package main

import (
	"fmt"
)

// Buffered - comunicação assíncrona
func bufferedExample() {
	ch := make(chan int, 3) // Buffer de 3

	// Pode enviar 3 mensagens sem receptor
	fmt.Println("Enviando msg 1")
	ch <- 1
	fmt.Println("Enviando msg 2")
	ch <- 2
	fmt.Println("Enviando msg 3")
	ch <- 3
	fmt.Println("Enviadas")

	// Agora recebe
	fmt.Println("Recebendo msg 1")
	fmt.Println(<-ch) // 1
	fmt.Println("Recebendo msg 2")
	fmt.Println(<-ch) // 2
	fmt.Println("Recebendo msg 3")
	fmt.Println(<-ch) // 3
	fmt.Println("Recebidas")
}

func main() {
	fmt.Println("Call buffered channel function")
	bufferedExample()
	fmt.Println("Voltei do call da buffered channel function")
}
```

**Quando usar cada um**?

| Tipo | Uso  |
| :--- | :--- |
| **Unbuffered** | Sincronização exata, garantia de entrega      |
| **Buffered**   | Processamento em lote, amortecimento de picos |

### Select: Multiplexando Channels

**Select** permite esperar por múltiplos canais simultaneamente - como um "switch" para canais.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Duas goroutines enviando em canais diferentes
	go func() {
		fmt.Println("go func 1")
		time.Sleep(2 * time.Second)
		fmt.Println("go func 1: enviando msg")
		ch1 <- "mensagem do canal 1"
		fmt.Println("go func 1: msg enviada")
	}()

	go func() {
		fmt.Println("go func 2")
		time.Sleep(1 * time.Second)
		fmt.Println("go func 2: enviando msg")
		ch2 <- "mensagem do canal 2"
		fmt.Println("go func 2: msg enviada")
	}()

	// Select espera o primeiro canal que receber dados
	for i := 0; i < 2; i++ {
		fmt.Println("Quem mandou msg?")
		select {
		case msg1 := <-ch1:
			fmt.Println("Recebido do canal 1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Recebido do canal 2:", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout!")
			return
		}
	}
	fmt.Println("Saindo")
}
```

**Padrões com Select**:

```go
// Timeout
select {
case msg := <-ch:
    processar(msg)
case <-time.After(5 * time.Second):
    fmt.Println("Timeout!")
}

// Non-blocking (default)
select {
case msg := <-ch:
    processar(msg)
default:
    fmt.Println("Nenhuma mensagem disponível")
}

// Tentar enviar sem bloquear
select {
case ch <- valor:
    fmt.Println("Enviado!")
default:
    fmt.Println("Canal cheio, não enviou")
}
```

### Range sobre Channels

`range` pode iterar sobre canais até que eles sejam fechados.

```go
package main

import (
	"fmt"
	"time"
)

func produtor(ch chan<- int) {
	fmt.Println("Entrando no produtor")
	for i := 1; i <= 5; i++ {
		fmt.Println("Enviando:", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Fecha o canal quando terminar
	fmt.Println("Saindo do produtor")
}

func consumidor(ch <-chan int) {
	fmt.Println("Entrando no consumidor")
	for valor := range ch { // Loop até canal ser fechado
		fmt.Println("Recebido:", valor)
	}
	fmt.Println("Canal fechado!")
	fmt.Println("Saindo do consumidor")
}

func main() {
	fmt.Println("range Channel iniciado!")
	ch := make(chan int)

	fmt.Println("Chamando produtor")
	go produtor(ch)
	fmt.Println("Chamando consumidor")
	consumidor(ch)
}
```

**⚠️ Importante sobre fechar canais**:

- Só o **produtor** deve fechar o canal
- Nunca feche um canal que outros podem estar enviando
- Receber de um canal fechado retorna o valor zero
- Use `v, ok := <-ch` para verificar se está aberto

### Pipeline: Produtor → Processador → Consumidor

**Pipeline** é um padrão clássico onde dados fluem por uma série de estágios.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Estágio 1: Produtor - gera números
func produtor(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Printf("📦 Produtor enviou: %d\n", n)
			out <- n
			time.Sleep(100 * time.Millisecond)
		}
		close(out)
		fmt.Println("📦 Produtor finalizado")
	}()
	return out
}

// Estágio 2: Processador - dobra os números
func processador(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			resultado := n * 2
			fmt.Printf("⚙️  Processador: %d → %d\n", n, resultado)
			out <- resultado
			time.Sleep(50 * time.Millisecond)
		}
		close(out)
		fmt.Println("⚙️  Processador finalizado")
	}()
	return out
}

// Estágio 3: Consumidor - imprime os resultados
func consumidor(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for resultado := range in {
		fmt.Printf("✅ Consumidor recebeu: %d\n", resultado)
	}
	fmt.Println("✅ Consumidor finalizado")
}

// Pipeline com múltiplos workers em paralelo
func pipelineParalelo() {
	fmt.Println("\n=== PIPELINE COM WORKERS PARALELOS ===")

	// Entrada
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Estágio 1: Produtor
	prod := produtor(numeros...)

	// Estágio 2: Processadores paralelos (fan-out)
	numWorkers := 3
	processadores := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		processadores[i] = processador(prod)
	}

	// Estágio 3: Consumidor (fan-in)
	consumidorChan := make(chan int)
	var wg sync.WaitGroup

	// Fan-in: junta os resultados dos processadores
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, ch := range processadores {
			for v := range ch {
				consumidorChan <- v
			}
		}
		close(consumidorChan)
	}()

	// Consumidor final
	wg.Add(1)
	go consumidor(consumidorChan, &wg)

	wg.Wait()
}

// Pipeline com contexto (cancellation)
func pipelineComContexto() {
	// Veremos amanhã com Context
	fmt.Println("\n=== PIPELINE COM CONTEXTO (amanhã) ===")
}

func main() {
	fmt.Println("=== PIPELINE SIMPLES ===")
	// Pipeline simples
	prod := produtor(1, 2, 3, 4, 5)
	proc := processador(prod)
	var wg sync.WaitGroup
	wg.Add(1)
	consumidor(proc, &wg)
	wg.Wait()

	// Pipeline com workers paralelos
	pipelineParalelo()
}
```

### Direcionalidade: `chan<-` e `<-chan`

Especifica se um canal só pode enviar ou só receber - aumenta segurança.

```go
// Função que só ENVIA
func sender(ch chan<- int) {
    ch <- 42
    // <-ch  // ERRO! Não pode receber
}

// Função que só RECEBE
func receiver(ch <-chan int) {
    valor := <-ch
    // ch <- 42  // ERRO! Não pode enviar
}

// Função que pode enviar E receber
func bidirectional(ch chan int) {
    ch <- 42
    valor := <-ch
}
```

### Erros e Confusões Comuns

| Erro | Sintoma | Solução |
| :--- | :------ | :------ |
| Deadlock enviando sem receptor | `fatal error: all goroutines are asleep` | Use buffer ou goroutine para enviar |
| Deadlock recebendo sem emissor | `fatal error: all goroutines are asleep` | Garanta que alguém enviará ou feche o canal |
| Enviar em canal fechado        | `panic: send on closed channel`          | Apenas o produtor deve fechar |
| Esquecer de fechar canal       | Loop `range` nunca termina               | Feche quando não houver mais dados |
| Canal sem buffer em loop       | Deadlock ou performance ruim             | Use buffer para processamento em lote |

### Exercício para Fixar

**Objetivo**: Criar um pipeline de processamento de logs com 3 estágios.

**Instruções**:

- **Estágio 1 - Coletor de Logs (`produtor`)**:
  - Gera 100 mensagens de log simuladas
  - Cada mensagem tem: timestamp, nível (`INFO/WARN/ERROR`), mensagem
  - Envia para o próximo estágio
- **Estágio 2 - Filtro e Processador**:
  - Recebe logs
  - Filtra apenas logs com nível INFO (opcional: pode receber um filtro)
  - Converte para maiúsculas
  - Envia para o próximo estágio
- **Estágio 3 - Consumidor**:
  - Recebe logs processados
  - Imprime cada log com prefixo "`[PROCESSADO]`"
  - Mantém estatísticas (total recebido, total processado)
- **Bônus**: Use múltiplos workers no estágio 2 (fan-out) e combine os resultados (fan-in)

**Esboço da Solução**:

```go
type Log struct {
    Timestamp time.Time
    Level     string
    Message   string
}

// Funções de pipeline
func coletorLogs() <-chan Log { /* ... */ }
func filtroProcessador(in <-chan Log, filtro string) <-chan string { /* ... */ }
func consumidor(in <-chan string, wg *sync.WaitGroup) { /* ... */ }
```

### Solução

```go
package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Log - struct representando uma entrada de log
type Log struct {
	Timestamp time.Time
	Level     string
	Message   string
}

// Stage 1: Coletor de Logs (Produtor)
func coletorLogs(numLogs int) <-chan Log {
	out := make(chan Log, numLogs)
	levels := []string{"INFO", "WARN", "ERROR", "INFO", "INFO"}
	messages := []string{
		"Serviço iniciado",
		"Conexão estabelecida",
		"Timeout na requisição",
		"Cache atualizado",
		"Usuário autenticado",
		"Erro no banco de dados",
		"Processamento concluído",
		"Health check OK",
	}

	go func() {
		defer close(out)
		for i := 0; i < numLogs; i++ {
			log := Log{
				Timestamp: time.Now(),
				Level:     levels[i%len(levels)],
				Message:   fmt.Sprintf("%s [%d]", messages[i%len(messages)], i),
			}
			fmt.Printf("📝 Coletor: [%s] %s\n", log.Level, log.Message)
			out <- log
			time.Sleep(50 * time.Millisecond) // Simula coleta
		}
		fmt.Println("📝 Coletor finalizado")
	}()
	return out
}

// Stage 2: Filtro e Processador (com múltiplos workers)
func filtroProcessador(in <-chan Log, nivelFiltro string, numWorkers int) <-chan string {
	out := make(chan string, 100)
	var wg sync.WaitGroup

	// Worker function
	worker := func(id int) {
		defer wg.Done()
		for log := range in {
			// Filtra pelo nível
			if log.Level != nivelFiltro {
				continue
			}

			// Processa: converte mensagem para maiúsculas
			processed := fmt.Sprintf("[Worker %d] %s | %s",
				id,
				log.Timestamp.Format("15:04:05"),
				strings.ToUpper(log.Message))

			fmt.Printf("⚙️  Worker %d processou: %s\n", id, log.Message)
			out <- processed
			time.Sleep(20 * time.Millisecond) // Simula processamento
		}
	}

	// Inicia workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Fecha o canal de saída quando todos os workers terminarem
	go func() {
		wg.Wait()
		close(out)
		fmt.Println("⚙️  Processador finalizado")
	}()

	return out
}

// Stage 3: Consumidor
func consumidor(in <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	var totalRecebidos, totalProcessados int
	var mutex sync.Mutex

	var innerWg sync.WaitGroup
	innerWg.Add(1)

	go func() {
		defer innerWg.Done()
		for processed := range in {
			mutex.Lock()
			totalRecebidos++
			totalProcessados++
			mutex.Unlock()

			fmt.Printf("✅ Consumidor: %s\n", processed)
		}
	}()

	// Timeout para mostrar estatísticas mesmo se pipeline demorar
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("⏰ Timeout no consumidor")
	case <-func() <-chan struct{} {
		ch := make(chan struct{})
		go func() {
			innerWg.Wait()
			close(ch)
		}()
		return ch
	}():
		// Pipeline completou
	}

	fmt.Printf("📊 Estatísticas: %d logs processados\n", totalProcessados)
}

func main() {
	fmt.Println("=== PIPELINE DE PROCESSAMENTO DE LOGS ===")
	fmt.Println()

	// Stage 1: Coletor (100 logs)
	logs := coletorLogs(20)

	// Stage 2: Processador (filtra INFO, 3 workers)
	processed := filtroProcessador(logs, "INFO", 3)

	// Stage 3: Consumidor
	var wg sync.WaitGroup
	wg.Add(1)
	consumidor(processed, &wg)

	wg.Wait()
	fmt.Println("\n✅ Pipeline completo!")
}
```

### O que Estudar para Aprofundar

- **Padrões de concorrência**: Fan-out, Fan-in, Worker pools
- **Select com timeouts**: `time.After()`, `time.Ticker()`, `time.Tick()`
- **Channel ownership**: Quem cria, quem fecha, quem usa
- **Context com channels**: Cancela pipelines (Dia 8)
- **Benchmarks**: Performance de unbuffered vs buffered

### Contexto Kubernetes: Como Você Vai Usar Isso

Em operadores K8s, channels são usados para:

- **Watch events**: Processar eventos de recursos K8s
- **Work queues**: Filas de reconciliação (controller-runtime)
- **Log aggregation**: Coletar e processar logs de múltiplos pods
- **Health checks**: Monitorar status de serviços

```go
// Exemplo real: controller-runtime usa work queues
type RateLimitingInterface interface {
    Add(item interface{})
    Get() (item interface{}, shutdown bool)
    Done(item interface{})
}
// Internamente usa channels e goroutines
```

### Checklist de Conclusão do Dia 6

- □ Entendo a diferença entre buffered e unbuffered channels
- □ Sei usar `select` para multiplexar canais
- □ Uso `range` para iterar sobre canais
- □ Compreendo direcionalidade (`chan<-` e `<-chan`)
- □ Evito deadlocks (não enviar sem receptor)
- □ Completei o exercício do pipeline de logs
- □ Entendo quando fechar canais (só o produtor)

## DIA 7: Tratamento de Erros

### Proposta

**Conceitos**:

- `error` é uma interface
- Retorno múltiplo (`valor, err`)
- `Errors.Is` e `errors.As` para wrapping
- `Panic/recover` - NÃO USE em código normal

**Analogia**: Diferente de exceptions (`try/catch`), erros são valores de retorno.

**Relação**: Crucial para serviços long-running em K8s - erros não podem quebrar tudo.

**Aplicação prática**:

```go
result, err := doSomething()
if err != nil {
    return fmt.Errorf("processando pod %s: %w", podName, err)
}
```

**Erro comum**: Ignorar erros com _ - SEMPRE trate ou propague.

**Exercício**:

- Crie função que tenta conectar a um servidor (mock), retry com backoff exponencial.

**Para aprofundar**: Erros em GO

### Erros São Valores, Não Exceções

**Analogia**: Em Java/C#, erros são como **alarmes** que interrompem o fluxo (try/catch). Em Go, erros são como **notas fiscais** - você recebe junto com o produto e decide o que fazer com ela.

**Características**:

- `error` é uma interface com um único método: `Error() string`
- Funções retornam `(valor, error)` - o erro é o último valor
- Você **deve** verificar erros explicitamente
- Não há `try/catch` - você decide como lidar com cada erro

```go
package main

import (
	"errors"
	"fmt"
)

// error é uma interface
type error interface {
	Error() string
}

// Criando erros simples
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não permitida")
	}
	return a / b, nil
}

func main() {
	// Padrão: sempre verifique o erro
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Resultado:", resultado)

	// Erro esperado
	resultado, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Erro esperado:", err)
		// Não retorna, continua execução
	}
}
```

### Criando e Usando Erros Customizados

```package main
package main

import (
	"fmt"
	"time"
)

// Erro customizado com mais informações
type TimeoutError struct {
	Operation string
	Timeout   time.Duration
	Cause     error
}

func (e TimeoutError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("timeout na operação '%s' após %v: %v",
			e.Operation, e.Timeout, e.Cause)
	}
	return fmt.Sprintf("timeout na operação '%s' após %v",
		e.Operation, e.Timeout)
}

// Função que pode retornar erro customizado
func operacaoComTimeout(operation string, timeout time.Duration) error {
	// Simula uma operação que pode timeout
	if timeout < time.Second {
		return TimeoutError{
			Operation: operation,
			Timeout:   timeout,
			Cause:     fmt.Errorf("timeout muito curto"),
		}
	}
	return nil
}

func main() {
	err := operacaoComTimeout("conectar", 100*time.Millisecond)
	if err != nil {
		// Type assertion para acessar campos específicos
		if timeoutErr, ok := err.(TimeoutError); ok {
			fmt.Printf("⚠️  Erro customizado: %s\n", timeoutErr.Error())
			fmt.Printf("   Operação: %s\n", timeoutErr.Operation)
			fmt.Printf("   Timeout: %v\n", timeoutErr.Timeout)
		}
	}
}
```

### Wrapping de Erros: `fmt.Errorf` com `%w`

Wrapping permite adicionar contexto a erros sem perder a causa original.

```go
package main

import (
	"errors"
	"fmt"
)

// Função de baixo nível
func conectarServidor(addr string) error {
	// Simula erro de conexão
	return fmt.Errorf("servidor %s: conexão recusada", addr)
}

// Função de médio nível
func iniciarServico(servico string) error {
	err := conectarServidor("localhost:8080")
	if err != nil {
		// %w mantém a causa original
		return fmt.Errorf("iniciando serviço %s: %w", servico, err)
	}
	return nil
}

// Função de alto nível
func main() {
	err := iniciarServico("api-gateway")
	if err != nil {
		// Verifica se o erro contém uma causa específica
		if errors.Is(err, fmt.Errorf("conexão recusada")) {
			fmt.Println("Erro específico detectado!")
		}

		// Desempacota para ver a cadeia
		fmt.Printf("Erro completo: %v\n", err)

		// Desempacota para ver a causa raiz
		fmt.Printf("Causa raiz: %v\n", errors.Unwrap(err))
	}
}
```

### `errors.Is` vs `errors.As`

```go
package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validação falhou para %s (%v): %s",
		e.Field, e.Value, e.Message)
}

func validarCampo(field string, value interface{}) error {
	if value == nil || value == "" {
		return ValidationError{
			Field:   field,
			Value:   value,
			Message: "valor não pode ser vazio",
		}
	}
	return nil
}

func processarDados() error {
	err := validarCampo("nome", "")
	if err != nil {
		return fmt.Errorf("processando dados: %w", err)
	}
	return nil
}

func main() {
	err := processarDados()
	if err != nil {
		// errors.Is: verifica se o erro (ou sua causa) é de um tipo específico
		if errors.Is(err, ValidationError{}) {
			fmt.Println("❌ Erro de validação detectado")
		}

		// errors.As: extrai o erro para um tipo específico
		var valErr ValidationError
		if errors.As(err, &valErr) {
			fmt.Printf("⚠️  Campo: %s, Mensagem: %s\n",
				valErr.Field, valErr.Message)
		}
	}
}
```

### Panic e Recover: Use com MUITA CUIDADO

**Regra de Ouro**: `panic` é para **erros irrecuperáveis**, não para fluxo normal.

```go
package main

import "fmt"

// ❌ NUNCA faça isso para fluxo normal
func exemploErrado() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado:", r)
		}
	}()

	panic("algo deu errado") // Isso é um anti-pattern!
}

// ✅ Use panic apenas para erros que não podem ser recuperados
func exemploCorreto() {
	// Exemplo: arquivo de configuração obrigatório não existe
	// Aqui panic é aceitável porque o programa não pode continuar
	config := lerConfiguracaoObrigatoria()
	// ... continua
}

func lerConfiguracaoObrigatoria() string {
	// Se não encontrar, panic é aceitável
	// ...
	return "config"
}

// ✅ Em serviços long-running, use retry com erros, não panic
func operacaoComRetry() error {
	for i := 0; i < 3; i++ {
		err := operacaoArriscada()
		if err == nil {
			return nil
		}
		// Log do erro e tenta novamente
	}
	return fmt.Errorf("falha após 3 tentativas")
}

func operacaoArriscada() error {
	return fmt.Errorf("erro simulado")
}
```

### Padrão de Retry com Backoff Exponencial (Exercício Prático)

Este é um padrão essencial para sistemas que interagem com serviços externos.

```go
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Erro de conexão customizado
type ConnectionError struct {
	Address string
	Attempt int
	Cause   error
}

func (e ConnectionError) Error() string {
	return fmt.Sprintf("conexão com %s falhou na tentativa %d: %v",
		e.Address, e.Attempt, e.Cause)
}

// Simula conexão com servidor (sucesso ou falha aleatória)
func conectarServidor(addr string) error {
	// Simula latência de rede
	time.Sleep(50 * time.Millisecond)

	// 70% de chance de falha
	if rand.Float32() < 0.7 {
		return fmt.Errorf("servidor %s não respondeu", addr)
	}
	return nil
}

// Função com retry e backoff exponencial
func conectarComRetry(addr string, maxRetries int) error {
	var lastErr error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		// Tentativa de conexão
		err := conectarServidor(addr)
		if err == nil {
			fmt.Printf("✅ Conexão com %s estabelecida na tentativa %d\n",
				addr, attempt+1)
			return nil
		}

		// Erro na tentativa
		lastErr = ConnectionError{
			Address: addr,
			Attempt: attempt + 1,
			Cause:   err,
		}

		// Se for a última tentativa, retorna o erro
		if attempt == maxRetries {
			break
		}

		// Calcula backoff exponencial: 1s, 2s, 4s, 8s...
		backoff := time.Duration(1<<uint(attempt)) * time.Second
		// Adiciona jitter (variação aleatória) para evitar "thundering herd"
		jitter := time.Duration(rand.Intn(500)) * time.Millisecond
		waitTime := backoff + jitter

		fmt.Printf("⚠️  Tentativa %d falhou. Aguardando %v antes de tentar novamente...\n",
			attempt+1, waitTime)
		time.Sleep(waitTime)
	}

	return fmt.Errorf("falha ao conectar com %s após %d tentativas: %w",
		addr, maxRetries+1, lastErr)
}

// Versão com jitter e backoff mais sofisticado
func conectarComRetryAvancado(addr string, maxRetries int) error {
	var lastErr error
	baseDelay := 1 * time.Second
	maxDelay := 30 * time.Second

	for attempt := 0; attempt <= maxRetries; attempt++ {
		err := conectarServidor(addr)
		if err == nil {
			fmt.Printf("✅ Conexão com %s estabelecida na tentativa %d\n",
				addr, attempt+1)
			return nil
		}

		lastErr = ConnectionError{
			Address: addr,
			Attempt: attempt + 1,
			Cause:   err,
		}

		if attempt == maxRetries {
			break
		}

		// Backoff exponencial com jitter
		// Formula: min(maxDelay, baseDelay * 2^attempt)
		backoff := baseDelay * time.Duration(1<<uint(attempt))
		if backoff > maxDelay {
			backoff = maxDelay
		}

		// Jitter: adiciona +/- 20% de variação
		jitterRange := float64(backoff) * 0.2
		jitter := time.Duration(rand.Float64()*jitterRange*2 - jitterRange)
		waitTime := backoff + jitter

		fmt.Printf("⚠️  Tentativa %d falhou. Aguardando %v antes de tentar novamente...\n",
			attempt+1, waitTime)
		time.Sleep(waitTime)
	}

	return fmt.Errorf("falha ao conectar com %s após %d tentativas: %w",
		addr, maxRetries+1, lastErr)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== RETRY COM BACKOFF EXPONENCIAL ===")
	err := conectarComRetry("api.k8s.local", 4)
	if err != nil {
		fmt.Printf("❌ Erro final: %v\n", err)

		// Verifica se é ConnectionError
		var connErr ConnectionError
		if errors.As(err, &connErr) {
			fmt.Printf("   Última tentativa: %d, Endereço: %s\n",
				connErr.Attempt, connErr.Address)
		}
	}

	fmt.Println("\n=== RETRY AVANÇADO COM JITTER ===")
	err = conectarComRetryAvancado("database.k8s.local", 5)
	if err != nil {
		fmt.Printf("❌ Erro final: %v\n", err)
	}
}
```

### Erros e Confusões Comuns

| Erro | Sintoma | Solução |
| :--- | :------ | :------ |
| Ignorar erro com `_`           | `result, _ := doSomething()`       | Sempre trate ou propague o erro   |
| Não propagar erro com `%w`     | Perde a causa original             | Use `fmt.Errorf("...: %w", err)`  |
| Usar `panic` para fluxo normal | Programa quebra desnecessariamente | Use error para erros esperados    |
| Comparar erros com `==`        | Não funciona com wrapped errors    | Use `errors.Is()` e `errors.As()` |
| Não verificar erro em `defer`  | `defer file.Close()` pode falhar   | Verifique erros em `defer` também |

### Exercício para Fixar

**Objetivo**: Criar um sistema de conexão com retry e circuit breaker.

**Instruções**:

- **Crie uma função `conectarAPI()` que**:
  - Simula conexão com API externa
  - Falha com 60% de chance
  - Pode retornar diferentes tipos de erro: Timeout, AuthError, RateLimitError
- **Implemente retry com backoff**:
  - Máximo de 5 tentativas
  - Backoff exponencial: 1s, 2s, 4s, 8s, 16s
  - Com jitter
- **Implemente Circuit Breaker**:
  - Se houver 3 falhas consecutivas, abre o circuito
  - Fica aberto por 10 segundos
  - Depois tenta novamente (half-open)
- **Teste diferentes cenários**:
  - Sucesso na primeira tentativa
  - Falhas com sucesso no retry
  - Falhas que abrem o circuit breaker

**Esboço da Solução**:

```go
type CircuitBreaker struct {
    state           string  // "closed", "open", "half-open"
    failures        int
    maxFailures     int
    timeout         time.Duration
    lastFailureTime time.Time
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    // Implementa lógica do circuit breaker
}
```

### Solução

```package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Tipos de erros
type TimeoutError struct {
	Duration time.Duration
}

func (e TimeoutError) Error() string {
	return fmt.Sprintf("timeout após %v", e.Duration)
}

type AuthError struct {
	Reason string
}

func (e AuthError) Error() string {
	return fmt.Sprintf("erro de autenticação: %s", e.Reason)
}

type RateLimitError struct {
	RetryAfter time.Duration
}

func (e RateLimitError) Error() string {
	return fmt.Sprintf("rate limit: aguarde %v", e.RetryAfter)
}

// Circuit Breaker
type CircuitBreaker struct {
	state           string
	failures        int
	maxFailures     int
	timeout         time.Duration
	lastFailureTime time.Time
	mutex           chan struct{} // Simples mutex com channel
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:       "closed",
		maxFailures: maxFailures,
		timeout:     timeout,
		mutex:       make(chan struct{}, 1),
	}
}

func (cb *CircuitBreaker) lock()   { cb.mutex <- struct{}{} }
func (cb *CircuitBreaker) unlock() { <-cb.mutex }

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.lock()
	defer cb.unlock()

	// Verifica estado do circuit breaker
	if cb.state == "open" {
		if time.Since(cb.lastFailureTime) > cb.timeout {
			fmt.Println("🔄 Circuit half-open: testando novamente...")
			cb.state = "half-open"
		} else {
			return fmt.Errorf("circuit breaker aberto (falhas: %d)", cb.failures)
		}
	}

	// Executa a função
	err := fn()

	if err != nil {
		cb.failures++
		cb.lastFailureTime = time.Now()

		// Verifica se deve abrir o circuito
		if cb.failures >= cb.maxFailures {
			cb.state = "open"
			fmt.Printf("🔴 Circuit breaker ABERTO (falhas: %d)\n", cb.failures)
		}
		return fmt.Errorf("chamada falhou: %w", err)
	}

	// Sucesso: reset
	if cb.state == "half-open" {
		fmt.Println("🟢 Circuit half-open: sucesso! Circuito fechado")
	}
	cb.state = "closed"
	cb.failures = 0
	return nil
}

// API simulada
func chamadaAPI() error {
	// Simula latência
	time.Sleep(50 * time.Millisecond)

	// Diferentes tipos de erro
	randVal := rand.Float32()
	switch {
	case randVal < 0.2:
		return TimeoutError{Duration: 5 * time.Second}
	case randVal < 0.35:
		return AuthError{Reason: "token expirado"}
	case randVal < 0.5:
		return RateLimitError{RetryAfter: 2 * time.Second}
	case randVal < 0.7:
		return fmt.Errorf("erro interno do servidor")
	default:
		return nil // Sucesso
	}
}

// Função com retry e circuit breaker
func chamadaComRetry(cb *CircuitBreaker, maxRetries int) error {
	var lastErr error
	baseDelay := 1 * time.Second

	for attempt := 0; attempt <= maxRetries; attempt++ {
		err := cb.Call(chamadaAPI)
		if err == nil {
			fmt.Printf("✅ Chamada bem-sucedida na tentativa %d\n", attempt+1)
			return nil
		}

		lastErr = err

		// Verifica se é erro de rate limit
		var rateErr RateLimitError
		if errors.As(err, &rateErr) {
			fmt.Printf("⏳ Rate limit: aguardando %v\n", rateErr.RetryAfter)
			time.Sleep(rateErr.RetryAfter)
			continue
		}

		if attempt == maxRetries {
			break
		}

		// Backoff exponencial com jitter
		backoff := baseDelay * time.Duration(1<<uint(attempt))
		jitter := time.Duration(rand.Intn(500)) * time.Millisecond
		waitTime := backoff + jitter

		fmt.Printf("⚠️  Tentativa %d falhou: %v. Aguardando %v...\n",
			attempt+1, err, waitTime)
		time.Sleep(waitTime)
	}

	return fmt.Errorf("falha após %d tentativas: %w", maxRetries+1, lastErr)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== SISTEMA DE CONEXÃO COM RETRY E CIRCUIT BREAKER ===")

	cb := NewCircuitBreaker(3, 5*time.Second)

	// Cenário 1: Múltiplas chamadas
	for i := 1; i <= 10; i++ {
		fmt.Printf("\n--- Chamada %d ---\n", i)
		err := chamadaComRetry(cb, 3)
		if err != nil {
			fmt.Printf("❌ Erro final: %v\n", err)
		}

		// Espera entre chamadas
		time.Sleep(1 * time.Second)
	}
}
```

### O que Estudar para Aprofundar

- **Erros customizados com `errors.New()`**: [Error handling best practices](https://go.dev/blog/error-handling)
- **Wrapping de erros em cadeia**: [Go 1.13 error wrapping](https://go.dev/blog/go1.13-errors)
- **Sentry/Logging**: Como integrar com ferramentas de monitoramento
- **Circuit Breaker patterns**: [Resilience patterns](https://resilience4j.readme.io/docs/circuitbreaker)
- **Error types em client-go**: Veja como o Kubernetes trata erros

### Contexto Kubernetes: Como Você Vai Usar Isso

Em operadores K8s, o tratamento de erros é crítico:

```go
// Exemplo real: controller-runtime
func (r *PodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    // Busca o pod
    var pod corev1.Pod
    if err := r.Get(ctx, req.NamespacedName, &pod); err != nil {
        if errors.Is(err, &NotFoundError{}) {
            // Pod não existe - requeue com delay
            return ctrl.Result{RequeueAfter: 5 * time.Minute}, nil
        }
        // Erro inesperado - requeue imediato
        return ctrl.Result{}, fmt.Errorf("erro ao buscar pod: %w", err)
    }
    
    // Processa o pod com retry
    if err := r.processPodWithRetry(&pod); err != nil {
        // Erro pode ser temporário - requeue com backoff
        return ctrl.Result{RequeueAfter: r.calculateBackoff()}, nil
    }
    
    return ctrl.Result{}, nil
}
```

### Checklist de Conclusão do Dia 7

- □ Entendo que `error` é uma interface, não exceção
- □ Sei criar erros customizados com informações adicionais
- □ Uso `fmt.Errorf("%w")` para wrapping
- □ Sei diferença entre `errors.Is` e `errors.As`
- □ Compreendo que `panic/recover` é para casos extremos
- □ Implementei retry com backoff exponencial
- □ Entendo o padrão Circuit Breaker
- □ Completei o exercício prático

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


# Go: Fundamentos de uma aplicação web

## About

> `https://cursos.alura.com.br/course/go-lang-web`

Faça esse curso de GoLang e:

- Crie uma aplicação web do zero com Go dentro das principais convenções
- Aprenda o que são structs na prática
- Saiba como conectar sua aplicação Go com banco de dados
- Aprofunde seus conhecimentos na linguagem criada pelo Google
- Melhore seu código com partials

## Servidor e struct de produtos

### Introdução

[00:00] Olá meu nome é Guilherme Lima e eu serei seu instrutor neste curso de Desenvolver uma Aplicação Web com Go. Vamos desenvolver uma aplicação web onde eu vou ser capaz de criar produtos e salvar suas informações esses produtos no banco de dados.

[00:16] Então vou criar um produto tênis, vou falar que a uma descrição dele é “bom para corrida”, o preço dele vai ser 199, quantidade um. Eu posso salvar essa informação, esse produto que eu criei está lá no nosso banco de dados.

[00:30] Vamos ser capaz de editar esse produto, tem quantidade um, a quantidade 2 alteramos e seremos capazes de deletar também o nosso produto. “Tem certeza que remover esse produto?”, dou um ok e aquele produto sai.

[00:44] Vamos desenvolver tudo isso utilizando nenhum framework, apenas a linguagem Go, todas as nossas tabelas, nossos layouts, as nossas páginas HTML, os nossos códigos do servidor, vamos fazer utilizando a biblioteca padrão do Go.

[00:58] Vamos importar umas outras bibliotecas também para criar essas funcionalidades para nosso projeto, vamos modularizar nosso projeto na convenção do MVC.

[01:08] Então teremos uma rota, criar um arquivo para cuidar só das nossas rotas, quem atende essas rotas, quem vai atender essas novas requisições, criaremos uma pasta de controle onde cada função do nosso controle vai ter a responsabilidade que cabe ao controle.

[01:26] Quando precisarmos fazer acesso ao banco de dados vamos pedir para um modelo nosso, que faz a conexão com banco de dados e traz as informações importantes.

[01:37] Esse curso vamos dar continuidade nos ensinamentos, nos nossos conhecimentos da linguagem Go. Vamos lá!

### Workspace e pacotes

[00:00] Vamos iniciar nosso projeto. A primeira coisa que eu vou fazer vai ser abrir o Visual Studio Code que já configuramos, na atividade anterior. Vou apertar command J e acessa o terminal. Antes começar a criar todo nosso código, existe um local específico para armazenar toda nossa aplicação Go.

[00:20] Todo nosso código-fonte tem um local específico, existe uma convenção de onde deve colocar o nosso código Go. Vou acessar com cd $GOPATH - tudo maiúsculo, acessei do LS, Observe que eu tenho que uma pasta chamada Go.

[00:38] Eu vou acessar essa pasta cd go, eu tenho três diretórios – bin, pkg e o crc, vamos recordar o que que é cada diretório que fica armazenado: o diretório bin é referente aos arquivos binários, ou seja, onde ficam todos os arquivos executáveis das nossas aplicações feitas em Go.

[01:06] O diretório pkg é referente ao pacote – de package em inglês, o pacote contém objetos referentes a todos os pacotes que têm na nossa aplicação e esse src de source - na tradução em português de fonte, é onde armazena todo o código-fonte, todos os programas escritos em Go, sejam nossos ou de programas de pacotes de terceiros.

[01:33] Ficam sempre armazenados dentro dessa pasta src, ou seja, nessa pasta que criamos os nossos projetos, então vou acessar essa pasta src, dou um ls, tem dois diretórios, um GitHub e o Golang.

[01:48] O Golang é onde fica tudo para o funcionamento da nossa linguagem, tudo que vamos utilizar, e é onde fica armazenado todo o código-fonte das aplicações feitas em Go. Então cd github.com dou um ls e repara que já tenho pastas dentro.

[02:12] Essa pasta src, esse diretório fica também os pacotes de terceiros instalados no nosso sistema. Dentro dessa pasta eu vou criar um diretório que seja o nome do meu usuário do GitHub, por exemplo, e dentro desse diretório vou colocar todos os meus projetos feitos em Go.

[02:32] Nesse caso eu já tenho o meu usuário GitHub que é Alura, se você não tem, você pode colocar mkdir e o nome do seu usuário, por exemplo, meu nome é Guilherme, vou colocar Guilherme, dou um enter e ele vai criar uma pasta nesse diretório e dentro dessa pasta posso salvar os meus projetos.

[02:49] Como eu já tenho – se você não tem você cria o seu, eu vou acessar cd alura, na sua casa você coloca cd e o nome desse diretório que você criou; vou dar enter e se dou ls repara que eu não tenho nada nesse diretório.

[03:02] E aí que a nossa aplicação de fato vai começar, para ficar legal, para não ficar com muitos diretórios na aplicação, vou escrever code . vou dar enter e ele vai abrir novo Visual Studio Code dentro deste diretório.

[03:22] Repara que já estou até no diretório da Alura, vou fechar o de trás e vou abrir um diretório “grandão”, se der no terminar ls, não vai mostrar nada, porque não tem arquivo nenhum. Aqui nossa brincadeira, nossa grande caminhada começa.

[03:42] Quero criar uma loja que tenha vários produtos, quero conseguir visualizar vários produtos na minha loja, então, eu tenho o nome do produto, tem a descrição - uma descrição breve, preço e quantidade.

[03:58] Assim como se eu digito no Google o teu www.google.com eu caio num endereço, alguém consegue atender essa requisição. Digita o site da Alura, alura.com.br, eu caio no site da Alura.

[04:11] Eu quero fazer a mesma coisa só que da minha máquina, não vou colocar o meu projeto na web ainda; eu quero da minha máquina conseguir ver o que que as outras pessoas que vão acessar o site da minha loja verão.

[04:24] Quando nós queremos utilizar o nosso próprio computador para fazer teste na próxima máquina, utiliza o endereço local host nesse caso utilizo a porta 8000. Quando eu digito isso não tem nada, mas eu também não criei nenhuma página.

[04:42] Vou criar uma página. O nosso curso não tem o foco em páginas HTML, CSS e assim por diante, então se você tem interesse você pode fazer outros cursos na Alura em nossa formação de front end.

[04:57] Eu vou criar index.html para ser a página principal da aplicação, vou apertar command J para fechar aquele nosso terminal, em HTML vou selecionar a segunda opção para ele trazer algumas informações.

[05:13] Algumas coisas feitas em HTML, o título vou chamar de loja, no lugar de localhost vai ficar loja, vou colocar H1 e vou chamar de Alura loja.

[05:28] Depois eu coloco nossos produtos, quero conseguir visualizar isso, se eu atualizo minha página repare que não vai acontecer nada, porque ainda não subiu o servidor.

### Subindo o servidor

[00:00] Eu quero que a página quando eu digito localhost, na porta 8000 suba o nosso servidor, então vou criar um novo arquivo, ele vai ser o nosso código Go principal, por convenção chama de main.go.

[00:15] Não coloquei nada e ele já está com uma marcação vermelha, estou esperando saber de qual pacote é todo esse código que você vai criar, do próprio Gol e precisa especificar qual é o pacote que nós estamos querendo.

[00:32] É referente a qual pacote aquele código pertence? Pertence ao pacote Main. Eu vou criar uma função, chamá-la de main e dentro dessa função mesmo eu quero subir um servidor, quero que alguém consiga ouvir uma requisição e responder essa requisição.

[00:52] Então para conseguir vou colocar aqui http.ListenAndServe tem alguém que que é capaz de ouvir e responder essa nossa requisição, ou seja, ouça e sirva essa requisição, vou colocar na porta “8000”, e eu não preciso passar nenhuma informação para essa nossa requisição.

[01:15] Vou colocar como New, vou salvar, ele já importou a biblioteca que pertence esse http, se eu volto e atualiza não acontece nada, vou executar esse programa, vou dar “Control + L” para limpar, um LS que tem aqui index.html e vou executar esse Go.

[01:40] Vou falar, “Go executa go run main.go”, quando executa repara que no meu terminal ficou travado, vou atualizar a página, ele falou 404 - página não encontrada. Subimos o servidor, mas onde referenciamos que essa nossa página index precisa ser respondida?

[02:08] Não fizemos isso, então eu vou destravar, interromper o servidor com “control + C”, eu trago o servidor, vou diminuir para visualizar melhor, deixar maior nossa fonte. Deu um erro de 404, eu não encontrei essa página; precisa vincular essa página.

[02:36] As páginas em um projeto Go chama de template, quer dizer que essa página é um template que é possível de carregar dentro do Go. A primeira coisa que eu vou fazer no nosso projeto não teremos apenas uma página e sim algumas páginas.

[02:55] Se deixar todas as páginas vai ficar um pouco bagunçado, se eu tenho uma página e o main.go, se eu tiver duas páginas, três, quatro, trinta, vai ficar difícil de encontrar a página.

[03:11] Para organizar o nosso projeto de uma forma melhor, eu vou criar uma pasta e vou chamar de templates e vou colocar todos os nossos arquivos HTML Vou criar uma variável que vai conseguir trazer todos os meus templates desse nosso diretório.

[03:37] Eu vou chamar templates de var temp. Todos os meus templates estarão armazenados dentro dessa variável, eu preciso conseguir trazer todos eles para que eles sejam executados; para que a pessoa quando acessar a nossa página consiga visualizar o nosso template.

[04:08] Tem uma função que faz isso, consegue visualizar, a função pertence ao pacote template e se chama must; o must encapsula todos os templates e devolve dois retornos.

[04:32] Ele vai devolver o template e uma mensagem de erro, se não tiver erro conseguimos carregar. Se ele não conseguiu trazer algum ele vai trazer a mensagem de erro.

[04:44] Temos o template, já encapsulamos todos eles, mas onde eles estão? Para isso utilizamos mais uma função chamada template Parse Glob, vou passar o caminho de onde estão todos nossos templates que é na pasta templates.

[05:15] Eu quero carregar na pasta todos os templates template.Must(template.ParseGlob(“templates/*.html”), tudo que tiver .html eu quero trazer para minha variável.

[05:29] Eu coloquei o símbolo do asterisco, se que tiver tudo que tiver .html ele vai trazer para mim e vai renderizar como um template, agora eu consegui de fato atender uma requisição e encaminhar, para encaminhar para minha página index eu preciso chamar novamente a pasta http.

[05:51] Temos uma função chamada HandleFunc, ela trouxe a documentação dela e o primeiro parâmetro que ela espera é “qual é o caminho que você quer atender?”. Qual é a requisição que você quer atender, eu quero atender a requisição de barra.

[06:10] Sempre que tiver barra eu quero que alguém consiga atender a requisição, vamos criar uma função para atender. Vou chamar essa função de index. Toda vez que tiver uma requisição com barra essa nossa função vai mandar para o index.

[06:29] Vou criar uma função index, ela vai precisar de dois parâmetros: o primeiro Response Writer que por convenção chama w http.ResponseWriter – ou seja toda vez que tiver uma requisição eu quero conseguir escrever, mostrar para a pessoa que vai mexer no site quem vai responder.

[07:10] Segundo parâmetro vamos criar apontando para http request, sempre que tiver uma requisição vai ter esses dois parâmetros, que na verdade nas outras requisições também vão ser necessários.

[07:30] Carregou todos os templates, falou toda vez que tiver uma requisição para barra quem vai atender vai ser index. Queremos de fato executar o nosso template de index, então vou chamar de temp.ExecuteTemplate; cuidado com as letras maiúsculas e minúsculas.

[07:45] Vou passar parâmetro para essa função para de fato executar o nosso template, vou passar quem consegue passar a nossa resposta que é o nosso w; nosso responsewriter, vai passar também quem ele vai exibir, mostrar, o index.

[08:19] E por último quer passar alguma informação para nossa página index. Vou salvar e ele já trouxe o http. Parece que está tudo certinho para dar um Command J para conseguir visualizar, vou limpar meu terminal, veremos muito nesse curso.

[08:51] Colocar na pasta certinho, eu tenho main e template. Vou executar o main e pedir para o go run – roda o arquivo main.go. Ele ficou travado, legal. Vou atualizar, ele não mostrou nenhuma mensagem, não tem nada na verdade para conseguir executar.

[09:12] Estranho. Por que não estamos conseguindo ver o index.html? Para conseguir executar esse index precisa embedar um código HTML, temos nosso código html funciona da mesma forma existe um código Go ali dentro também que consegue fazer associação de tudo que HTML para conversar com o Go também.

[09:46] Vou pegar o index.html vou colocar duas chaves, vou definir que isso vai ser o meu index, então {{define “Index”}}, lá no final eu vou colocar mais duas chaves e vou escrever {{end}}, acaba nosso index.

[10:10] Vou salvar esse arquivo, o index vou deixar o I maiúsculo, salvei dos dois lados, eu passei o Index maiúsculo e referente ao Index maiúsculo, vou parar meu servidor “Control + C” e vou executar de novo com run.

[10:34] Ele atualizou a página. Vamos ver, apareceu lá o Alura loja, legal, só que eu não quero exibir apenas Alura loja, eu quero exibir alguns produtos na minha loja, então rapidamente eu vou criar um table.

[10:47] Dentro dessa table eu vou ter umathread para ficar bonito, dentro dela eu vou ter uma linha tr e dentro dessa tr eu quero ter para cada produto um nome, uma descrição, preço, quantidade.

[11:08] Então vou colocar th vezes quatro, eu quero quatro ths, um nome, uma descrição, preço e para finalizar eu quero uma quantidade. Criei a minha thread, vou criar agora o tbody – o corpo da minha tabela.

[11:36] Dentro do corpo eu vou ter tr que vai ser o meu primeiro produto e dentro desse tr vou ter quatro td, td*4, deixei aberto, agora sim bem melhor, o produto que vou criar vai ser uma camiseta.

[05:29] Eu gosto bastante de camiseta, vou colocar uma camiseta “bem bonita” que é a descrição, o preço dela vai ser 29 reais e a quantidade vou falar que eu tenho 10; vou criar mais um produto, último produtos tr.

[12:10] Eu tenho td vezes quatro, os quatro campos, vou criar o meu próximo produto vai ser um notebook “muito rápido”, o preço dele é 1999 e eu tenho só 2 em estoque, salvei, Command J.

[12:33] Voltei para o terminal, vou parar aplicação “Control + C”, vou limpar “Control + L”, rodei de novo, quando volto para página e atualizo eu tenho os produtos, está estranho o visual, mas consegui criar uma página HTML.

[12:56] Vincular a página HTML junto com o nosso código Go e tem aqui o nosso servidor Go funcionando.

### Adicionando o bootstrap

Vamos melhorar o visual da nossa aplicação?

Para isso, vamos adicionar o [bootstrap via CDN](https://getbootstrap.com/docs/4.3/getting-started/download/) e utilizar este incrível framework web com código-fonte aberto para desenvolvimento de componentes de interface e melhorar nosso front-end.

- Altere o conteúdo do arquivo index.html, substituindo o código desenvolvido pelo conteúdo abaixo:

```html
{{define "Index"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T"
        crossorigin="anonymous">
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM"
        crossorigin="anonymous"></script>
    <title>Alura loja</title>
</head>
<nav class="navbar navbar-light bg-light mb-4">
    <a class="navbar-brand" href="/">Alura Loja</a>
</nav>

<body>
    <div class="container">
        <section class="card">
            <div>
                <table class="table table-striped table-hover mb-0">
                    <thead>
                        <tr>
                            <th>Nome</th>
                            <th>Descrição</th>
                            <th>Preço</th>
                            <th>Quantidade</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>Camiseta</td>
                            <td>Bem bonita</td>
                            <td>29</td>
                            <td>10</td>
                        </tr>
                        <tr>
                            <td>Notebook</td>
                            <td>Muito rápido</td>
                            <td>1999</td>
                            <td>2</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </section>

    </div>
</body>

</html>
{{end}}
```

### Criando struct de produtos

[00:00] Na atividade anterior, melhoramos o visual do site, adicionando Bootstrap, vamos ver como ficou. Vamos só ver como é que ficou, subiu o nosso servidor com Go run. Alura loja,

[00:19] Só que tem algo interessante, um ponto que podemos melhorar no nosso código que é o seguinte. Sempre que eu quero criar um produto novo eu vou lá no HTML, cria uma linha, uma nova tr, depois cria 4 tds, onde a primeira é referente ao nome, a descrição, preço e a quantidade.

[00:35] Seguindo essa ordem. Repara que está um pouco difícil para eu criar os meus produtos têm que mexer direto no código HTML. Se eu quiser criar outro produto de novo mais uma td, mais tr, mais quatro tds.

[00:46] Não está legal isso, será que existe uma forma da melhorar a criação dos nossos produtos? E a resposta é sim. Podemos criar uma estrutura, modelo de produtos onde cada produto já sabemos quais são os campos deles, tem o nome, descrição, preço e quantidade.

[01:03] Isso que vamos fazer agora, criar uma estrutura lá no nosso código principal, na nossa main para eu criar uma estrutura que utiliza a palavrinha do Go struct, struct é simplesmente uma coleção de variáveis que formam o novo tipo, quero criar um tipo e vai usar também essa palavrinha do Go.

[01:31] Como é o nome desse tipo de produto? Uma estrutura de produto, abre chaves e fecha chaves. Vamos criar a nossa estrutura de produto que vai ter um nome, o tipo dele, é uma string, a descrição também uma string.

[02:06] Também tem um preço que vai ser float 64 e temos uma quantidade, essa quantidade é 20. Criamos novo tipo, uma estrutura nova - essa estrutura tem o nome, descrição, preço, quantidade, todo esse conjunto de variáveis chama de produto.

[02:35] Agora eu quero de fato criar o meu tipo de produto, criar produtos utilizando este modelo, essa estrutura e vamos fazer agora, para os meus produtos eu não quero criar simplesmente criar um produto, eu quero uma lista de produtos.

[02:52] Eu quero no Go um slice de produtos, eu quero escrever produtos := vou atribuindo nessa variável uma lista. Vou colocar a colchetes, uma lista de produto e todos os produtos estarão dentro dessa chave.

[03:18] Eu vou criar o meu primeiro produto criando uma chave. Vamos ver duas formas de criar os nossos produtos, eu posso criar assim nome:, camiseta, tem também a descrição, que vou chamar : abre e fecha chaves – por quê?

[03:40] Chaves porque estou trabalhando com string, a descrição vai ser uma camiseta azul bonita, “bem bonita”, para ficar diferente daquela que já tinha. Preço da camiseta é um pouquinho mais caro, é 39.

[04:06] A quantidade dessa camiseta 5. Eu vou criar esse novo produto de uma forma diferente, eu passei os campos descrição, preço, quantidade, posso passar direto o conteúdo valor desse meu novo produto.

[04:28]. Então vai ser um tênis confortável ,duas strings, o preço dele vai ser 99, não está muito caro, está na promoção na loja 89 e a quantidade dele vai ser 3. Vou criar um produto e esse produto vai ser um fone, um fone “muito bom” e eu tenho o preço dele é 59.

[05:18] Eu tenho apenas dois. Eu cria minha lista de produtos, no final eu também coloco uma virgula, mesmo sendo nessa linha. Cria uma lista de produtos. Como que eu mando esses meus produtos para index.

[05:32] Eu não quero visualizar os produtos que estão lá no meu HTML direto, eu quero visualizar os produtos que eu acabei de instanciar na minha struct de produto, repara que no nosso método index está passando executa esse template do index.

[05:52] Ele vai responder, vai mostrar na tela, escrever na tela o que se define como nosso HTML e vai passar para ele de dado - não está passando nada, criou instancia de produto, mas não está passando esses valores.

[06:08] Eu vou fazer no lugar do nil eu vou passar os produtos, certo se eu fizer só isso vai funcionar? Vamos testar para ver. Eu salvei aparentemente certinho, eu vou parar o meu servidor e vou rodar o meu servidor mais uma vez. Vamos ver se vai visualizar produtos diferente.

[06:26] Não visualizou. Mas passou esses dados, passou para index, mas não está conseguindo ver. Em que momento que nós alteramos o conteúdo dela? Ou de onde que está de fato trazendo as propriedades da nossa lista de produtos para cá, não fizemos isso.

[06:51] Para começar vou minimizar o nosso terminal, vou remover esse primeiro e segundo produto, vou ficar só com um e vamos fazer um bem legal. Vamos fazer um range, vamos pegar cada listinha que criamos, cada produto e vai falar cada produto vai ser uma nova tr, quatro trs, uma td e quatro tds de forma dinâmica.

[07:13] Vou remover antes do nosso tbody tem o nosso thead; eu vou colocar um código Go. Coloquei duas chaves, vou escrever range e vou dar um ponto, porque eu quero trazer todos os valores, eu quero range de produtos na minha tela.

[07:38] No lugar de exibir o conteúdo camiseta, eu quero exibir o nome do produto que eu estou passando, esse nome é o nome da lista, eu removo o nome, incluo o código Go mais uma vez e coloco ponto, escrevo nome. Faça o mesmo para os outros também.

[07:57] Eu vou remover a descrição e vou colocar .descricao e fecho, sempre com duas chaves e utilizando o ponto; no preço remove, duas chaves, ponto preço e para finalizar a quantidade, mesmo procedimento.

[08:27] Antes de salvar precisa fechar, fala eu quero que você pegue cada produto da lista e coloca o produto um vai ser o nome, descrição, preço 1, quantidade 1, acabou esse produto.

[08:50] Vamos fechar essa nossa range, embaixo coloco a palavra end, para cada produto vai ter essa repetição, salvei, vou no meu main.go, vou ver o meu terminalzinho de novo, Command J.

[09:04] Comando “Control + C” para parar o terminal, “Control + L” para limpar, rodei de novo, agora vamos ver se funcionou mesmo. Vou atualizar, está lá. Camiseta azul bem bonita, o tênis confortável, fone muito bonito. Será que funciona mesmo? Vamos testar e criar mais um produto.

[09:21] No main.go eu quero criar um produto novo, coloco as chaves e vou colocar um nome engraçado para ver. Vai ser “Produto novo”, a descrição é muito legal, o preço dele está 1,99, e a quantidade dele eu tenho só um, é superpromoção.

[09:54] Salvei, vem no terminal, cancela meu terminal, derrubo, subo de novo, eu esqueci da virgula no final, não pode, vamos ver, rodei, voltei, vamos ver se aparece produto novo e apareceu o produto novo.

[10:12] Repara que está gerando uma tabela dinâmica de produtos agora, utilizando a struct.

### Faça como eu fiz na aula

**Sua vez!**

Vamos criar uma loja, com vários produtos diferentes, onde as pessoas podem acessar e ver uma lista com todos meus produtos e alguns detalhes deles, como o nome, descrição, preço e quantidade.

> Para iniciar o desenvolvimento da aplicação, crie um servidor com go e uma página html com uma tabela de produtos. Em seguida, deixe a tabela de produtos dinâmica por meio de instâncias de uma struct, com base no conteúdo visto nos vídeos.

**Resumo do código**

- Para criar o servidor e conseguir exibir uma página html, crie um arquivo main.go e utilize as funções HandleFunc e ListenAndServe do pacote http, conforme ilustra a imagem abaixo:
package main

```go
import (
    "html/template"
    "net/http"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
    http.HandleFunc("/", Index)
    http.ListenAndServe(":8000", nil)
}
func Index(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "Index", nil)
```

- Em seguida, crie uma pasta para armazenar a página html chamada templates e uma página index.html, definindo entre 2 chaves {{ }} a referência passada na função ExecuteTemplate.

```html
{{define "Index"}}
<!DOCTYPE html>

código omitido...

</html>
{{end}}
```

- Para organizar melhor a forma de exibir e manter nossos produtos, crie uma coleção de variáveis e formar um novo tipo chamado Produto através do prefixo type, seguido do nome e o sufixo struct:

```go
type Produto struct {
    Nome       string
    Descricao  string
    Preco      float64
    Quantidade int
}
```

- Construa um slice de produtos, com todos os produtos que queremos exibir na nossa página index.html:

```go
func index(w http.ResponseWriter, r *http.Request) {
    produtos := []Produto{
        {Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
        {"Tenis", "Confortável", 89, 3},
        {"Fone", "Muito bom", 59, 2},
        {"Produto novo", "Muito legal", 1.99, 1},
    }

    temp.ExecuteTemplate(w, "Index", produtos)
}
```

- Para finalizar, altera o arquivo index.html para exibir todos os produtos de forma dinâmica:

```html
<tbody>
  {{range .}}
  <tr>
      <td>{{.Nome}}</td>
      <td>{{.Descricao}}</td>
      <td>{{.Preco}}</td>
      <td>{{.Quantidade}}</td>
  </tr>
  {{end}}
```

O gabarito deste exercício é o passo a passo demonstrado no vídeo. Tenha certeza de que tudo está certo antes de continuar.

### Um pouco mais sobre struct

Inicialmente, em nosso projeto, adicionamos os produtos direto no Html, criando uma nova tr e quatro td para representar, respectivamente: o nome, a descrição, o preço e a quantidade do produto.

Em seguida, criamos uma struct de Produto e tornamos nosso Html dinâmico, instanciando os produtos no arquivo main.go.

Sabendo disso, analise as afirmações abaixo e marque as verdadeiras.

- **Qualquer valor armazenado em uma struct, pode ser alterado.**
  - *Certo! As structs são tipos mutáveis, sendo assim, qualquer valor pode ser alterado. Para verificar um exemplo, você pode clicar no link a seguir para acessar The Go Playland.*
- **Uma struct é simplesmente uma coleção variáveis que formam um novo tipo.**
  - *Certo! No nosso exemplo, a struct de Produto possui 4 variáveis: nome e descrição do tipo string, preço do tipo float64 e quantidade do tipo inteiro.*
- **Podemos criar a instância de uma struct utilizando as chaves {} e, dentro delas, os valor dos campos na ordem que foram criados.**
  - *Certo! Diferente de outras linguagens de programação que utilizam os ( ) (parênteses), em Go, utilizamos { } (chaves).*
- Uma struct é simplesmente uma coleção variáveis do mesmo tipo que formam um novo tipo.

### O que aprendemos?

**Nessa aula:**

- Criamos o nosso projeto no workspace correto, dentro do GOPATH (dentro da pasta src, github.com, seguido do nome de usuário do Github);
- Aprendemos como subir um servidor através da função http.ListenAndServe(), exibindo uma tabela com nossos produtos;
- Em seguida criamos uma struct de Produto, onde instanciamos alguns produtos e exibimos de forma dinâmica em nossa index.html.

**Projeto desenvolvido nesta aula**

[Neste link](), você fará o download do projeto feito até esta aula.

Caso queira visualizar o código desenvolvido até aqui, [clique neste link](https://github.com/alura-cursos/web_com_golang).

**Na próxima aula**

Vamos instalar o Postgres para armazenar nossos produtos de forma segura, conectar nosso projeto e exibir os produtos que estão cadastrados no banco de dados!

## Conectando com banco de dados (33m)

### Instalando o Postgres


### Conectando com o banco


### Exibindo dados do banco


### Faça como eu fiz na aula


### Imports


### O que aprendemos?


## Refatoração e página de novos produtos (47m)

### Modularizando o código


### Modularizando um Pouco Mais


### Página de novos produtos


### Formulário de novos produtos


### Buscando dados da página


### Salvando produto no banco


### Faça como eu fiz na aula


### Dividindo o código


### O que aprendemos?


## Deletando produtos e partials (25m)

### Deletando produtos


### Verificando com Js


### Código duplicado e partials


### Faça como eu fiz na aula


### Quando vamos deletar...


### O que aprendemos?


## Atualizando e editando produtos (36m)

### Tela para atualizar produto


### Exibindo dados do produto


### Atualizando o produto


### Faça como eu fiz na aula


### Routes e routes.go


### O que aprendemos?


### Conclusão


### Parabéns



## Congratulations!

You have completed your course. Share your success on social media or email.

That's all folks!!!
___
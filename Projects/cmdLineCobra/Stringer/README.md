# Let's build a CLI in Go with Cobra

## Reference

> [Article: Let's build a CLI in Go with Cobra](https://www.thorsten-hans.com/lets-build-a-cli-in-go-with-cobra/)

## What is [Cobra](https://github.com/spf13/cobra)

Cobra is a library for creating powerful modern CLI applications.

Cobra is used in many Go projects such as [Kubernetes](https://kubernetes.io/), [Hugo](https://gohugo.io/), and [GitHub CLI](https://github.com/cli/cli) to name a few. [This list](https://github.com/spf13/cobra/blob/main/site/content/projects_using_cobra.md) contains a more extensive list of projects using Cobra.

## The hands-on

- Create a new CLI from scratch with Go and Cobra

```sh
# at the project folder
cd cmdlnCobraStringer

# open project in VSCode
code .

# initialize the project
go mod init cmdlnCobraStringer

# install cobra-cli package
go install github.com/spf13/cobra-cli@latest

# initialize a Cobra CLI project
... /go/bin/cobra-cli init

# create the repository structure base
mkdir -p pkg
mkdir -p bin

# edit and insert info about command
## edit cmd/root.go

# try and compile
go run main.go --help

go build -o ./bin/stringer
./bin/stringer --help

# adding the business-logic
mkdir -p pkg/geral
mkdir -p pkg/reverse
mkdir -p pkg/inspect
## edit pkg/geral/stringer.go
## edit pkg/inspect/inspect.go
## edit pkg/reverse/reverse.go

# adding commands
... /go/bin/cobra-cli add inspect
... /go/bin/cobra-cli add reverse

# edit and insert info about commands inserted
## edit cmd/inspect.go
## edit cmd/reverse.go


# add command to our Cobra CLI
... /go/bin/cobra-cli add stringer
# mkdir -p cmd/stringer
# touch cmd/stringer/root.go

# create the repository structure for this command
mkdir -p pkg/stringer
touch pkg/stringer/stringer.go

# add Cobra as a dependency
go get -u github.com/spf13/cobra@latest

```


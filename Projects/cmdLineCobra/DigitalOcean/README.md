# How To Use the Cobra Package in Go

> `https://www.digitalocean.com/community/tutorials/how-to-use-the-cobra-package-in-go`

## Step 1 - Setting Up Your Development Environment

The go mod command initializes a new project in the specified working directory. The command also creates a go.mod file in the directory for managing your project’s dependencies.

```sh
go mod init cmdlnCobraDOcean
```

## Step 2 - Getting Started the Cobra Package

You can run this command in the terminal of your project’s working directory to install the Cobra-cli package.

```sh
go install github.com/spf13/cobra-cli@latest
```

The command installs the Cobra-cli package as a CLI executable.

After installing the package, you can initialize a Cobra CLI project with the init command of the command line tool.

```sh
/home/tarso/go/bin/cobra-cli init
```

You must install your CLI package with the install command to access the changes and interact with your CLI app.

```sh
go install
```

After installing your CLI project application, you can run the tool with commands to test the tool and its functionalities.

```sh
/home/tarso/go/bin/cmdlnCobraDOcean
```

## Step 3 - Retrieving the Time in a Timezone With the Time Package

Create a time.go file with this command in the cmd directory.

```sh
cd cmd && touch time.go
```



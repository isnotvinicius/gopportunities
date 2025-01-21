# Anotacoes do projeto

## Iniciando o projeto

Utilizei o comando `go mod init <nome-do-modulo>` para iniciar o projeto com um arquivo go.mod. Este arquivo salva a versao sendo utilizada do Go e tambem os packages externos que serao utilizados no mesmo.

Alem de executar o comando de modulo, criei um arquivo `main.go` e nomear o `package main` para poder compilar o projeto.

## Importando bibliotecas externas (Gin Gonic)

Para importar uma biblioteca externa em Go, existem varias maneiras, no projeto utilizei a forma mais simples e facil.

Adicionei o import da biblioteca no meu arquivo `main.go` e executei o comando `go mod tidy`. Este comando limpa o arquivo `go.mod`, removendo imports nao utilizados e adicionando imports que nao foram adicionados ainda

```
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    
}
```

Ao executar o comando `go mod tidy`, ele gera um arquivo `go.sum` que seria o equivalente ao `composer.lock` do laravel ou o `package-lock.json` do javascript, servindo como um freeze de versoes das bibliotecas adicionadas no projeto.


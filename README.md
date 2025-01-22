# Project annotations

This project uses the Go language and it's library Gin Gonic to create an API using SQLite and Swagger.

Project based on Arthur 404 dev [video](https://youtu.be/wyEYpX5U4Vg?si=EX_H8vZqy07PHPoz)

## Initializing the project

I've used the command `go mod init <nome-do-modulo>` to initialize the project with a `go.mod` file. This file saves the Go version being used and also external packages that we may import later on (such as Gin Gonic).

I also created the `main.go` file to initialize the main package so I could compile the project.

## Importing external packages (Gin Gonic)

Para importar uma biblioteca externa em Go, existem varias maneiras, no projeto utilizei a forma mais simples e facil.

There are several ways to import an external package, I used the most simple and easier way.

I added the import of the package on my `main.go` file and runned the command `go mod tidy`. This command cleans the `go.mod` file, removing all unused imports and adding imports that are being used but not declared on the `go.mod` file.

```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    
}

$ go mod tidy
```

After running `go mod tidy`, a `go.sum` will be generated, which is the equivalent of `composer.lock` from Laravel or `package-lock.json` from JavaScript.

## New module

It's not that good to create everything inside the main package of the application, so I created a new file (also a new package) so I can manage things properly, like routing.

As a code standard, every new directory becomes a new package. So I created a new directory called `router` and also a package router `router.go`, inside of it I created the `func Initialize()` to initialize the package and called this function inside my main package.

```go
// Router

package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}

// Main
package main

import "github.com/isnotvinicius/gopportunities/router"

func main() {
	router.Initialize()
}
```

P.S: To export something from one file to another, like the initialize function, the thing you want to export NEEDS to start with a capital letter instead of lowercase. If the `Initialize()` function was declared with a lowercase `i`, the function could never be accessed on the main package, because `Initialize() != initialize()`.

## Splitting the routes into a new file

When creating multiple routes, the router package would be a mess because each route has it's own logic and they will all be together in the same function. So I created a new file inside the router package called `routes.go`. This file will be the responsible for managing the routes properly.

Everything that is declared inside this new file can be accessed on the `router.go` file, just call it as it is.

```go
// routes.go
package router

func exampleFunction() {
	print("Hello, World!")
}

// router.go
package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
    exampleFunction() // This can be called without import because it's on the same package, just a different file.

	router := gin.Default()

	router.Run(":8080")
}

```

Inside the `routes.go`, I created a function called `initializeRoutes()` that receives the router and returns a group of routes called v1. 

```go
package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) { // Receives the router as param and it's type (a pointer to gin.Engine)
	v1 := router.Group("/api/v1") // Defines a group of routes with the prefix /api/v1
	{
		v1.GET("/opening", func(ctx *gin.Context) { // Defines a route with the method GET and the path /api/v1/opening (group + route)
			ctx.JSON(http.StatusOK, gin.H{ // Return a JSON response with the status code 200 and a message using the gin.H helper
				"message": "GET Opening",
			})
		})
	}
}
```
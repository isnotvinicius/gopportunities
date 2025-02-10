# Project notes

This project uses the Go language and it's library Gin Gonic to create an API using SQLite and Swagger.

Project based on Arthur 404 dev [video](https://youtu.be/wyEYpX5U4Vg?si=EX_H8vZqy07PHPoz)

## Initializing the project

I've used the command `go mod init <nome-do-modulo>` to initialize the project with a `go.mod` file. This file saves the Go version being used and also external packages that we may import later on (such as Gin Gonic).

I also created the `main.go` file to initialize the main package so I could compile the project.

## Importing external packages (Gin Gonic)

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

## Handlers

In order to keep the routes more easily to read and indentify them, I separeted the anonymous func of the route into a handler package, which will deal with the logic for the route, separating things properly.

Inside the handler package I created multiple files, one for each of the routes declared before. After that, instead of calling an anonymous function I called the function inside the handler file that corresponds to the current route. This makes the `routes.go` file to be more readable and easy to indentify the routes.

## Database

This is a simple CRUD, so I need a Database. In this project I'm using `GORM` as the ORM and `SQLite` as a database.

First, I installed GORM following the documentation provided by it's website.

The I created the package `schemas`, where I will create all the schemas that will be used on the api. This project only has one, but I could add more by creating new files. Different from the other packages, this one won't have a `schemas.go` file, instead each schema will have a file.

```
.
└── schemas
    └── opening.go
    └── other_schema.go
    └── another_schema.go
```

The schema file will look like this:

```go
package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

```

This struct created will act as a model to the database table. Same thing with other ORM's out there.

## The config module

In order to define some configurations of the project, I created a config package. Define the instance of the database is one of the many things this module can have. I did it by declaring a var in the package scope that receives the gorm and then a `Init()` function to initialize the config. The function for now will be empty..

```go
package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func Init() error {
	return nil
}
```

Before doing anything on the api, the config needs to be initialized, I did this by adding the `config.Init()` function on my main package.

```go
package main

import (
	"fmt"

	"github.com/isnotvinicius/gopportunities/config"
	"github.com/isnotvinicius/gopportunities/router"
)

func main() {
	// Initialize configs
	err := config.Init()

	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize the router package
	router.Initialize()
}
```

If error has any value besides `nil`, I print an error and return the `main func`.

## Logger

It's good to return an error message or even `panic()` if the config couldn't be initialized, but in this case I will create a logger in order to save this errors. For this, I created the logger file inside the config package. Inside the `logger.go`, I created a struct called Logger and defined all log types that the API should have.

Then I created a function that receives the log prefix and return the log using a writer, the prefix and a few flags to enhace the log message.

Finally I created a few functions to return the logging message with and without formating.

```go
package config

import (
	"io"
	"log"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Non-formatted loggers
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Formatted loggers

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
```

P.S: the parenthesis before the func name is to declare a receiver to the func we are creating. In this case I'm saying that each of the func's I created will be avaiable in the Logger struct because it's pointer is declared as the receiver of the given func. `func (variable *Struct) FuncName(params)`

Now to implement the logger, we declare the logger on the main package and call a function inside config that initialize our Logger.

```go
// main.go
package main

import (
	"github.com/isnotvinicius/gopportunities/config"
	"github.com/isnotvinicius/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main package")

	// Initialize configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config init resulted in error: %v", err)
		return
	}

	// Initialize the router package
	router.Initialize()
}

// config.go
package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
```

## Creating and connecting to the database

Inside the config package, I created a `sqlite.go` file to initialize the connection and connect to it. Just followed the [documentation](https://gorm.io/docs/) guide to connect to it and used the logger created before to display some errors (when they occur). I also added some validation for when the SQLite file does not exist, it will create automatically for me, this ensures that the api will never run without the database file being created.

```go
package config

import (
	"os"

	"github.com/isnotvinicius/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	// Check if the database file exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating it...")

		// Create the database directory
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		// Create the database file
		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		// Close the file to avoid errors
		file.Close()
	}

	// Create database and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Migrate the schema created before
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("Failed to migrate schema: %v", err)
		return nil, err
	}

	return db, nil
}


```

After configuring the database on the `sqlite.go` file, I added it's initialization on the `config.go` file to start the database as soon as the api starts running (because the first thing the main package does is initialize the config package).

```go
package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize the Database
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
```

## JSON Response

To facilitate data manipulation, I created a struct inside my `schemas/opening.go` called `OpeningResponse` that maps all the model fields to a json correspondant, so when we retrieve data they come as a json.

```go
type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"` // When empyt, omits it from the response
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}
```

## Finishing the Handlers

### Create Opening

Finally, after a lot of configuration and creating things to help, I started finishing the handlers and actually creating the CRUD.

Before creating the actual logic in the handlers, I added the logger and database inits on the `handler.go` file so they can be accessed on any of the handlers created. Then I initialized the handler inside the `routes.go` file.

```go
// handler/handler.go
package handler

import (
	"github.com/isnotvinicius/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}

// router/routes.go
func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.Init()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.PostOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningHandler)
	}
}
```

I created a `handler/request.go` file to manage the requests of the API. This file will have a type struct for each of the handlers (CRUD methods) to map which fields the request should receive. This would be the "equivalent" of creating a request file in Laravel and adding the field and their types there. Go will ignore fields that are sent on the request but are NOT declared on the struct being used.

Besides creating the struct, I added a validation function that uses a helper to return a message when a field is empty on the request.

```go
package handler

import "fmt"

func errParamIsRequired(name string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// createOpening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}

	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}

	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if r.Salary == 0 {
		return errParamIsRequired("salary", "int")
	}

	return nil
}
```

Then on the `handler/createOpening.go` file, I just create a request variable and initialize it with the struct I just created. Using the request variable to store data will ensure that ONLY the fields declared on the struct will be used, any additional fields will be simply ignored. I binded the json of the request to the request variable and used the `Validate()` to check if data sent is correct before storing it in the database.

```go
package handler

import (
	"github.com/gin-gonic/gin"
)

func PostOpeningHandler(ctx *gin.Context) {
	// Initialize request with data from the body mapped by the struct
	request := CreateOpeningRequest{}

	// bind the JSON body to the request variable using gin context
	ctx.BindJSON(&request)

	// Validate the request
	// Declares err and do request.Validate(), if err is not nil I log the error. Common syntax in Go.
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		return
	}
}
```

To manage the response and sanitize it, I created a `handler/response.go` file. Inside it, a `sendError()` function was declared and it uses the gin context to return the error response, just like the documentation of Gin Gonic says but in a way that all handlers can use it. I also created a `sendSucess()` function to return the response in a success case with the headers and the body content.

```go
package handler

import (
	"github.com/gin-gonic/gin"

	"fmt"

	"net/http"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": message,
		"status":  code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}
```

Then on the `createOpening.go` file I do the logic to receive the request, validate all data and insert it on the database using all the functions created previously.
```go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/isnotvinicius/gopportunities/schemas"
)

func PostOpeningHandler(ctx *gin.Context) {
	// Initialize request with data from the body mapped by the struct
	request := CreateOpeningRequest{}

	// bind the JSON body to the request variable using gin context
	ctx.BindJSON(&request)

	// Validate the request
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Maps the request values to a variable to avoid any data we don't want
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error while creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
```

### Delete Opening

The rest of the handlers were easy to implement, just needed to follow the logic behind the create and adapt according to the operation being executed at the time. For delete I just needed to retrieve the id on the queryParam, find the opening and then delete it, always checking for errors and logging them.

```go
package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isnotvinicius/gopportunities/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	// Get the id from the query parameters and validates it
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// Find the opening and send an error when not found
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	// Deletes the opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error while deleting opening with id %s on database", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}
```

### List Openings

Just did a find method and logged an error when none opening is found.

```go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isnotvinicius/gopportunities/schemas"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error while listing openings on database")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}

```
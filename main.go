package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// When creating an API you start with the API endpoint. The API endpoint is what the client interacts with.
// endpoint 1 : /quotes
// GET- Get the list of ALL quotes in the form of JSON.
// POST - Add a quote in the form of JSON
// endpoint 2 : /quotes/:id
// GET - get a specific quote specified by the ID
type napoleonQuote struct {
	ID     string `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

// As of right now the data is stored in memory, though usually this is stored in a database.
// Because its stored in memory, each time the server is closed, data cannot be retrieved. Only when server is running can the data be retrieved.
var quotes = []napoleonQuote{

	{ID: "1", Quote: "My enemies are many, my equals are none.", Author: "Napoleon"},
	{ID: "2", Quote: "A general's most important talent is to know the mind of the soldier and gain his confidence... \n He is not a machine that must be made to move, he is reasonable being who needs leadership.", Author: "Napoleon"},
	{ID: "3", Quote: "His small size and and punny face did not put him in their favour... but as soon as he put on his general's hat, he seemed to grow by two feet.", Author: "Messena"},
	{ID: "4", Quote: "I no longer regarded myself as a simple general, but as a man called upon to decide the fates of peoples", Author: "Napoleon"},
}

// Now a handler needs to be created.
// When the client makes a request at GET /quotes, all the quotes must be returned as JSON.
// To do this you need 1.Logic to prepare the response 2. Code to map the request path to your logic.

func getQuotes(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, quotes)
}

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open("DATABASE_URL"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func main() {
	db, err := gorm.Open(postgres.Open("DATABASE_URL"), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "You cannot connect to the database", err)
	}

	database, _ := db.DB()
	err = database.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "The ping did not work my boy", err)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATBASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "My brother, you are unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: &V\n", err)
		os.Exit(1)
	}
	fmt.Println(greeting)

	router := gin.Default()
	router.GET("/quotes", getQuotes)
	router.Run("localhost:9000")
}

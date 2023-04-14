package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type napoleonQuote struct {
	ID     string `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var quotes = []napoleonQuote{

	{ID: "1", Quote: "My enemies are many, my equals are none.", Author: "Napoleon"},
	{ID: "2", Quote: "A general's most important talent is to know the mind of the soldier and gain his confidence... \n He is not a machine that must be made to move, he is reasonable being who needs leadership.", Author: "Napoleon"},
	{ID: "3", Quote: "His small size and and punny face did not put him in their favour... but as soon as he put on his general's hat, he seemed to grow by two feet.", Author: "Messena"},
	{ID: "4", Quote: "I no longer regarded myself as a simple general, but as a man called upon to decide the fates of peoples", Author: "Napoleon"},
}

func getQuotes(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, quotes)
}

func main() {
	router := gin.Default()
	router.GET("/quotes", getQuotes)
	router.Run("localhost:9000")
}

package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "math/rand"
  "time"
)

// type declaration
type joke struct {
	id        int
	beginning string
	punchline string
}

// the jokes list goes here
var jokes = []joke{
	{
		id:        1,
		beginning: "Why are snails slow?",
		punchline: "Because they’re carrying a house on their back.",
	},
	{
		id:        2,
		beginning: "What does a storm cloud wear under his raincoat?",  
		punchline: "Thunderwear.",
	},
	{
		id:        3,
		beginning: "How does the ocean say hi?",
		punchline: "It waves!",
	},
}

// main function
func main() {
  rand.Seed(time.Now().Unix())
  n := rand.Int() % len(jokes)
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
	 c.JSON(http.StatusOK, gin.H {
     "beginning": jokes[n].beginning,
     "punchline": jokes[n].punchline,
     "id": jokes[n].id,
   })
  })
  r.Run()
}
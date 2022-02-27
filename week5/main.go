package main

import (
	"geek/week5/route"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	route.Route(engine)

	log.Fatal(engine.Run())
}

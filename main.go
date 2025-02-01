package main

import (
	"MATHB/app/midwares"
	"MATHB/app/utils/env"
	"MATHB/config/router"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	router.Init(r)
	err := r.Run(fmt.Sprintf(":%d", env.Port))
	if err != nil {
		log.Fatal("ServerStartFailed", err)
	}
}

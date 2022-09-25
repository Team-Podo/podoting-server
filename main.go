package main

import (
	"fmt"
	"github.com/Team-Podo/podoting-server/database/migration"
	"github.com/Team-Podo/podoting-server/routes"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/Team-Podo/podoting-server/utils/aws"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	utils.SetEnv()
}

func main() {
	go migration.Init()

	r := gin.Default()
	routes.Routes(r)

	aws.Init()

	fmt.Printf("Podoting running on %s mode \n", os.Getenv("APP_ENV"))

	err := r.Run(":80")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

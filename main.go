package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kwanok/podonine/database/migration"
	"github.com/kwanok/podonine/routes"
	"github.com/kwanok/podonine/utils"
	"os"
)

func init() {
	utils.SetEnv()
}

func main() {
	migration.Init()

	r := gin.Default()
	routes.Routes(r)

	fmt.Printf("Podoting running on %s mode \n", os.Getenv("APP_ENV"))

	err := r.Run(":80")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

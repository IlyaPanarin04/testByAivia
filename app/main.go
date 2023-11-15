package main

import (
	"github.com/gin-gonic/gin"
	"testByAivia/config"
	"testByAivia/internal/service"
)

func main() {
	serv := gin.Default()

	newConfig := config.NewConfig()

	/*pgConn := postgres.Open(newConfig.Datastore.Postgres)

	db, err := repository.NewConnection(&pgConn)
	if err != nil {
		log.Fatalf(`Error in connection to db: %v`, err)
	}*/

	s := service.NewService( /*db,*/ serv)

	s.StartServer(newConfig.Port)
}

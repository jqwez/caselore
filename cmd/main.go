package cmd

import (
	"fmt"
	"github.com/jqwez/caselore/database"
	"github.com/jqwez/caselore/server"
)

func Main() {
	// config.InitializeConfig()
	repository := database.NewRepository().MustSetFromEnv()

	s := server.NewServer(repository)
	fmt.Println("Running Server on port 8080")
	s.Run()

}

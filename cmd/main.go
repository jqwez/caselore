package cmd

import (
	"fmt"
	"github.com/jqwez/caselore/repo"
	"github.com/jqwez/caselore/server"
)

func Main() {
	// config.InitializeConfig()
	repository := repo.GetRepo()

	s := server.NewServer(repository)
	fmt.Println("Running Server on port 8080")
	s.Run()

}

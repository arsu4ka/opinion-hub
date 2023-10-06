package main

import (
	"github.com/aru4ka/opinion-hub/internal/app/configs"
	"github.com/aru4ka/opinion-hub/internal/app/server"
)

func main() {
	config, err := configs.NewPostgresServerConfig(true)
	if err != nil {
		panic(err)
	}

	s := server.New(config)
	panic(s.Start())
}

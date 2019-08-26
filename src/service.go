package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml"
)

type Service struct {
	Router *gin.Engine
	Config *toml.Tree
}

func (s *Service) init() {
	config, err := toml.LoadFile("config/pica.toml")
	if err != nil{
		panic(err)
	}
	s.Config = config

	s.initRouter()
}
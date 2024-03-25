package main

import "github.com/takwot/tech-strelka.git/pkg/database"

type Config struct {
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`
	Postgres *database.Config `yaml:"postgres"`
}

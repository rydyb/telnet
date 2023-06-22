package main

import (
	"fmt"
	"time"

	"github.com/alecthomas/kong"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/rydyb/telnet"
)

var cli struct {
	Host    string `name:"host"`
	Port    int    `name:"port"`
	Command string `arg:"command"`
	Debug   bool   `name:"debug" default:"false"`
	Timeout int    `name:"timeout" default:"1"`
}

func main() {
	kong.Parse(&cli)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cli.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	client := &telnet.Client{
		Timeout: time.Duration(cli.Timeout) * time.Second,
		Address: fmt.Sprintf("%s:%d", cli.Host, cli.Port),
	}
	if err := client.Open(); err != nil {
		log.Fatal().Err(err).Msg("failed to open client connection")
	}
	defer client.Close()

	out, err := client.Exec(cli.Command)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to acquire measurement names")
	}
	fmt.Println(out)
}

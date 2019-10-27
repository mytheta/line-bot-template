package main

import (
	"flag"
	"fmt"

	"github.com/mytheta/line-bot-template/pkg/server"

	"github.com/mytheta/line-bot-template/conf"
)

var state = flag.String("s", "local", "local/prd")

func main() {
	flag.Parse()
	if err := conf.Setup(fmt.Sprintf("conf/env/%s.toml", *state)); err != nil {
		fmt.Sprintf("%s", err)
		return
	}

	server.Init()
}

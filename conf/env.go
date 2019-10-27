package conf

import (
	"log"
	"runtime"

	"github.com/BurntSushi/toml"
)

var conf Config

type (
	// Config is all config
	Config struct {
		Bot BotConfig
	}

	// BotConfig is line bot config
	BotConfig struct {
		ChannelSecret string
		ChannelToken  string
	}
)

// Setup is env setup
func Setup(file string) error {
	runtime.GOMAXPROCS(runtime.NumCPU())
	_, err := toml.DecodeFile(file, &conf)
	if err != nil {
		// Error Handling
		log.Panic(err)
		return err
	}

	return nil
}

// GetBotConfig is get line bot config
func GetBotConfig() *BotConfig {
	return &conf.Bot
}

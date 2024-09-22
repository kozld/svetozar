package initializers

import "github.com/kozld/svetozar/config"

func InitializeConfig() *config.Config {
	return config.LoadConfig()
}

package initializers

import (
	"path/filepath"

	tdlib "github.com/zelenin/go-tdlib/client"

	"github.com/kozld/svetozar/config"
)

func InitializeTDLibClient(conf *config.Config) (*tdlib.Client, error) {
	authorizer := tdlib.ClientAuthorizer()
	go tdlib.CliInteractor(authorizer)

	authorizer.TdlibParameters <- &tdlib.SetTdlibParametersRequest{
		UseTestDc:           false,
		DatabaseDirectory:   filepath.Join(".tdlib", "database"),
		FilesDirectory:      filepath.Join(".tdlib", "files"),
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseMessageDatabase:  true,
		UseSecretChats:      false,
		ApiId:               conf.TgApiID,
		ApiHash:             conf.TgApiHash,
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
	}

	tdlib.SetLogVerbosityLevel(&tdlib.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 1,
	})

	client, err := tdlib.NewClient(authorizer)
	if err != nil {
		return nil, err
	}

	return client, err
}

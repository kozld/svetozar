package dependencies

import (
	"go.uber.org/zap"

	"github.com/zelenin/go-tdlib/client"

	"github.com/kozld/svetozar/config"
	repo "github.com/kozld/svetozar/internal/repositories/message"
	"github.com/kozld/svetozar/pkg/workerpool"
)

// Container is a DI container for application
type Container struct {
	TDLibClient *client.Client
	Workerpool  *workerpool.WorkerPool
	Repository  *repo.Repository
	Config      *config.Config
	Logger      *zap.Logger
}

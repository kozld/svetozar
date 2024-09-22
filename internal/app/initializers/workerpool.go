package initializers

import (
	"github.com/kozld/svetozar/internal/services/listener"
	"github.com/kozld/svetozar/internal/services/validator"
	"github.com/kozld/svetozar/pkg/workerpool"
)

func InitializeWorkerpool(
	t1 *listener.Service,
	t2 *validator.Service,
) *workerpool.WorkerPool {
	wp := workerpool.NewWorkerPool(3)
	wp.Start()

	wp.Submit(t1)
	wp.Submit(t2)

	return wp
}

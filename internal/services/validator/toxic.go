package validator

import (
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/kozld/svetozar/internal/models"
	repo "github.com/kozld/svetozar/internal/repositories/message"
	"github.com/kozld/svetozar/pkg/workerpool"
)

var (
	_ workerpool.Task = (*Service)(nil)
)

type (
	Service struct {
		repo   *repo.Repository
		logger *zap.Logger
	}
)

const (
	ScriptsDir     = "scripts"
	ScriptFileName = "rubert_toxic.py"
	ExecCommand    = "python3"

	ValidateDelay = 3 * time.Second
)

var (
	filePath = path.Join(ScriptsDir, ScriptFileName)
)

func New(
	repo *repo.Repository,
	logger *zap.Logger,
) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) Validate() error {
	messages, err := s.repo.ByStatus(models.MessageNew)
	if err != nil {
		return err
	}

	if len(messages) == 0 {
		return nil
	}

	for _, message := range messages {
		cmd := exec.Command(ExecCommand, filePath, message.Text)
		out, err := cmd.Output()
		if err != nil {
			continue
		}

		re := regexp.MustCompile(`[^0-9]`)
		out = re.ReplaceAll(out, []byte(""))

		status, _ := strconv.Atoi(string(out))
		message.Status = models.Status(status + 1)
	}

	return s.repo.Update(messages...)
}

func (s *Service) Run() {
	for {
		<-time.After(ValidateDelay)

		if err := s.Validate(); err != nil {
			continue
		}
	}
}

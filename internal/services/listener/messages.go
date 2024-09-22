package listener

import (
	"time"

	tdlib "github.com/zelenin/go-tdlib/client"
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
		client *tdlib.Client
		repo   *repo.Repository
		logger *zap.Logger
	}
)

const (
	FlushDelay = 10 * time.Second
)

func New(
	client *tdlib.Client,
	repo *repo.Repository,
	logger *zap.Logger,
) *Service {
	return &Service{
		client: client,
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) ListenMessages() {
	listener := s.client.GetListener()
	defer listener.Close()

	messages := make([]*models.Message, 0)

	for {
		select {
		case <-time.After(FlushDelay):
			if len(messages) == 0 {
				continue
			}

			s.repo.Update(messages...)
			messages = make([]*models.Message, 0)

		case update := <-listener.Updates:
			if update.GetClass() == tdlib.ClassUpdate {
				newMessage, ok := update.(*tdlib.UpdateNewMessage)
				if ok {
					messageText, ok := newMessage.Message.Content.(*tdlib.MessageText)
					if ok {
						messages = append(messages, &models.Message{
							ID:     newMessage.Message.Id,
							ChatID: newMessage.Message.ChatId,
							Text:   messageText.Text.Text,
							Status: models.MessageNew,
						})
					}
				}
			}
		}
	}
}

func (s *Service) Run() {
	s.ListenMessages()
}

package message

import "log/slog"

type MessageService struct {
	repo *MessageRepository

	logger *slog.Logger
}

func NewService(repo *MessageRepository, logger *slog.Logger) *MessageService {
	return &MessageService{
		repo:   repo,
		logger: logger,
	}
}

func (s MessageService) SaveMsg(msg string) error {
	if err := s.repo.SaveMsg(msg); err != nil {
		return err
	}

	//TODO: функция для отправки в очередь на обработку

	return nil
}

//TODO: что делать если сообщение попало в бд и апп крашулся, т.е. сообщение есть в бд, но его нет в очереди, при этом нужно избеать повторной обработки
// бд + (при перезапуске проверяет какие сообщения есть в очереди) кафка

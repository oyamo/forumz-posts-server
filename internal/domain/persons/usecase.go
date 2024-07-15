package persons

import (
	"context"
	"go.uber.org/zap"
)

type UseCase struct {
	personRepository PersonRepository
	logger           *zap.SugaredLogger
}

func (u *UseCase) Upsert(ctx context.Context, dto *Person) error {
	err := u.personRepository.Upsert(ctx, dto)
	if err != nil {
		u.logger.Errorw("error while upserting persons", "error", err)
		return err
	}

	return nil
}

func NewUseCase(personRepository PersonRepository, logger *zap.SugaredLogger) *UseCase {
	return &UseCase{
		personRepository: personRepository,
		logger:           logger,
	}
}

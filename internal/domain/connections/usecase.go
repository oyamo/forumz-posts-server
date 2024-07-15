package connections

import (
	"context"
)

type UseCase struct {
	connectionRepository Repository
}

func (uc *UseCase) Connect(ctx context.Context, dto *CreateConnectionDTO) error {
	connection := &Connection{
		UserId:      dto.UserId,
		ConnectedTo: dto.ConnectionTo,
	}
	err := uc.connectionRepository.Save(ctx, connection)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) Disconnect(ctx context.Context, dto *CreateConnectionDTO) error {
	err := uc.connectionRepository.Delete(ctx, dto.UserId, dto.ConnectionTo)
	if err != nil {
		return err
	}

	return nil
}

func NewUseCase(connectionRepository Repository) *UseCase {
	return &UseCase{connectionRepository: connectionRepository}
}

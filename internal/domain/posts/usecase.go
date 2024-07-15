package posts

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UseCase struct {
	postsRepo Repository
	logger    *zap.SugaredLogger
}

func (u UseCase) Create(ctx context.Context, dto CreatePostDTO) (*Post, error) {
	postId, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	post := Post{
		Id:       postId,
		PersonId: dto.PersonId,
		Content:  dto.Content,
	}

	err = u.postsRepo.Create(ctx, &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (u UseCase) Find(ctx context.Context, id uuid.UUID) (*Post, error) {
	post, err := u.postsRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (u UseCase) Delete(ctx context.Context, id uuid.UUID, initiatorId uuid.UUID) error {
	err := u.postsRepo.Delete(ctx, id, initiatorId)
	if err != nil {
		return err
	}

	return nil
}

func NewUseCase(postsRepo Repository, logger *zap.SugaredLogger) *UseCase {
	return &UseCase{
		postsRepo: postsRepo,
		logger:    logger,
	}
}

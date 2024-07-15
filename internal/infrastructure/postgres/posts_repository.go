package postgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"posts-server/internal/domain/posts"
)

type postgresPostsRepository struct {
	db *sql.DB
}

func (p postgresPostsRepository) Create(ctx context.Context, post *posts.Post) error {
	stmt, err := p.db.Prepare(`insert into post (id, person_id, content) values ($1, $2, $3)`)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(ctx,
		post.Id,
		post.PersonId,
		post.Content,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p postgresPostsRepository) Get(ctx context.Context, id uuid.UUID) (*posts.Post, error) {
	stmt, err := p.db.Prepare(`select id, person_id, content, datetime_created, last_modified from post where id = $1`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var post posts.Post
	err = stmt.QueryRowContext(ctx, id).Scan(
		&post.Id,
		&post.PersonId,
		&post.Content,
		&post.DatetimeCreated,
		&post.LastModified)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p postgresPostsRepository) Delete(ctx context.Context, id uuid.UUID, initiatorId uuid.UUID) error {
	stmt, err := p.db.Prepare(`delete from post where id = $1 and person_id = $2`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id, initiatorId)
	if err != nil {
		return err
	}
	return nil
}

func NewPostsRepository(db *sql.DB) posts.Repository {
	return &postgresPostsRepository{
		db: db,
	}
}

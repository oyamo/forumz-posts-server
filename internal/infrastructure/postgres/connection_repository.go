package postgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"posts-server/internal/domain/connections"
)

type psqlConnectionRepository struct {
	db *sql.DB
}

func (p psqlConnectionRepository) Save(ctx context.Context, connection *connections.Connection) error {
	stmt, err := p.db.Prepare(`insert into connection(user_id, connected_to)
		values ($1, $2) on conflict (user_id, connected_to) do nothing`)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.ExecContext(
		ctx,
		connection.UserId,
		connection.ConnectedTo,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p psqlConnectionRepository) Delete(ctx context.Context, id, connectedTo uuid.UUID) error {
	stmt, err := p.db.Prepare(`delete from connection where user_id = $1 and connected_to = $2`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id, connectedTo)
	if err != nil {
		return err
	}

	return nil
}

func NewConnectionRepository(db *sql.DB) connections.Repository {
	return &psqlConnectionRepository{
		db: db,
	}
}

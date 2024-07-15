package postgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"posts-server/internal/domain/persons"
)

type psqlPersonRespository struct {
	db *sql.DB
}

func (repo psqlPersonRespository) Upsert(ctx context.Context, person *persons.Person) error {
	stmt, err := repo.db.Prepare(`insert into person (id,first_name, email_address, username) 
    values ($1, $2, $3, $4) 
    on conflict(id) do update set first_name = $2, email_address = $3, username = $4`)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		person.Id,
		person.FirstName,
		person.EmailAddress,
		person.Username,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo psqlPersonRespository) Find(ctx context.Context, u uuid.UUID) (*persons.Person, error) {
	stmt, err := repo.db.Prepare(`select id, first_name, last_name, email_address, password_hash, username, status, dob, datetime_created, last_modified from person where id = $1`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var person persons.Person

	row := stmt.QueryRowContext(ctx, u)
	err = row.Scan(
		&person.Id,
		&person.FirstName,
		&person.Status,
		&person.DatetimeCreated,
		&person.LastModified,
	)

	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (repo psqlPersonRespository) FindByUsername(ctx context.Context, username string) (*persons.Person, error) {
	stmt, err := repo.db.Prepare(`select id, first_name, last_name, email_address, password_hash, username, status, dob, datetime_created, last_modified from person where username = $1 or email_address = $1 limit 1`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var person persons.Person
	row := stmt.QueryRowContext(ctx, username)
	err = row.Scan(
		&person.Id,
		&person.FirstName,
		&person.Status,
		&person.DatetimeCreated,
		&person.LastModified,
	)

	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (repo psqlPersonRespository) Exists(ctx context.Context, u uuid.UUID) (bool, error) {
	stmt, err := repo.db.Prepare(`select exists(select 1 from person where id = $1)`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	var exists bool
	err = stmt.QueryRowContext(ctx, u).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (repo psqlPersonRespository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	stmt, err := repo.db.Prepare(`select exists(select 1 from person where email_address = $1)`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	var exists bool
	err = stmt.QueryRowContext(ctx, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (repo psqlPersonRespository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	stmt, err := repo.db.Prepare(`select exists(select 1 from person where username = $1)`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var exists bool
	err = stmt.QueryRowContext(ctx, username).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func NewPersonRepository(db *sql.DB) persons.PersonRepository {
	return &psqlPersonRespository{
		db: db,
	}
}

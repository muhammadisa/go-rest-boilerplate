package repository

import (
	"errors"
	"time"

	"github.com/gocraft/dbr/v2"
	"github.com/muhammadisa/go-rest-boilerplate/api/apps/foobar"
	"github.com/muhammadisa/go-rest-boilerplate/api/models"
	uuid "github.com/satori/go.uuid"
)

type foobarRepository struct {
	Sess *dbr.Session
}

// NewFoobarRepository function
func NewFoobarRepository(sess *dbr.Session) foobar.Repository {
	return &foobarRepository{
		Sess: sess,
	}
}

func (foobarRepositories *foobarRepository) Fetch() (
	[]models.Foobar,
	error,
) {
	var err error
	var foobars []models.Foobar

	_, err = foobarRepositories.Sess.Select("*").
		From("foobars").
		Load(&foobars)
	if err != nil {
		return nil, err
	}
	return foobars, nil
}

func (foobarRepositories *foobarRepository) Create(
	foobar *models.Foobar,
) error {
	var err error

	_, err = foobarRepositories.Sess.InsertInto("foobars").
		Columns("id", "foobar_content", "created_at").
		Record(foobar).
		Exec()
	if err != nil {
		return err
	}
	return nil
}

func (foobarRepositories *foobarRepository) Update(
	foobar *models.Foobar,
) error {
	var err error

	_, err = foobarRepositories.Sess.Update("foobars").
		Where("id = ?", foobar.ID.String()).
		SetMap(map[string]interface{}{
			"foobar_content": foobar.FoobarContent,
			"updated_at":     time.Now(),
		}).
		Exec()
	if err != nil {
		return err
	}
	return nil
}

func (foobarRepositories *foobarRepository) Delete(id uuid.UUID) error {
	var err error

	result, err := foobarRepositories.Sess.DeleteFrom("foobars").
		Where("id = ?", id.String()).
		Exec()
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Foobar not found")
	}
	return nil
}

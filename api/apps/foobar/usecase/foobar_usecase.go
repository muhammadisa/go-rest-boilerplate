package usecase

import (
	"time"

	"github.com/muhammadisa/go-rest-boilerplate/api/apps/foobar"
	"github.com/muhammadisa/go-rest-boilerplate/api/models"
	uuid "github.com/satori/go.uuid"
)

type foobarUsecase struct {
	foobarRepository foobar.Repository
}

// NewFoobarUsecase function
func NewFoobarUsecase(foobarRepository foobar.Repository) foobar.Usecase {
	return &foobarUsecase{
		foobarRepository: foobarRepository,
	}
}

func (foobarUsecases foobarUsecase) Fetch() ([]models.Foobar, error) {
	results, err := foobarUsecases.foobarRepository.Fetch()
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (foobarUsecases foobarUsecase) Create(foobar *models.Foobar) error {
	foobar.ID = uuid.NewV4()
	foobar.CreatedAt = time.Now()
	err := foobarUsecases.foobarRepository.Create(foobar)
	if err != nil {
		return err
	}
	return nil
}

func (foobarUsecases foobarUsecase) Update(foobar *models.Foobar) error {
	foobar.UpdatedAt = time.Now()
	err := foobarUsecases.foobarRepository.Update(foobar)
	if err != nil {
		return err
	}
	return nil
}

func (foobarUsecases foobarUsecase) Delete(id uuid.UUID) error {
	return foobarUsecases.foobarRepository.Delete(id)
}

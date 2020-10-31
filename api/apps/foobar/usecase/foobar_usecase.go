package usecase

import (
	"fmt"
	"sync"
	"time"

	"github.com/muhammadisa/goredisku"

	"github.com/muhammadisa/go-rest-boilerplate/api/apps/foobar"
	"github.com/muhammadisa/go-rest-boilerplate/api/models"
	uuid "github.com/satori/go.uuid"
)

type foobarUsecase struct {
	foobarRepository foobar.Repository
	cacheCommand     *goredisku.GoRedisKu
}

// NewFoobarUsecase function
func NewFoobarUsecase(
	foobarRepository foobar.Repository,
	cacheCommand *goredisku.GoRedisKu,
) foobar.Usecase {
	return &foobarUsecase{
		foobarRepository: foobarRepository,
		cacheCommand:     cacheCommand,
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
	fatalError := make(chan error, 1)

	foobar.ID = uuid.NewV4()
	foobar.CreatedAt = time.Now()
	key := fmt.Sprintf("%s:%s", foobar.ID.String(), "foobar")
	err := foobarUsecases.cacheCommand.WT(key, &foobar, func() {
		fmt.Println("Write Through db event")
		err := foobarUsecases.foobarRepository.Create(foobar)
		if err != nil {
			fatalError <- err
		}
	})
	if err != nil {
		return err
	}
	close(fatalError)
	return <-fatalError
}

func (foobarUsecases foobarUsecase) Update(foobar *models.Foobar) error {
	fatalError := make(chan error, 1)

	foobar.UpdatedAt = time.Now()
	key := fmt.Sprintf("%s:%s", foobar.ID.String(), "foobar")
	err := foobarUsecases.cacheCommand.WB(key, &foobar, func(
		wg *sync.WaitGroup,
		mtx *sync.Mutex,
	) {
		mtx.Lock()
		fmt.Println("Write Back db event")
		err := foobarUsecases.foobarRepository.Update(foobar)
		if err != nil {
			fatalError <- err
		}
		mtx.Unlock()
		wg.Done()
	})
	if err != nil {
		return err
	}
	close(fatalError)
	return <-fatalError
}

func (foobarUsecases foobarUsecase) Delete(id uuid.UUID) error {
	key := fmt.Sprintf("%s:%s", id.String(), "foobar")
	foobarUsecases.cacheCommand.Del(key)
	return foobarUsecases.foobarRepository.Delete(id)
}

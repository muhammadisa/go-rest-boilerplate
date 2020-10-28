package errhandler

import (
	"log"
)

type event func()

// HandleErrorThenReturn error handler returning
func HandleErrorThenReturn(err error) error {
	if err != nil {
		return err
	}
	return nil
}

// HandleError basic
func HandleError(err error, panics bool) {
	if err != nil {
		log.Fatalf("Error %v", err)
		if panics {
			panic(err)
		}
	}
}

// HandleErrorWithEvent with event
func HandleErrorWithEvent(err error, panics bool, evnt event) {
	if err != nil {
		evnt()
		log.Fatalf("Error %v", err)
		if panics {
			panic(err)
		}
	}
}

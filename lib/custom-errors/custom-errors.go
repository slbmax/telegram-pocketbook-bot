package custom_errors

import (
	"errors"
	"fmt"
)

func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func WrapIfErr(msg string, err error) error {
	if err == nil {
		return nil
	}

	return Wrap(msg, err)
}

func New(message string) error {
	return errors.New(message)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

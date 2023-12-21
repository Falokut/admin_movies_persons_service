package service

import (
	"errors"
	"fmt"
	"regexp"

	movies_persons_service "github.com/Falokut/admin_movies_persons_service/pkg/admin_movies_persons_service/v1/protos"
)

var ErrInvalidParam = errors.New("invalid param value, param must contain only digits and commas")
var ErrEmptyParam = errors.New("invalid param value, param mustn't be empty")

func validateParams(param *movies_persons_service.GetPersonsRequest) error {
	if param.PersonsIDs == "" {
		return ErrEmptyParam
	}
	if err := checkParam(param.PersonsIDs); err != nil {
		return err
	}

	return nil
}

func validateLimitAndPage(page, limit int32) error {
	if page <= 0 {
		return fmt.Errorf("%s error: %w", "page must be > 0", ErrInvalidArgument)
	}
	if limit < 10 || limit > 100 {
		return fmt.Errorf("%s error: %w", "limit must in range [10;100]", ErrInvalidArgument)
	}

	return nil
}

func checkParam(val string) error {
	exp := regexp.MustCompile("^[!-&!+,0-9]+$")

	if !exp.Match([]byte(val)) {
		return ErrInvalidParam
	}

	return nil
}

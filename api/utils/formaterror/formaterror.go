package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "username") {
		return errors.New("Username already taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email already taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title already taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect password")
	}
	return errors.New("Incorrect details")
}

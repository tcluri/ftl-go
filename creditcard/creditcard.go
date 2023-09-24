package creditcard

import (
	"errors"
)

type card struct {
	number string
}

func New(num string) (card, error) {
	if num == "" {
		return card{}, errors.New("number must not be empty")
	}
	return card{num}, nil
}

func (c *card) Number() string {
	return c.number
}

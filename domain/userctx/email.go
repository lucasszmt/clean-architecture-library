package userctx

import "regexp"

type Email struct {
	address string
}

func NewEmail(address string) (*Email, error) {
	pattern, _ := regexp.Compile(`[A-Za-z]+(.*)@[a-z]+(.*)(\.com(\.[a-z]+)?)?`)
	if pattern.MatchString(address) == true {
		return &Email{address: address}, nil
	}
	return nil, ErrInvalidEmail
}

func (e *Email) String() string {
	return e.address
}

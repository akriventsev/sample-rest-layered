package entities

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Married   bool      `json:"married"`
	Password  string    `json:"-"`
	Login     string    `json:"login"`
	Birthday  time.Time `json:"birthday"`
}

const (
	AllowedAge = 18
)

var (
	ErrAgeNotAllowed = errors.New("age must be greater than 18")
)

func (u User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func AgeGte(birthday time.Time, at time.Time, age int) bool {
	ageAt := birthday.AddDate(age, 0, 0)

	return ageAt.Before(at) || ageAt.Equal(at)
}

func (u User) Validate() error {
	if !AgeGte(u.Birthday, time.Now(), AllowedAge) {
		return ErrAgeNotAllowed
	}

	return nil
}

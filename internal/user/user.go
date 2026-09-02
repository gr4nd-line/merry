package user

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	NameMustNotBeEmptyError                  = errors.New("name must not be empty")
	InvalidEmailError                        = errors.New("invalid e-mail")
	PasswordMustNotBeEmptyError              = errors.New("password must not be empty")
	InvalidPasswordLenghtError               = errors.New("password length must be greater then 5")
	PasswordMustContainNumberError           = errors.New("password must contain a number")
	PasswordMustContainLetterError           = errors.New("password must contain a letter")
	PasswordMustContainSpecialCharacterError = errors.New("password must contain a special character")
)

type User struct {
	id        string
	name      string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(name, email, password string) (User, error) {
	if ok, err := validateName(strings.TrimSpace(name)); !ok {
		return User{}, err
	}
	if ok, err := validateEmail(email); !ok {
		return User{}, err
	}

	if ok, err := validatePassword(password); !ok {
		return User{}, err
	}

	hasPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("error hashing password")
	}

	return User{
		id:        uuid.NewString(),
		name:      name,
		email:     email,
		password:  string(hasPass),
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func (u User) Id() string {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) Email() string {
	return u.email
}

func (u User) Password() string {
	return u.password
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

func (u User) UpdatedAt() time.Time {
	return u.updatedAt
}

//TODO: Add methods to change name, password and e-mail

func validateName(name string) (bool, error) {
	if len(name) <= 0 {
		return false, NameMustNotBeEmptyError
	}
	return true, nil
}

func validateEmail(email string) (bool, error) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	ok := emailRegex.MatchString(email)
	if !ok {
		return false, InvalidEmailError
	}

	return true, nil
}

func validatePassword(password string) (bool, error) {
	if len(password) == 0 {
		return false, PasswordMustNotBeEmptyError
	}

	if len(password) < 6 {
		return false, InvalidPasswordLenghtError
	}

	var hasLetter, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsLetter(char):
			hasLetter = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasLetter {
		return false, PasswordMustContainLetterError
	}
	if !hasNumber {
		return false, PasswordMustContainNumberError
	}
	if !hasSpecial {
		return false, PasswordMustContainSpecialCharacterError
	}

	return true, nil
}

package user

import "testing"

type validateNameTestCase struct {
	title    string
	input    string
	expected bool
	err      error
}

func TestValidateName(t *testing.T) {
	tests := []validateNameTestCase{
		{
			title:    "should validate a valid name",
			input:    "John Doe",
			expected: true,
			err:      nil,
		},
		{
			title:    "should invalidate a empty name",
			input:    "",
			expected: false,
			err:      NameMustNotBeEmptyError,
		},
	}

	for _, test := range tests {
		result, err := validateName(test.input)
		if result != test.expected || err != test.err {
			t.Errorf("test title: %s, expected: %t, got: %t -- error expected: %s, error got: %s", test.title, test.expected, result, test.err, err)
		}
	}

}

type validateEmailTestCase struct {
	title    string
	input    string
	expected bool
	err      error
}

func TestValidateEmail(t *testing.T) {
	tests := []validateEmailTestCase{
		{
			title:    "should validate a valid email",
			input:    "johndoe@email.com",
			expected: true,
			err:      nil,
		},
		{
			title:    "should validate a valid email with .(dot)",
			input:    "john.doe@email.com",
			expected: true,
			err:      nil,
		},
		{
			title:    "should invalidate a email without @",
			input:    "johndoeemail.com",
			expected: false,
			err:      InvalidEmailError,
		},
		{
			title:    "should invalidate a email without .(dot)",
			input:    "johndoe@emailcom",
			expected: false,
			err:      InvalidEmailError,
		},
		{
			title:    "should invalidate a email without domain",
			input:    "johndoe@",
			expected: false,
			err:      InvalidEmailError,
		},
	}

	for _, test := range tests {
		result, err := validateEmail(test.input)
		if result != test.expected || err != test.err {
			t.Errorf("test title: %s, expected: %t, got: %t -- error expected: %s, error got: %s", test.title, test.expected, result, test.err, err)
		}
	}

}

type validatePasswordTestCase struct {
	title    string
	input    string
	expected bool
	err      error
}

func TestValidatePassword(t *testing.T) {
	tests := []validatePasswordTestCase{
		{
			title:    "should validate a valid password",
			input:    "Pass@123",
			expected: true,
			err:      nil,
		},
		{
			title:    "should invalidate a empty password",
			input:    "",
			expected: false,
			err:      PasswordMustNotBeEmptyError,
		},
		{
			title:    "should invalidate a password with less then 6 digits",
			input:    "Ps@12",
			expected: false,
			err:      InvalidPasswordLenghtError,
		},
		{
			title:    "should invalidate a password without a number",
			input:    "Psasas@",
			expected: false,
			err:      PasswordMustContainNumberError,
		},
		{
			title:    "should invalidate a password without special characters",
			input:    "Passa12312",
			expected: false,
			err:      PasswordMustContainSpecialCharacterError,
		},
		{
			title:    "should invalidate a password without a letter",
			input:    "@123$%@!112312",
			expected: false,
			err:      PasswordMustContainLetterError,
		},
	}

	for _, test := range tests {
		result, err := validatePassword(test.input)
		if result != test.expected || err != test.err {
			t.Errorf("test title: %s, expected: %t, got: %t -- error expected: %s, error got: %s", test.title, test.expected, result, test.err, err)
		}
	}

}

func TestNewUser(t *testing.T) {
	testTitle := "test new user"

	name := "John Doe"
	email := "john.doe@email.com"
	pass := "Pass@123"

	user, err := NewUser(name, email, pass)
	if err != nil {
		t.Errorf("test title: %s -- error: %s", testTitle, err)
	}

	if len(user.Id()) == 0 {
		t.Errorf("test title: %s -- problem with user id: %s", testTitle, user.Id())
	}

	if name != user.Name() {
		t.Errorf("test title: %s -- user name expected: %s, got: %s", testTitle, name, user.Name())
	}

	if email != user.Email() {
		t.Errorf("test title: %s -- user email expected: %s, got: %s", testTitle, email, user.Email())
	}

	if pass == user.Password() {
		t.Errorf("test title: %s -- user password not expected: %s, got: %s", testTitle, pass, user.Password())
	}
}

package validators
// TODO do actual input validation
import "errors"

func ValidateUsername(username string) error {
	if len(username) < 4 { // TODO
		return errors.New("username cannot be shorter than 4 characters")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 12 { // TODO
		return errors.New("password cannot be shorter than 12 characters")
	}
	return nil
}

func ValidateEmail(email string) error {
	return nil // TODO
}
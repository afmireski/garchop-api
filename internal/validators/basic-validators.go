package validators

import (
	"regexp"
	"time"

	"github.com/google/uuid"
)

func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}

func IsValidName(name string, min int, max int) bool {
	length := len(name)
	return min <= length && length <= max
}

func IsValidUuid(id string) bool {
	_, err := uuid.Parse(id)

	return err == nil
}

func IsValidNumericId(id int32) bool {
	return id > 0
}

func IsValidAge(birthDate time.Time, age int) bool {
	// Calculate the date `age`` years ago from today
	minDate := time.Now().AddDate(-age, 0, 0)

	// Compare the birth date with the minimum date
	return !birthDate.After(time.Now()) && (birthDate.Before(minDate) || birthDate.Equal(minDate))
}

func IsPhoneNumber(phone string) bool {
	if phone == "" {
		return false
	}
	return regexp.MustCompile(`^\+55\D*([1-9]{2})\D*9?\D*(\d{4})\D*(\d{4})$`).MatchString(phone)
}

func MinLength(texr string, min int) bool {
	return len(texr) >= min
}

func IsValidPassword(password string) bool {
	re := regexp.MustCompile(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*()])(?=.{6,})`)
	return re.MatchString(password)
}

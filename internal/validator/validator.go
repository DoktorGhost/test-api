package validator

import (
	"api/internal/storage"
	"errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
	"strings"
)

// Функция для валидации имени
func validateName(name string) (string, error) {
	// Имя должно содержать только буквы и иметь длину от 2 до 50 символов
	re := regexp.MustCompile(`^[а-яА-Яa-zA-Z]{2,50}$`)
	if re.MatchString(name) {
		f := cases.Title(language.Und)
		return f.String(name), nil
	}
	return "", errors.New("invalid name")
}

// Функция для валидации фамилии
func validateLastname(surname string) (string, error) {
	// Фамилия должна содержать только буквы и иметь длину от 2 до 50 символов
	re := regexp.MustCompile(`^[а-яА-Яa-zA-Z]{2,50}$`)
	if re.MatchString(surname) {
		f := cases.Title(language.Und)
		return f.String(surname), nil
	}
	return "", errors.New("invalid surname")
}

// Функция для валидации email
func validateEmail(email string) (string, error) {
	// Простая проверка формата email с использованием регулярного выражения
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if re.MatchString(email) {
		return strings.ToLower(email), nil
	}
	return "", errors.New("invalid email")
}

// Функция для валидации возраста
func validateAge(age uint) (uint, error) {
	// Возраст должен быть положительным числом и не превышать 150
	if age > 0 && age <= 150 {
		return age, nil
	}
	return 0, errors.New("invalid age")
}

func ValidateStruct(user storage.User) error {
	if _, err := validateName(user.Firstname); err != nil {
		return err
	}
	if _, err := validateLastname(user.Lastname); err != nil {
		return err
	}
	if _, err := validateAge(user.Age); err != nil {
		return err
	}
	if _, err := validateEmail(user.Email); err != nil {
		return err
	}
	return nil
}

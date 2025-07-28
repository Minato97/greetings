package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formarts := []string{
		"Hola, %v! Bienvenido",
		"Que bueno verte, %v!",
		"Saludo, %v! Encantado de conocerte!",
	}
	return formarts[rand.Intn(len(formarts))]
}

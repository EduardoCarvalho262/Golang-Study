package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	//Cria um Map
	messages := make(map[string]string)

	//Um for que percorre o slice de nomes
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		//Adiciona a chave name com a mensagem que tem
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

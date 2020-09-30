package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrWordDoesNotExist = errors.New("Word does not exist")

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrWordDoesNotExist
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}

func main() {
}

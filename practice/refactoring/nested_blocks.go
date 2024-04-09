package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	s1 := "Hello"
	s2 := "World"

	s, err := join(s1, s2, 20)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(s)
}

func join(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

func concatenate(s1 string, s2 string) (string, error) {
	return s1 + s2, nil
}

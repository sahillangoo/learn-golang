package util

import "os"

func ReadTextFile(file string) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		// error
		return "", err
	} else {
		// success
		return string(content), nil
	}
}

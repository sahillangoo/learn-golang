package util

import "os"

func WriteTextFile(file string, content string) error {
	return os.WriteFile(file, []byte(content), 0644)

}

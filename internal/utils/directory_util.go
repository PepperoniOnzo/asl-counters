package utils

import "os"

func GetExecutionPath() (string, error) {
	executablePath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return executablePath, nil
}

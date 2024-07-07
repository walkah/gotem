package util

import (
	"fmt"
	"os"
	"strings"
)

func GetRelativePath(absolutePath string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(absolutePath, homeDir) {
		relativePath := strings.TrimPrefix(absolutePath, fmt.Sprintf("%s/", homeDir))
		return relativePath, nil
	}

	return absolutePath, nil
}

func GetAbsolutePath(relativePath string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(relativePath, "/") {
		relativePath = fmt.Sprintf("%s/%s", homeDir, relativePath)
	}

	return relativePath, nil
}

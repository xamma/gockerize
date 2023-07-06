package kube

import (
	"fmt"
	"os"
)

func SaveManifestToFile(manifest, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(manifest)
	if err != nil {
		return fmt.Errorf("failed to write manifest to file: %v", err)
	}

	return nil
}
